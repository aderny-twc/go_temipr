package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/aderny-twc/go_temipr/application"
)

func main() {
	app := application.New()

	signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
