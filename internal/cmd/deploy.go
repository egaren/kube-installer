package cmd

import (
	"github.com/egaren/kube-installer/internal/kube"
	"github.com/spf13/cobra"
	"log"
	"os"
)

type DeployCmd struct {
	command  *cobra.Command
	flavour  string // short -f ENV KUBE_INSTALLER_FLAVOUR
	version  string // short -v ENV KUBE_INSTALLER_VERSION
	nodeAddr string // short -n ENV KUBE_INSTALLER_NODEADDR
}

func NewDeployCmd() *DeployCmd {
	c := &DeployCmd{
		command: &cobra.Command{
			Use:   "deploy",
			Short: "kubernetes cluster deployment",
		},
	}
	c.command.Flags().StringVarP(&c.flavour, "flavour", "f", os.Getenv("KUBE_INSTALLER_FLAVOUR"), "flavor of kubernetes to deploy, env KUBE_INSTALLER_FLAVOUR")
	c.command.Flags().StringVarP(&c.version, "version", "v", os.Getenv("KUBE_INSTALLER_VERSION"), "version of kubernetes to deploy, env KUBE_INSTALLER_VERSION")
	c.command.Flags().StringVarP(&c.nodeAddr, "nodeip", "n", os.Getenv("KUBE_INSTALLER_NODEADDR"), "nodeaddr where to deploy kubernetes, env KUBE_INSTALLER_NODEADDR")
	c.command.Run = c.Run
	return c
}

func (c *DeployCmd) Run(commnd *cobra.Command, args []string) {
	switch c.flavour {
	case "k3s":
		if err := kube.DeployK3sCluster(c.version, c.nodeAddr); err != nil {
			log.Fatal(err)
		}

	}
}
