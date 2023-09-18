package main

import (
	"github.com/filipegorges/memoria/api"
)

func main() {
	srv := api.NewAPI()
	if err := srv.Run(":8080"); err != nil {
		panic(err)
	}
	defer srv.Close()
}
