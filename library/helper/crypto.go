package helper

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Bytes(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
