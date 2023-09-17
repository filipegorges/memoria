package memoria

import (
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

const (
	MEASUREMENT = "history"
)

func SaveToDB(writer api.WriteAPI, quality, version int, txt, ss string) {
	writer.WritePoint(write.NewPoint(
		MEASUREMENT,
		map[string]string{
			"quality": string(quality),
			"version": string(version),
		},
		map[string]interface{}{
			"ocr": txt,
			"ss":  ss,
		},
		time.Now(),
	))
}
