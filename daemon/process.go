package daemon

import (
	"encoding/base64"
	"log"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/otiai10/gosseract/v2"
)

type Memoria struct {
	targetDisplay      int
	captureQuality     int
	version            int
	recordingFrequency time.Duration
	tesseractClient    *gosseract.Client
	influxDBWriteAPI   *api.WriteAPI
}

func NewMemoria(display int, captureFreq time.Duration, tc *gosseract.Client, db *api.WriteAPI) *Memoria {
	if captureFreq == 0 {
		captureFreq = 10 * time.Second
	}
	return &Memoria{
		targetDisplay:      display,
		captureQuality:     80,
		version:            1,
		recordingFrequency: captureFreq,
		tesseractClient:    tc,
		influxDBWriteAPI:   db,
	}
}

func (m *Memoria) Record() {
	for {
		ss, err := capture(m.targetDisplay, m.captureQuality)
		if err != nil {
			panic(err)
		}

		txt, err := ocr(m.tesseractClient, ss)
		if err != nil {
			panic(err)
		}

		ssInBase64 := base64.StdEncoding.EncodeToString(ss)

		saveToDB(*m.influxDBWriteAPI, m.captureQuality, m.version, txt, ssInBase64)
		log.Println(txt)
		time.Sleep(m.recordingFrequency)
	}
}
