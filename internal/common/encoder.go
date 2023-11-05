package common

import (
	"encoding/hex"
	"hash/fnv"
)

// Encode string in 32bit hash
func EncodeString(s string) string {
	hash := fnv.New32a()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}
