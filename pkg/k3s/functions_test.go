package k3s

import (
	"fmt"
	"testing"
)

func TestK3s_GetBinary(t *testing.T) {

	k := NewK3sBuilder("v1.24.1+k3s1", "192.168.0.2")
	if err := k.GetBinary(); err != nil {
		t.Fail()
	}
}

func TestK3s_ValidateBinary(t *testing.T) {
	k := NewK3sBuilder("v1.24.1+k3s1", "192.168.0.2")
	res, err := k.ValidateBinary()
	if err != nil {
		t.Fail()
	}
	fmt.Println(res)
}

func TestK3s_InstallService(t *testing.T) {
	k := NewK3sBuilder("v1.24.1+k3s1", "192.168.0.2")
	if err := k.InstallService(); err != nil {
		fmt.Println(err)
	}
}

func TestK3s_MakeLinks(t *testing.T) {
	k := NewK3sBuilder("v1.24.1+k3s1", "192.168.0.2")
	if err := k.MakeLinks(); err != nil {
		fmt.Println(err)
	}
}
