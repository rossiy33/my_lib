//go:build !darwin && !windows
// +build !darwin,!windows

package keymanager

type KeyManager struct{}

func (d KeyManager) Set(name string, secret []byte) error {
	// ここに実装を追加するか、エラーを返します。
	return fmt.Errorf("unsupported OS")
}

func (d KeyManager) Get(name string) ([]byte, error) {
	return "", fmt.Errorf("unsupported OS")
}
