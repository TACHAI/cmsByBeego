package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(text []byte)string  {
	ctx:=md5.New()
	ctx.Write(text)
	return hex.EncodeToString(ctx.Sum(nil))
}