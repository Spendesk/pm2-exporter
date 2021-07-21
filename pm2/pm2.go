package pm2

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"

	"pm2-exporter/settings"

	"github.com/prometheus/client_golang/prometheus"
)

type process struct {
	Pid    int    `json:"pid"`
	Name   string `json:"name"`
	PmID   int    `json:"pm_id"`
	Monit  monit  `json:"monit"`
	Pm2Env pm2Env `json:"pm2_env"`
}

type monit struct {
	Memory int     `json:"memory"`
	CPU    float32 `json:"cpu"`
}

type pm2Env struct {
	RestartTime int    `json:"restart_time"`
	Version     string `json:"version"`
	NodeVersion string `json:"node_version"`
}

var (
	memoryUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pm2_process_memory_usage",
			Help: "Memory Usage from PM2",
		},
		[]string{"pm_name", "pm_id", "version", "node_version"},
	)

	cpuUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pm2_process_cpu_usage",
			Help: "CPU Usage from PM2",
		},
		[]string{"pm_name", "pm_id", "version", "node_version"},
	)

	restartTime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "pm2_process_restart_time",
			Help: "Restart Count from PM2",
		},
		[]string{"pm_name", "pm_id", "version", "node_version"},
	)
)

func init() {
	prometheus.MustRegister(memoryUsage)
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(restartTime)
}

func GetPm2Info() {
	for {
		var list []process

		// exec pm2 command and get data
		out, err := exec.Command(settings.PM2Path, "jlist").Output()
		if err != nil {
			log.Printf("exec error")
			log.Println(err)
		}

		// transform exec output to struct
		err = json.Unmarshal(out, &list)
		if err != nil {
			log.Printf("json unmarshal error")
			log.Println(err)
		}

		fmt.Println(list)

		// parse data
		for _, process := range list {
			memoryUsage.WithLabelValues(
				process.Name,
				strconv.Itoa(process.PmID),
				process.Pm2Env.Version,
				process.Pm2Env.NodeVersion,
			).Set(float64(process.Monit.Memory))

			cpuUsage.WithLabelValues(
				process.Name,
				strconv.Itoa(process.PmID),
				process.Pm2Env.Version,
				process.Pm2Env.NodeVersion,
			).Set(float64(process.Monit.CPU))

			restartTime.WithLabelValues(
				process.Name,
				strconv.Itoa(process.PmID),
				process.Pm2Env.Version,
				process.Pm2Env.NodeVersion,
			).Set(float64(process.Pm2Env.RestartTime))
		}

		time.Sleep(time.Duration(settings.Refresh) * time.Second)
	}
}
