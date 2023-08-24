package gen

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

// GenerateRandString 汎用的なランダム文字列を生成する
func GenerateRandString(length int) (key []byte, err error) {
	key = make([]byte, length)
	_, err = rand.Read(key)
	return
}

// GenerateKey ランダム文字列を生成し、base64エンコードする
func GenerateKey(length int) (key string, err error) {
	str, err := GenerateRandString(length)
	key = base64.RawURLEncoding.EncodeToString(str)
	return
}

// GenerateCodeChallenge PKCE用のハッシュ値を生成する
func GenerateCodeChallenge(verifier string) (code string) {
	s256 := sha256.Sum256([]byte(verifier))
	code = base64.RawURLEncoding.EncodeToString(s256[:])
	return
}
