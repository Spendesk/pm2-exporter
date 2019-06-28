package settings

import (
	"github.com/urfave/cli"
)

var (
	PM2Path string
	Refresh int64
	Port    int64
)

// NewContext - Get configuration from env vars or command parameters
func NewContext() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "pm2_path, pp",
			Value:       "pm2",
			Destination: &PM2Path,
			EnvVar:      "PM2_PATH",
			Usage:       "Path to PM2 command if not present in $PATH",
		},
		cli.Int64Flag{
			Name:        "refresh, r",
			Value:       30,
			Destination: &Refresh,
			EnvVar:      "REFRESH",
			Usage:       "PM2 status refresh interval",
		},
		cli.Int64Flag{
			Name:        "port, p",
			Value:       10100,
			Destination: &Port,
			EnvVar:      "PORT",
			Usage:       "Exporter port",
		},
	}
}
