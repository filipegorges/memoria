# memoria
Generate a searchable visual history of your activity on your computer.

This project is both a tool for personal use as well as a means to practice a few buzzwordy shenanigans. **WIP**!

## Installation
The `daemon` service relies on `tesseract`, which is available on multiple platforms. For ubuntu (native or `WSL`), install these packages:

```bash
sudo apt-get install tesseract-ocr \ 
                    libtesseract-dev \
                    libleptonica-dev

```

On the project's root folder, run:
```bash
go mod tidy
```

NOTE: must have `go` version `1.21.0` or greater installed

Before running the `daemon` application, make sure to have InfluxDB running:
```bash
docker run \
      -p 8086:8086 \
      -v ./db:/var/lib/influxdb2 \
      influxdb:latest
```

Run the `daemon` application from the project's root with:
```bash
go run ./daemon/cmd/main.go
```

TODO: fix `Dockerfile` so that the project can be run just via a simple `docker compose up`

## Daemon
Takes a screenshot every 10 seconds, runs OCR on it and stores the extracted text and the base64 value of the screenshot on InfluxDB.

## API
Exposes endpoints to search and visualize screenshots captured by the Daemon. 

## UI
Consumes data from the API and presents an interface to interact with it.