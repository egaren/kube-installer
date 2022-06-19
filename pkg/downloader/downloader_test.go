package downloader

import (
	"fmt"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	if err := DownloadFile("k3s-arm64", "https://github.com/k3s-io/k3s/releases/download/v1.24.1+k3s1/k3s-arm64"); err != nil {
		fmt.Println(err)
	}
}
