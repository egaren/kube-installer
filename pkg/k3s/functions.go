package k3s

import (
	"fmt"
	"github.com/egaren/kube-installer/crypto"
	"github.com/egaren/kube-installer/pkg/downloader"
	"github.com/egaren/kube-installer/pkg/systemd"
	"io"
	"os"
	"strings"
	"text/template"
)

type K3s struct {
	version  string
	NodeAddr string
}

// NewK3sBuilder init
func NewK3sBuilder(version string, nodeAddr string) *K3s {
	k := K3s{
		version:  version,
		NodeAddr: nodeAddr,
	}
	return &k
}

// CreateDirectories - creates directories if not exists
func (k *K3s) CreateDirectories() error {
	if !checkDirsFiles(configDir) {
		if err := os.MkdirAll(configDir, executableMode); err != nil {
			return err
		}
	}
	if !checkDirsFiles(configYamlDir) {
		if err := os.MkdirAll(configYamlDir, executableMode); err != nil {
			return err
		}
	}
	return nil
}

// CreateFiles - creates configuration file if not exists
func (k *K3s) CreateFiles() error {
	if !checkDirsFiles(configFile) {
		f, err := os.Create(configFile)
		if err != nil {
			return err
		}
		if err = f.Close(); err != nil {
			return err
		}
	}
	return nil
}

// GetBinary - downloads binary
func (k *K3s) GetBinary() error {
	url, fileName := buildUrl(k.version, binaryName)
	err := downloader.DownloadFile("/tmp/"+fileName, url)
	if err != nil {
		return err
	}
	url, fileName = buildUrl(k.version, shaSumName)
	err = downloader.DownloadFile("/tmp/"+fileName, url)
	if err != nil {
		return err
	}
	return nil
}

// ValidateBinary - downloaded binary validation
func (k *K3s) ValidateBinary() (bool, error) {
	binName := getBinaryName(binaryName)
	shaSumFilePath := buildFilePath(shaSumName, downloadPath)
	binaryFilePath := buildFilePath(binName, downloadPath)
	sha256Hash, err := getHashValue(shaSumFilePath, binName)
	if err != nil {
		return false, err
	}
	sha256Sum, err := crypto.Sha256Sum(binaryFilePath)
	if err != nil {
		return false, err
	}
	if strings.Compare(sha256Sum, sha256Hash) == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

// InstallBinary - installs binary to proper place
func (k *K3s) InstallBinary() error {
	binarySourcePath := buildFilePath(binaryName, downloadPath)
	binaryDestinationPath := buildFilePath(binaryName, binDir)
	//source location of the binary
	sourceFile, err := os.Open(binarySourcePath)
	if err != nil {
		return err
	}
	//destination location of the binary
	destFile, err := os.Create(binaryDestinationPath)
	if err != nil {
		if err = sourceFile.Close(); err != nil {
			return err
		}
		return err
	}
	defer func() {
		if err = destFile.Close(); err != nil {
		}
	}()
	//copies binary to destination
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}
	//changing permission of the binary
	err = os.Chmod(binaryDestinationPath, executableMode)
	if err != nil {
		return err
	}
	return nil
}

// MakeLinks - makes links to binary
func (k *K3s) MakeLinks() error {
	k3sPath, criCtlPath, kubectlPath, ctrPath := getLinks()
	binaryDestinationPath := buildFilePath(binaryName, binDir)
	arch := checkArchitecture()
	if arch != "amd64" {
		err := os.Symlink(binaryDestinationPath, k3sPath)
		if err != nil {
			return err
		}
	}
	err := os.Symlink(binaryDestinationPath, criCtlPath)
	if err != nil {
		return err
	}
	err = os.Symlink(binaryDestinationPath, kubectlPath)
	if err != nil {
		return err
	}
	err = os.Symlink(binaryDestinationPath, ctrPath)
	if err != nil {
		return err
	}
	return nil
}

// InstallService - installs service file
func (k *K3s) InstallService() error {
	dest := fmt.Sprintf("%s/%s.service", systemDirectory, binaryName)
	s, err := os.Create(dest)
	if err != nil {
		return err
	}
	serviceTemplate := template.Must(template.New(svc).Parse(svc))
	err = serviceTemplate.Execute(s, k)
	if err != nil {
		return err
	}
	return nil
}

// StartService - starts k3s service
func (k *K3s) StartService() error {
	systemdClient := systemd.NewSystemdClient()
	//reload daemon
	if err := systemdClient.Reload(); err != nil {
		return err
	}
	unit := fmt.Sprintf("%s.service", binaryName)
	if err := systemdClient.Start(unit); err != nil {
		return err
	}
	if err := systemdClient.Enable(unit); err != nil {
		return err
	}
	return nil
}
