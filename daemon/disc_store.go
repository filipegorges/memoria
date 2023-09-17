package daemon

import (
	"fmt"
	"log"
	"os"
	"time"
)

func saveToDisc(img []byte) error {
	fileName := fmt.Sprintf("./assets/screenshot_%s.jpeg", time.Now().Format("20060102_150405"))
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()
	_, err = file.Write(img)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}
	log.Printf("Screenshot saved as %s\n", fileName)
	return nil
}
