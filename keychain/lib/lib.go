package lib

import (
	"encoding/base64"
	"os/user"
)

func GetUserName() string {
	usr, _ := user.Current()
	return usr.Username
}

func EncB64(src []byte) string {
	return base64.RawStdEncoding.EncodeToString(src)
}

func DecB64(src string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(src)
}
