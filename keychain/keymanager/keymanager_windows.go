//go:build windows
// +build windows

package keymanager

import (
	"github.com/danieljoos/wincred"
	"github.com/rossiy33/my_lib/keychain/lib"
)

type KeyManager struct{}

func (w KeyManager) Set(name string, secret []byte) error {
	cred := wincred.NewGenericCredential(name)
	cred.CredentialBlob = []byte(lib.EncB64(secret))
	return cred.Write()
}

func (w KeyManager) Get(name string) ([]byte, error) {
	cred, err := wincred.GetGenericCredential(name)
	if err != nil {
		return nil, err
	}
	return lib.DecB64(string(cred.CredentialBlob))
}
