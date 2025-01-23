package app

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

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
			// Normalize the 16-bit color values (0-65535) to the range [0, 1]
			image.p[y][x] = Pixel{
				r: float64(red) / 65535.0,
				g: float64(green) / 65535.0,
				b: float64(blue) / 65535.0,
				a: float64(alpha) / 65535.0,
			}
		}
	}
	return image, nil
}

func SaveImage(img *Img, filename string) error {
	fileSplit := strings.Split(filename, ".")
	if len(fileSplit) <= 0 {
		return errors.New("file extension not provided for image")
	}
	extension := fileSplit[len(fileSplit)-1]
	switch extension {
	case "jpg":
	case "jpeg":
		saveAsJpeg(img, filename)
	case "png":
		saveAsPng(img, filename)
	}
	return nil
}

/*
Private functions beyond this point
*/
func processImageForSave(img *Img) *image.RGBA {
	rgba := image.NewRGBA(image.Rect(0, 0, img.w, img.h))
	for y := 0; y < img.h; y++ {
		for x := 0; x < img.w; x++ {
			r := uint8(img.p[y][x].r * 255)
			g := uint8(img.p[y][x].g * 255)
			b := uint8(img.p[y][x].b * 255)
			a := uint8(img.p[y][x].a * 255)
			rgba.Set(x, y, color.RGBA{r, g, b, a})
		}
	}
	return rgba
}

// SaveAsPNG function to save the Img as a PNG file
func saveAsPng(img *Img, filename string) error {
	rgba := processImageForSave(img)
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

// SaveAsJPEG function to save the Img as a PNG file
func saveAsJpeg(img *Img, filename string) error {
	rgba := processImageForSave(img)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := jpeg.Encode(file, rgba, nil); err != nil {
		return err
	}

	return nil
}
