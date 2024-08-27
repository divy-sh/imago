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

// load image
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
	image, err := NewImage(height, width)
	if err != nil {
		return nil, err
	}
	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			red, green, blue, alpha := m.At(y, x).RGBA()
			image.p[x][y] = Pixel{
				r: float64(red),
				g: float64(green),
				b: float64(blue),
				a: float64(alpha),
			}
		}
	}
	return image, nil
}

// SaveAsPNG function to save the Img as a PNG file
func (img *Img) SaveAsPNG(filename string) error {
	// Create a new RGBA image with the same width and height
	rgba := image.NewRGBA(image.Rect(0, 0, img.w, img.h))

	// Populate the RGBA image with data from the Img struct
	for y := 0; y < img.h; y++ {
		for x := 0; x < img.w; x++ {
			r := uint8(img.p[y][x].r * 255)
			g := uint8(img.p[y][x].g * 255)
			b := uint8(img.p[y][x].b * 255)
			a := uint8(img.p[y][x].a * 255)
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

func (img *Img) HorizontalFlip() (*Img, error) {
	newImg, _ := NewImage(img.h, img.w)
	for i := range img.h {
		for j := range img.w {
			newImg.p[i][j] = img.p[i][img.w-j-1]
		}
	}
	return newImg, nil
}
