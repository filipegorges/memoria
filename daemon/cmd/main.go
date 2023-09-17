package main

import (
	"encoding/base64"
	"log"
	"time"

	"github.com/filipegorges/memoria"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/otiai10/gosseract/v2"
)

const (
	QUALITY                    = 100
	VERSION                    = 1
	ITERATION_DELAY_IN_SECONDS = 10
	DISPLAY                    = 0
	INFLUX_DB_API_TOKEN        = "PpVRM1LN0rVqjbPOnaxMK9Tr7Jo7FwuXkXaJATRwGvyGYnemrshcH_DNI4DUhklBVSeJOXfspS1_Eoop4U-y5Q=="
	INFLUXDB_TOKEN             = "tHQFmdtBctPgqA1JCGZi7KxbcwEOsqvOFfySe8Bs-y21JxFXSw8yg88XRAlSDEFgMf9R4pMVWeMGe6UJyn0oVQ=="
	DB_URL                     = "http://localhost:8086"
	DB_ORG                     = "memoria"
	DB_BUCKET                  = "memoria"
)

func main() {
	db := influxdb2.NewClient(DB_URL, INFLUX_DB_API_TOKEN)
	defer db.Close()
	writeAPI := db.WriteAPI(DB_ORG, DB_BUCKET)

	tc := gosseract.NewClient()
	defer tc.Close()

	for {
		ss, err := memoria.Capture(DISPLAY, QUALITY)
		if err != nil {
			panic(err)
		}

		txt, err := memoria.OCR(tc, ss)
		if err != nil {
			panic(err)
		}

		ssInBase64 := base64.StdEncoding.EncodeToString(ss)

		memoria.SaveToDB(writeAPI, QUALITY, VERSION, txt, ssInBase64)
		log.Println(txt)
		time.Sleep(ITERATION_DELAY_IN_SECONDS * time.Second)
	}
}
