package common

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func Md5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func FormatTime(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}
