package main

import (
	"flag"
	"time"

	"github.com/liabio/sparrow-tools/pkg/rest"
)

func main() {

	var (
		timeout                 time.Duration
		listenPort string
		interPath string
	)

	flag.StringVar(&listenPort, "listen-port", ":10000", "The port binds to.")
	flag.StringVar(&interPath, "inter-path", "/inter/timeout/simulate", "interface uri for simulate timeout")
	flag.DurationVar(&timeout, "timeout",  time.Duration(20) * time.Second, "interface timeout")
	flag.Parse()

	opt := &rest.RunOptions{
		Timeout: timeout,
		InterPath: interPath,
		ListenPort: listenPort,
	}

	rest.Run(opt)
}
