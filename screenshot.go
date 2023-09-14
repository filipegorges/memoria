package memoria

import (
	"bytes"
	"errors"
	"fmt"
	"image/jpeg"

	"github.com/kbinani/screenshot"
)

func Capture(display, quality int) ([]byte, error) {
	// Get the total number of screens
	n := screenshot.NumActiveDisplays()

	// If there's no active display, exit
	if n <= 0 {
		return nil, errors.New("no active displays found")
	}

	// For the purpose of this example, we're capturing the screenshot of the primary display
	bounds := screenshot.GetDisplayBounds(display)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return nil, fmt.Errorf("failed to capture screen: %v", err)
	}

	var buf bytes.Buffer
	err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality})
	if err != nil {
		return nil, fmt.Errorf("failed to encode image: %v", err)
	}

	return buf.Bytes(), nil
}
