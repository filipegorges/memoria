package daemon

import (
	"fmt"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

const (
	measurement = "history"
)

func saveToDB(writer api.WriteAPI, quality, version int, txt, ss string) {
	writer.WritePoint(write.NewPoint(
		measurement,
		map[string]string{
			"quality": fmt.Sprint(quality),
			"version": fmt.Sprint(version),
		},
		map[string]interface{}{
			"ocr": txt,
			"ss":  ss,
		},
		time.Now(),
	))
}
