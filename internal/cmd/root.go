package cmd

import (
	"github.com/egaren/kube-installer/core"
	"github.com/spf13/cobra"
)

type RootCmd struct {
	Command *cobra.Command
}

func NewRootCmd() *RootCmd {
	c := &RootCmd{
		Command: &cobra.Command{
			Use:     "kube-installer",
			Short:   "Kube installer to deploy easier kubernetes clusters",
			Long:    ``,
			Version: core.Version,
		},
	}
	c.Command.SetVersionTemplate("kubeinstaller version: {{.Version}}\n")
	cobra.OnInitialize(c.initConfig)
	return c
}

func (c *RootCmd) initConfig() {

}
