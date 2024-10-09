package hash

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

func Sh3CrpytoString(data string) string {
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
