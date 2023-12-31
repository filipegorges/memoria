package daemon

import (
	"github.com/otiai10/gosseract/v2"
)

func ocr(c *gosseract.Client, ss []byte) (string, error) {
	err := c.SetImageFromBytes(ss)
	if err != nil {
		return "", err
	}
	text, err := c.Text()
	if err != nil {
		return "", err
	}
	return text, nil
}
