package editor

import (
	"bytes"
	"fmt"
	img "image"
	"image/png"
	"io"

	"github.com/anthonynsimon/bild/transform"
)

func ProcessImage(reader io.Reader, width, height int) []byte {
	image, _, err := img.Decode(reader)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}

	image = transform.Resize(image, width, height, transform.Gaussian)

	buffer := bytes.NewBuffer(nil)
	encoder := &png.Encoder{
		CompressionLevel: png.NoCompression,
	}
	encoder.Encode(buffer, image)
	return buffer.Bytes()
}
