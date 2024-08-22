package main

import (
	"log"

	"github.com/GusTeixeira/habits-tracker/runner"
)

func main() {
	if err := runner.Run("habits-tracker"); err != nil {
		log.Fatal(err)
	}
}
