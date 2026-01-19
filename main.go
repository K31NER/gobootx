package main

import (
	"log"

	"github.com/K31NER/gobootx/cmd"
	_ "github.com/K31NER/gobootx/internal/clean"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}