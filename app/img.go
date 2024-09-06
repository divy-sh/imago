package app

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Pixel struct {
	r float64
	g float64
	b float64
	a float64
}

type Img struct {
	p [][]Pixel
	h int
	w int
}

// create a new blank image
func NewImage(height, width int) (*Img, error) {
	if height <= 0 || width <= 0 {
		return nil, errors.New("invalid image dimensions")
	}
	pixels := make([][]Pixel, height)

	for i := range pixels {
		pixels[i] = make([]Pixel, width)
	}
	return &Img{
		p: pixels,
		h: height,
		w: width,
	}, nil
}

// Load image from file path
func Load(path string) (*Img, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	bounds := m.Bounds()

	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y
	image, err := NewImage(height, width)
	if err != nil {
		return nil, err
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			red, green, blue, alpha := m.At(x, y).RGBA()
			image.p[y][x] = Pixel{
				r: float64(red) / 255.0,
				g: float64(green) / 255.0,
				b: float64(blue) / 255.0,
				a: float64(alpha) / 255.0,
			}
		}
	}
	return image, nil
}

// SaveAsPNG function to save the Img as a PNG file
func (img *Img) SaveAsPNG(filename string) error {
	rgba := image.NewRGBA(image.Rect(0, 0, img.w, img.h))
	for y := 0; y < img.h; y++ {
		for x := 0; x < img.w; x++ {
			r := uint8(img.p[y][x].r * 0xffff)
			g := uint8(img.p[y][x].g * 0xffff)
			b := uint8(img.p[y][x].b * 0xffff)
			a := uint8(img.p[y][x].a * 0xffff)
			rgba.Set(x, y, color.RGBA{r, g, b, a})
		}
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := png.Encode(file, rgba); err != nil {
		return err
	}

	return nil
}

// flip image horizontally
func (img *Img) HorizontalFlip() (*Img, error) {
	return process(
		func(i, j int, newImg *Img) {
			newImg.p[i][j] = img.p[i][img.w-j-1]
		}, img,
	)
}

// flip image vertically
func (img *Img) VerticalFlip() (*Img, error) {
	return process(
		func(i, j int, newImg *Img) {
			newImg.p[i][j] = img.p[img.h-i-1][j]
		}, img,
	)
}

func (img *Img) Brighten(bVal float64) (*Img, error) {
	return process(
		func(i, j int, newImg *Img) {
			newImg.p[i][j].r = clampPixelValue(img.p[i][j].r + bVal)
			newImg.p[i][j].g = clampPixelValue(img.p[i][j].g + bVal)
			newImg.p[i][j].b = clampPixelValue(img.p[i][j].b + bVal)
			newImg.p[i][j].a = img.p[i][j].a
		}, img,
	)
}

func (img *Img) GetRed() (*Img, error) {
	return process(
		func(i, j int, newImg *Img) {
			newImg.p[i][j].r = img.p[i][j].r
			newImg.p[i][j].a = img.p[i][j].a
		}, img,
	)
}

func (img *Img) GetGreen() (*Img, error) {
	return process(
		func(i, j int, newImg *Img) {
			newImg.p[i][j].g = img.p[i][j].g
			newImg.p[i][j].a = img.p[i][j].a
		}, img,
	)
}
func (img *Img) GetBlue() (*Img, error) {
	return process(
		func(i, j int, newImg *Img) {
			newImg.p[i][j].b = img.p[i][j].b
			newImg.p[i][j].a = img.p[i][j].a
		}, img,
	)
}

func (img *Img) GetGrayScaleByValue() (*Img, error) {
	return process(
		func(i, j int, newImg *Img) {
			maxVal := max(max(img.p[i][j].r, img.p[i][j].g), img.p[i][j].b)
			newImg.p[i][j].r = maxVal
			newImg.p[i][j].g = maxVal
			newImg.p[i][j].b = maxVal
			newImg.p[i][j].a = img.p[i][j].a
		}, img,
	)
}

func (img *Img) GetGrayScaleByIntensity() (*Img, error) {
	return process(
		func(i, j int, newImg *Img) {
			average := (img.p[i][j].r + img.p[i][j].g + img.p[i][j].b) / 3
			newImg.p[i][j].r = average
			newImg.p[i][j].g = average
			newImg.p[i][j].b = average
			newImg.p[i][j].a = img.p[i][j].a
		}, img,
	)
}

func (img *Img) Blur() (*Img, error) {
	blurFilter := [][]float64{
		{1.0 / 16, 1.0 / 8, 1.0 / 16},
		{1.0 / 8, 1.0 / 4, 1.0 / 8},
		{1.0 / 16, 1.0 / 8, 1.0 / 16},
	}
	return applyFilter(img, &blurFilter)
}

func (img *Img) Sharpen() (*Img, error) {
	sharpenFilter := [][]float64{
		{-1.0 / 8, -1.0 / 8, -1.0 / 8, -1.0 / 8, -1.0 / 8},
		{-1.0 / 8, 1.0 / 4, 1.0 / 4, 1.0 / 4, -1.0 / 8},
		{-1.0 / 8, 1.0 / 4, 1.0, 1.0 / 4, -1.0 / 8},
		{-1.0 / 8, 1.0 / 4, 1.0 / 4, 1.0 / 4, -1.0 / 8},
		{-1.0 / 8, -1.0 / 8, -1.0 / 8, -1.0 / 8, -1.0 / 8},
	}
	return applyFilter(img, &sharpenFilter)
}

func (img *Img) HaarCompress(ratio float32) (*Img, error) {
	if ratio < 0 || ratio > 1 {
		return nil, errors.New("invalid compression ratio")
	}
	return haarCompress(img, ratio)
}

func clampPixelValue(val float64) float64 {
	return min(max(val, 0), 0xfffff)
}

func process(f func(int, int, *Img), img *Img) (*Img, error) {
	newImg, _ := NewImage(img.h, img.w)
	for i := range img.h {
		for j := range img.w {
			f(i, j, newImg)
		}
	}
	return newImg, nil
}

func applyFilter(img *Img, filterPointer *[][]float64) (*Img, error) {
	filter := *filterPointer
	return process(
		func(h, w int, newImg *Img) {
			steps := len(filter) / 2
			for i := -steps; i <= steps; i++ {
				for j := -steps; j <= steps; j++ {
					if h+i >= 0 && h+i < img.h && w+j >= 0 && w+j < img.w {
						newImg.p[h][w].r += clampPixelValue(filter[i+steps][j+steps] * img.p[h+i][w+j].r)
						newImg.p[h][w].g += clampPixelValue(filter[i+steps][j+steps] * img.p[h+i][w+j].g)
						newImg.p[h][w].b += clampPixelValue(filter[i+steps][j+steps] * img.p[h+i][w+j].b)
					}
				}
			}
			newImg.p[h][w].a = img.p[h][w].a
		}, img,
	)
}
