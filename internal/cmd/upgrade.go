package cmd

import (
	"github.com/egaren/kube-installer/internal/kube"
	"github.com/spf13/cobra"
	"log"
	"os"
)

type UpgradeCmd struct {
	command *cobra.Command
	flavour string // short -f ENV KUBE_INSTALLER_FLAVOUR
	version string // short -v ENV KUBE_INSTALLER_VERSION
}

func NewUpgradeCmd() *UpgradeCmd {
	c := &UpgradeCmd{
		command: &cobra.Command{
			Use:   "upgrade",
			Short: "kubernetes cluster upgrade",
		},
	}
	c.command.Flags().StringVarP(&c.flavour, "flavour", "f", os.Getenv("KUBE_INSTALLER_FLAVOUR"), "flavor of kubernetes to deploy, env KUBE_INSTALLER_FLAVOUR")
	c.command.Flags().StringVarP(&c.version, "version", "v", os.Getenv("KUBE_INSTALLER_VERSION"), "version of kubernetes to deploy, env KUBE_INSTALLER_VERSION")
	c.command.Run = c.Run
	return c
}

func (c *UpgradeCmd) Run(commnd *cobra.Command, args []string) {
	switch c.flavour {
	case "k3s":
		if err := kube.UpgradeK3sCluster(c.version); err != nil {
			log.Fatal(err)
		}

	}
}
