package editor

import (
	"bytes"
	"fmt"
	img "image"
	"image/png"
	"io"

	blr "github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/transform"
)

func ProcessImage(reader io.Reader, cX, cY, width, height, angle, resize, blur string) []byte {
	image, _, err := img.Decode(reader)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}

	x, y, w, h, a, r, b :=
		parseI(cX, 0),
		parseI(cY, 0),
		parseI(width, image.Bounds().Dx()),
		parseI(height, image.Bounds().Dy()),
		parseF(angle, 0),
		parseI(resize, 1),
		parseF(blur, 0)

	fmt.Printf("T: %v x %v - F: %v x %v - A: %v° - B: %v%%\n", w, h, image.Bounds().Dx(), image.Bounds().Dy(), a, b)

	if x != 0 || y != 0 {
		image = transform.Crop(image, img.Rect(x, y, w, h))
	} else if w != image.Bounds().Dx() || h != image.Bounds().Dy() {
		image = transform.Resize(image, w, h, transform.Linear)
	}

	if b != 0 {
		image = blr.Gaussian(image, b)
	}

	if a != 0 {
		image = transform.Rotate(image, a, &transform.RotationOptions{
			ResizeBounds: r > 0,
			Pivot:        &img.Point{0, 0},
		})
	}

	buffer := bytes.NewBuffer(nil)
	encoder := &png.Encoder{
		CompressionLevel: png.NoCompression,
	}
	encoder.Encode(buffer, image)
	return buffer.Bytes()
}
