package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

//	"log"
        "github.com/hashicorp/hcl/hclsimple"
	"github.com/mitchellh/colorstring"
)


type Config struct {
	IOMode  string        `hcl:"io_mode"`
	Service ServiceConfig `hcl:"service,block"`
}

type ServiceConfig struct {
	Protocol   string          `hcl:"protocol,label"`
	Type       string          `hcl:"type,label"`
	ListenAddr string          `hcl:"listen_addr"`
	Processes  []ProcessConfig `hcl:"process,block"`
}

type ProcessConfig struct {
	Type    string   `hcl:"type,label"`
	Command []string `hcl:"command"`
}

func main() {
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		search := arg
		if info, err := os.Stat(arg); err == nil && info.IsDir() {
			search = fmt.Sprintf("%s/*.tf", arg)
		}
		files, err := filepath.Glob(search)
		if err != nil {
			colorstring.Printf("[red]Error finding files: %s", err)
		}
		for _, filename := range files {
			fmt.Printf("Checking %s ... ", filename)
			file, err := ioutil.ReadFile(filename)
			if err != nil {
				colorstring.Printf("[red]Error reading file: %s\n", err)
				break
			}
//			#_, err = hclparse.NewParser(string(file))
                        var config Config
			err = hclsimple.DecodeFile(filename, nil, &config)
			if err != nil {
				colorstring.Printf("[red]Error parsing file: %s\n", err)
				break
			}
			colorstring.Printf("[green]OK!\n")
		}
	}

}
