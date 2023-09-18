package main

import (
	"log"

	"github.com/akshayur0404/go-aws-s3/pkg/config"
	"github.com/akshayur0404/go-aws-s3/pkg/di"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cant load config")
	}
	di.InitializeAPI(cfg)

}
