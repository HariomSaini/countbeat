package main

import (
	"os"

	"github.com/HariomSaini/countbeat/cmd"
	"github.com/elastic/beats/filebeat/beater"
	"github.com/elastic/beats/libbeat/cmd/instance"

	_ "github.com/HariomSaini/countbeat/include"
)

var RootCmd = cmd.GenRootCmdWithSettings(beater.New, instance.Settings{Name: "countbeat"})

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
