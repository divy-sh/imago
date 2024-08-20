package app

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"sort"
	"sync"
)

type Img struct {
	r [][]float64
	g [][]float64
	b [][]float64
	a [][]float64
	h int
	w int
}

func Load(path string) (*Img, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	m, _, err := image.Decode(reader)
	defer reader.Close()
	if err != nil {
		return nil, err
	}
	bounds := m.Bounds()

	height := bounds.Max.X - bounds.Min.X
	width := bounds.Max.Y - bounds.Min.Y
	red := make([][]float64, height)
	green := make([][]float64, height)
	blue := make([][]float64, height)
	alpha := make([][]float64, height)

	// Allocate space for each row in the slices.
	for i := range red {
		red[i] = make([]float64, width)
		green[i] = make([]float64, width)
		blue[i] = make([]float64, width)
		alpha[i] = make([]float64, height)
	}

	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			r, g, b, a := m.At(y, x).RGBA()
			red[x][y] = float64(r)
			green[x][y] = float64(g)
			blue[x][y] = float64(b)
			alpha[x][y] = float64(a)
		}
	}
	return &Img{
		r: red,
		g: green,
		b: blue,
		a: alpha,
		h: height,
		w: width,
	}, nil
}

// SaveAsPNG function to save the Img as a PNG file
func (img *Img) SaveAsPNG(filename string) error {
	// Create a new RGBA image with the same width and height
	rgba := image.NewRGBA(image.Rect(0, 0, img.w, img.h))

	// Populate the RGBA image with data from the Img struct
	for y := 0; y < img.h; y++ {
		for x := 0; x < img.w; x++ {
			r := uint8(img.r[y][x] * 255)
			g := uint8(img.g[y][x] * 255)
			b := uint8(img.b[y][x] * 255)
			a := uint8(img.a[y][x] * 255)
			rgba.Set(x, y, color.RGBA{r, g, b, a})
		}
	}

	// Create a file to save the PNG
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the RGBA image as a PNG and write it to the file
	if err := png.Encode(file, rgba); err != nil {
		return err
	}

	return nil
}

func (img *Img) Compress(ratio float64) (*Img, error) {
	if ratio < 0 || ratio > 1 {
		return nil, errors.New("invalid compression ratio")
	}

	size := max(nextPow2(img.h), nextPow2(img.w))
	imageData := make([][][]float64, 4)
	for i := range imageData {
		imageData[i] = make([][]float64, size)
		for j := range imageData[i] {
			imageData[i][j] = make([]float64, size)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < img.h; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < img.w; j++ {
				imageData[0][i][j] = img.r[i][j]
				imageData[1][i][j] = img.g[i][j]
				imageData[2][i][j] = img.b[i][j]
				imageData[3][i][j] = img.a[i][j]
			}
		}(i)
	}
	wg.Wait()

	for i := 0; i < 4; i++ {
		haarTransform2D(imageData[i], size)
		setValuesZero(imageData[i], ratio)
		inverseHaarTransform2D(imageData[i], size)
	}

	red := make([][]float64, img.h)
	green := make([][]float64, img.h)
	blue := make([][]float64, img.h)
	alpha := make([][]float64, img.h)

	// Allocate space for each row in the slices.
	for i := range red {
		red[i] = make([]float64, img.w)
		green[i] = make([]float64, img.w)
		blue[i] = make([]float64, img.w)
		alpha[i] = make([]float64, img.w)
	}
	for i := 0; i < img.h; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < img.w; j++ {
				red[i][j] = imageData[0][i][j]
				green[i][j] = imageData[1][i][j]
				blue[i][j] = imageData[2][i][j]
				alpha[i][j] = imageData[3][i][j]

			}
		}(i)
	}
	wg.Wait()

	return &Img{
		r: red,
		g: green,
		b: blue,
		a: alpha,
		h: img.h,
		w: img.w,
	}, nil
}

func setValuesZero(comp [][]float64, compRatio float64) {
	if compRatio == 0 {
		return
	}

	var flattenedComp []float64
	for _, rows := range comp {
		for _, val := range rows {
			if math.Abs(val) >= 0.001 {
				flattenedComp = append(flattenedComp, math.Abs(val))
			}
		}
	}
	sort.Float64s(flattenedComp)
	minNum := flattenedComp[int(compRatio*float64(len(flattenedComp)-1))]

	var wg sync.WaitGroup
	for i := range comp {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := range comp[i] {
				if math.Abs(comp[i][j]) <= minNum {
					comp[i][j] = 0
				}
			}
		}(i)
	}
	wg.Wait()
}

// haarTransform2D applies the Haar transform on a 2D float64 array.
func haarTransform2D(comp [][]float64, size int) {
	c := size
	for c > 1 {
		len := c
		var wg sync.WaitGroup

		for i := 0; i < size; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				haarPartial(comp[i], len)
			}(i)
		}
		wg.Wait()

		for i := 0; i < size; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				col := make([]float64, len)
				for j := 0; j < len; j++ {
					col[j] = comp[j][i]
				}
				haarPartial(col, len)
				for j := 0; j < len; j++ {
					comp[j][i] = col[j]
				}
			}(i)
		}
		wg.Wait()

		c /= 2
	}
}

// inverseHaarTransform2D applies the inverse Haar transform on a 2D float64 array.
func inverseHaarTransform2D(comp [][]float64, size int) {
	c := 2
	for c <= size {
		len := c
		var wg sync.WaitGroup

		for i := 0; i < size; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				col := make([]float64, len)
				for j := 0; j < len; j++ {
					col[j] = comp[j][i]
				}
				inverseHaarPartial(col, len)
				for j := 0; j < len; j++ {
					comp[j][i] = col[j]
				}
			}(i)
		}
		wg.Wait()

		for i := 0; i < size; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				inverseHaarPartial(comp[i], len)
			}(i)
		}
		wg.Wait()

		c *= 2
	}
}

// haarPartial applies a partial Haar transform (one step) on a float64 array.
func haarPartial(comp []float64, length int) {
	sqrt2 := math.Sqrt(2)
	pass := make([]float64, length)
	var wg sync.WaitGroup

	for i := 0; i < length; i += 2 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			pass[i/2] = (comp[i] + comp[i+1]) / sqrt2
			pass[(i+length)/2] = (comp[i] - comp[i+1]) / sqrt2
		}(i)
	}
	wg.Wait()

	copy(comp, pass)
}

// inverseHaarPartial applies a partial inverse Haar transform (one step) on a float64 array.
func inverseHaarPartial(comp []float64, length int) {
	sqrt2 := math.Sqrt(2)
	pass := make([]float64, length)
	var wg sync.WaitGroup

	for i := 0; i < length; i += 2 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			pass[i] = (comp[i/2] + comp[(i+length)/2]) / sqrt2
			pass[i+1] = (comp[i/2] - comp[(i+length)/2]) / sqrt2
		}(i)
	}
	wg.Wait()

	copy(comp, pass)
}

// nextPow2 returns the next power of 2 greater than or equal to num.
func nextPow2(num int) int {
	res := 1
	for res < num {
		res <<= 1
	}
	return res
}
