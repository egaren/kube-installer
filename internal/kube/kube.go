package kube

import (
	"github.com/egaren/kube-installer/pkg/k3s"
)

// DeployK3sCluster - deploys k3s cluster
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
	// Step 5 - Reload Daemon
	if v {
		if err = k.ReloadDaemon(); err != nil {
			return err
		}
	}
	// Step 6 - Start service
	if v {
		if err = k.StartService(); err != nil {
			return err
		}
	}
	// Step 7 - Enable Service
	if v {
		if err = k.EnableService(); err != nil {
			return err
		}
	}
	// Step 8 - Cleanup
	if v {
		if err = k.Remove(true); err != nil {
			return err
		}
	}
	return nil
}

// UpgradeK3sCluster - upgrades K3s cluster
func UpgradeK3sCluster(version string) error {
	// Step 1 - Get packages
	k := k3s.NewK3sBuilder(version, "")
	err := k.GetBinary()
	if err != nil {
		return err
	}
	// Step 2 - Validate packages
	v, err := k.ValidateBinary()
	if err != nil {
		return err
	}
	// Step 3 - Stop Service
	if v {
		if err = k.StopService(); err != nil {
			return err
		}
	}
	// Step 3 - Upgrades
	if v {
		err = k.InstallBinary()
		if err != nil {
			return err
		}
	}
	// Step 4 - Starts service
	if v {
		if err = k.StartService(); err != nil {
			return err
		}
	}
	// Step 5 - Cleanup
	if v {
		if err = k.Remove(true); err != nil {
			return err
		}
	}
	return nil
}

// Uninstall - uninstalls k3s cluster
func Uninstall(u bool) error {
	k := k3s.NewK3sBuilder("", "")
	if u {
		// Step 1 - stop service
		if err := k.StopService(); err != nil {
			return err
		}
		// Step 2 - cleanup
		if err := k.Remove(false); err != nil {
			return err
		}
	}
	return nil
}
