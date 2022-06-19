package k3s

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// checkDirectory - checks if directory exists. if exists return true, not exists false
func checkDirsFiles(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// checkArchitecture - checks hardware architecture
func checkArchitecture() string {
	return runtime.GOARCH
}

// getBinaryName - returns binary file name struct
func getBinaryName(name string) string {
	arch := checkArchitecture()
	if arch == "amd64" {
		return fmt.Sprintf("%s", name)
	} else {
		return fmt.Sprintf("%s-%s", name, arch)
	}
}

// buildUrl - builds k3s download url
func buildUrl(version string, fileName string) (string, string) {
	arch := checkArchitecture()
	if arch == "amd64" && fileName == binaryName {
		return fmt.Sprintf("%s/%s/%s", sourceUrl, version, fileName), fmt.Sprintf("%s", fileName)
	} else if fileName == binaryName {
		return fmt.Sprintf("%s/%s/%s-%s", sourceUrl, version, fileName, arch), fmt.Sprintf("%s-%s", fileName, arch)
	} else {
		return fmt.Sprintf("%s/%s/%s-%s.txt", sourceUrl, version, fileName, arch), fmt.Sprintf("%s-%s.txt", fileName, arch)
	}
}

// buildFilePath - builds k3s binary filepath
func buildFilePath(fileName string, path string) string {
	arch := checkArchitecture()
	if arch == "amd64" && fileName == binaryName {
		return fmt.Sprintf("%s/%s", path, fileName)
	} else if fileName == binaryName {
		return fmt.Sprintf("%s/%s-%s", path, fileName, arch)
	} else {
		return fmt.Sprintf("%s/%s-%s.txt", path, fileName, arch)
	}
}

// getHashValue - get hash value from file
func getHashValue(filePath string, filter string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, filter) {
			str := strings.Split(line, " ")
			return str[0], nil
		}
	}
	if err = scanner.Err(); err != nil {
		return "", err
	}
	return "", nil
}

// getLinks - Links names
func getLinks() (k3sPath string, criCtlPath string, kubectlPath string, ctrPath string) {
	ctrPath = fmt.Sprintf("%s/%s", binDir, ctr)
	kubectlPath = fmt.Sprintf("%s/%s", binDir, kubectl)
	criCtlPath = fmt.Sprintf("%s/%s", binDir, criCtl)
	k3sPath = fmt.Sprintf("%s/%s", binDir, binaryName)
	return k3sPath, criCtlPath, kubectlPath, ctrPath
}
