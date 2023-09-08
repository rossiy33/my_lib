package keychain

type CredentialManager interface {
	Set(name string, secret []byte) error
	Get(name string) ([]byte, error)
}
