package main

import (
	"context"
	"log"
	"time"
)

func main() {
	app := New()

	startCtx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}
}
