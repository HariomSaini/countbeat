package main

import (
	"os"

	"github.com/HariomSaini/countbeat/cmd"

	_ "github.com/HariomSaini/countbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
