package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"pm2-exporter/pm2"
	"pm2-exporter/settings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "pm2-exporter"
	app.Flags = settings.NewContext()
	app.Action = run
	app.Version = "v1.1.0"

	err := app.Run(os.Args)
	if err != nil {
		log.Printf("app run error")
		log.Fatal(err)
	}
}

func run(ctx *cli.Context) {
	go pm2.GetPm2Info()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "/metrics")
	})

	http.Handle("/metrics", promhttp.Handler())
	log.Printf("starting exporter with port %v", 10100)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", settings.Port), nil))
}
