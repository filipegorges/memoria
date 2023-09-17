# memoria
Generate a searchable visual history of your activity on your computer.

This project is both a tool for personal use as well as a means to practice a few buzzwordy shenanigans.

## Daemon
Takes a screenshot every 10 seconds, runs OCR on it and stores the extracted text and the base64 value of the screenshot on InfluxDB.

## API
Exposes endpoints to search and visualize screenshots captured by the Daemon. 

## UI
Consumes data from the API and presents an interface to interact with it.