package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/kamiertop/gin-wrap/internal/router"
	cfg "github.com/kamiertop/gin-wrap/pkg/config"
	"github.com/kamiertop/gin-wrap/pkg/log"
)

func initialize() (err error) {
	flag.Parse()
	if err = cfg.Init(); err != nil {
		return err
	}

	log.Init()

	return nil
}

func main() {
	if err := initialize(); err != nil {
		fmt.Println("initialize failed, err: ", err)
		os.Exit(1)
	}
	engine := router.Init()

	server := &http.Server{
		Handler: engine,
		Addr:    ":8080",
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
