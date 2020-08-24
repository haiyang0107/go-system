package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Enc(str []byte) (st string) {
	hashValue := md5.New()
	hashValue.Write(str)
	return hex.EncodeToString(hashValue.Sum(nil))
}
