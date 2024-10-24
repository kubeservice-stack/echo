/*
Copyright 2024 The KubeService-Stack Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"github.com/gin-gonic/gin"
	"github.com/oklog/run"

	_ "github.com/kubeservice-stack/echo/docs"
	"github.com/kubeservice-stack/echo/internal/goruntime"
	_ "github.com/kubeservice-stack/echo/pkg/favicon"
	_ "github.com/kubeservice-stack/echo/pkg/health"
	_ "github.com/kubeservice-stack/echo/pkg/metrics"
	"github.com/kubeservice-stack/echo/pkg/routers"
	"github.com/kubeservice-stack/echo/pkg/version"

	logging "github.com/kubeservice-stack/common/pkg/logger"
	_ "github.com/kubeservice-stack/common/pkg/metrics"
)

const (
	ServerName           = "echo"
	DefaultMemlimitRatio = 0.85
)

var (
	mainLogger = logging.GetLogger("cmd", ServerName)
	printVer   bool
	printShort bool
)

func main() {
	app := kingpin.New(ServerName, "http server name")
	listenAddress := app.Flag(
		"listen-address",
		"address on which to expose metrics (disabled when empty)").
		String()

	app.Flag("version", "Prints current version.").Default("false").BoolVar(&printVer)
	app.Flag("short-version", "Print just the version number.").Default("false").BoolVar(&printShort)

	if _, err := app.Parse(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stdout, err)
		mainLogger.Error(fmt.Sprint("parse args err: ", err))
		os.Exit(2)
	}
	if printShort || printVer {
		Print(os.Stdout, ServerName)
		os.Exit(0)
	}

	if *listenAddress == "" {
		mainLogger.Error("--listen-address is empty")
		os.Exit(2)
	}

	mainLogger.Info("Starting server")

	goruntime.SetMaxProcs(mainLogger)
	goruntime.SetMemLimit(mainLogger, DefaultMemlimitRatio)

	var (
		g           run.Group
		ctx, cancel = context.WithCancel(context.Background())
	)

	defer cancel()

	r := gin.New()
	router.Router(r)
	srv := http.Server{
		Addr:              *listenAddress,
		WriteTimeout:      time.Second * 600,
		ReadHeaderTimeout: time.Second * 60,
		ReadTimeout:       time.Second * 60,
		IdleTimeout:       time.Second * 60,
		Handler:           r,
		MaxHeaderBytes:    1 << 20,
	}

	g.Add(func() error {
		mainLogger.Info("Starting web server", logging.Any("listenAddress", *listenAddress))
		return srv.ListenAndServe()
	}, func(error) {
		srv.Close()
	})

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	g.Add(func() error {
		select {
		case <-term:
			mainLogger.Info("Received SIGTERM, exiting gracefully...")
		case <-ctx.Done():
		}

		return nil
	}, func(error) {})

	if err := g.Run(); err != nil {
		mainLogger.Error("Failed to run", logging.Error(err))
		os.Exit(1)
	}
}

// Print version information to a given out writer.
func Print(out io.Writer, program string) {
	if printShort {
		fmt.Fprint(out, version.BuildContext())
		return
	}
	if printVer {
		fmt.Fprint(out, version.Print(program))
	}
}
