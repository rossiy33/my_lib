//go:build darwin
// +build darwin

package keymanager

import (
	"keychain/lib"
	"os/exec"
	"strings"
)

type KeyManager struct{}

func (k KeyManager) Set(name string, secret []byte) error {
	cmd := exec.Command("security", "add-generic-password", "-a", lib.GetUserName(), "-s", name, "-w", lib.EncB64(secret))
	return cmd.Run()
}

func (k KeyManager) Get(name string) ([]byte, error) {
	cmd := exec.Command("security", "find-generic-password", "-a", lib.GetUserName(), "-s", name, "-w")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	b64 := strings.TrimSpace(string(output))
	return lib.DecB64(b64)
}
