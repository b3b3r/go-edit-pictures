package filter

import (
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
)

// Filter is the interface to process differents filters
type Filter interface {
	Process(srcPath, distPath string) error
}

//GrayScale is the type for gray scale filter
type GrayScale struct{}

// Process process picture effects
func (g GrayScale) Process(srcPath, dstPath string) error {
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	// Create a grayscale version of the image with higher contrast and sharpness.
	img := imaging.Grayscale(src)
	img = imaging.AdjustContrast(img, 20)
	img = imaging.Sharpen(img, 2)

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	opts := jpeg.Options{Quality: 90}
	return jpeg.Encode(dstFile, img, &opts)
}

//Blur is the type for gray scale filter
type Blur struct{}

// Process process picture effects
func (b Blur) Process(srcPath, dstPath string) error {
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	img := imaging.Blur(src, 3.5)

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	opts := jpeg.Options{Quality: 90}
	return jpeg.Encode(dstFile, img, &opts)
}
