package main

import (
	"github.com/egaren/kube-installer/internal/cmd"
	"log"
)

func main() {
	rootCmd := cmd.InitialiseRootCmd()
	if err := rootCmd.Command.Execute(); err != nil {
		log.Fatal(err)
	}
}
