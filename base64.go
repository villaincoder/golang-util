package util

import (
	"encoding/base64"
	"encoding/binary"
)

func Base64ToUint64(base64String string) (uint64Value uint64, err error) {
	bytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return
	}
	uint64Value = binary.BigEndian.Uint64(bytes)
	return
}

func Uint64ToBase64(uint64Value uint64) (base64String string) {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64Value)
	base64String = base64.StdEncoding.EncodeToString(buf)
	return
}

func Int64ToBase64(int64Value int64) (base64String string) {
	base64String = Uint64ToBase64(uint64(int64Value))
	return
}
