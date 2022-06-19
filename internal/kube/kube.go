package kube

import (
	"github.com/egaren/kube-installer/pkg/k3s"
)

func DeployK3sCluster(version string, nodeAddr string) error {
	// Step 1 - Get packages
	k := k3s.NewK3sBuilder(version, nodeAddr)
	err := k.GetBinary()
	if err != nil {
		return err
	}
	// Step 2 - Install packages
	v, err := k.ValidateBinary()
	if v {
		err = k.InstallBinary()
		if err != nil {
			return err
		}
	}
	if v {
		err = k.MakeLinks()
		if err != nil {
			return err
		}
	}
	// Step 3 - Install services
	if v {
		err = k.InstallService()
		if err != nil {
			return err
		}
	}
	// Step 4 - Create necessary files and directories
	if v {
		err = k.CreateDirectories()
		if err != nil {
			return err
		}
		err = k.CreateFiles()
		if err != nil {
			return err
		}
	}
	// Step 5 - Start service
	if v {
		if err = k.StartService(); err != nil {
			return err
		}
	}

	return nil
}
