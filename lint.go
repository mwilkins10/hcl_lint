package main

import (
	"log"
	"github.com/hashicorp/hcl/hclsimple"
)

type Config struct {
	LogLevel string `hcl:"log_level"`
}

func main() {
	var config Config
	err := hclsimple.DecodeFile("/home/mwilkins/deploy-tool/infrastructure/almc-aws-inf/dev/almc/commonservices/account.hcl", nil, &config)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	log.Printf("Configuration is %#v", config)
}
