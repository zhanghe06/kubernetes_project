package utils

import (
	"encoding/base64"
)

func Base64Endoce(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func Base64Dedoce(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}
