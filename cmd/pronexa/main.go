// cmd/pronexa/main.go
package main

import (
	"flag"
	"log"
	"os"

	"pronexa/internal/pronexa"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := pronexa.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
