# imago

`imago` is a command-line interface (CLI) image editor written in Go. It supports a variety of image processing operations for enhancing or modifying images.

## List of Supported Operations

- **Blur**: Apply a blur filter to the image.
- **Sharpen**: Apply a sharpen filter to the image.
- **Greyscale**: Convert the image to greyscale.
- **Sepia**: Apply a sepia filter to the image.
- **Change Brightness**: Adjust the brightness of the image.
- **Get Specific Color Components**: Extract specific color components (red, green, blue) from the image.
- **Horizontal Flip**: Flip the image horizontally.
- **Vertical Flip**: Flip the image vertically.
- **Apply Greyscale by Value**: Convert the image to greyscale using the value component of the pixels.
- **Apply Greyscale by Intensity**: Convert the image to greyscale using the intensity component of the pixels.
- **Apply Greyscale by Luma**: Convert the image to greyscale using the luma component of the pixels.
- **Compress an Image**: Compress the image using the Haar wavelet transform.
- **Generate a Histogram**: Generate a histogram of the image.

## TODO

- **Adjust Levels**: Adjust the level of an image using black, mid, and white values.
- **Color Correct an Image**: Apply color correction to the image.

## How to run

- run this command in the root directory.
```
    go build .
```
- run a command using -
```
    ./imago <command> <input image> <output image> <other arguments>
```
