package cmd

import (
	"github.com/egaren/kube-installer/internal/kube"
	"github.com/spf13/cobra"
	"log"
	"os"
)

type UninstallCmd struct {
	command   *cobra.Command
	flavour   string // short -f ENV KUBE_INSTALLER_FLAVOUR
	uninstall bool
}

func NewUninstallCmd() *UninstallCmd {
	c := &UninstallCmd{
		command: &cobra.Command{
			Use:   "uninstall",
			Short: "kubernetes cluster uninstall",
		},
	}
	c.command.Flags().StringVarP(&c.flavour, "flavour", "f", os.Getenv("KUBE_INSTALLER_FLAVOUR"), "flavor of kubernetes to deploy, env KUBE_INSTALLER_FLAVOUR")
	c.command.Flags().BoolVarP(&c.uninstall, "uninstall", "u", true, "uninstall kubernetes cluster if true")
	c.command.Run = c.Run
	return c
}

func (c *UninstallCmd) Run(commnd *cobra.Command, args []string) {
	switch c.flavour {
	case "k3s":
		if err := kube.Uninstall(c.uninstall); err != nil {
			log.Fatal(err)
		}

	}
}
