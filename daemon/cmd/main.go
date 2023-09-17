package main

import (
	"fmt"
	"log"
	"time"

	"github.com/filipegorges/memoria/daemon"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/otiai10/gosseract/v2"
)

type Env struct {
	Quality          int    `envconfig:"MEMORIA_QUALITY"`
	Version          int    `envconfig:"MEMORIA_VERSION"`
	CaptureFrequency int    `envconfig:"MEMORIA_CAPTURE_FREQUENCY"`
	Display          int    `envconfig:"MEMORIA_DISPLAY"`
	InfluxDBAPIToken string `envconfig:"MEMORIA_DB_API_TOKEN"`
	InfluxDBURL      string `envconfig:"MEMORIA_DB_URL"`
	InfluxDBOrg      string `envconfig:"MEMORIA_DB_ORG"`
	InfluxDBBucket   string `envconfig:"MEMORIA_DB_BUCKET"`
}

func main() {
	log.Println("Starting Memoria daemon...")
	log.Println("Reading environment variables...")

	// TODO: we'll want to remove this direct dependency over time
	if err := godotenv.Load("./daemon/.env/local.env"); err != nil {
		panic(fmt.Sprintf("Error loading .env file: %v", err))
	}
	var env Env
	if err := envconfig.Process("MEMORIA", &env); err != nil {
		panic(err)
	}

	// TODO: remove this log as it'll print the API key in it
	log.Printf("%#v", env)
	log.Println("Environment variables read successfully")
	log.Println("Starting daemon...")
	db := influxdb2.NewClient(env.InfluxDBURL, env.InfluxDBAPIToken)
	defer db.Close()
	writeAPI := db.WriteAPI(env.InfluxDBOrg, env.InfluxDBBucket)

	tc := gosseract.NewClient()
	defer tc.Close()

	memoria := daemon.NewMemoria(env.Display, time.Duration(env.CaptureFrequency)*time.Second, tc, &writeAPI)
	log.Println("Daemon started successfully")
	memoria.Record()
}
