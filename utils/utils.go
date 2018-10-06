package utils

import "encoding/hex"

// BytesToHex converts from an array of bytes to a hex encoded string
func BytesToHex(bytesArray []byte) string {
	r := "0x"
	h := hex.EncodeToString(bytesArray)
	r = r + h
	return r
}

// HexToBytes converts from a hex string into an array of bytes
func HexToBytes(h string) ([]byte, error) {
	b, err := hex.DecodeString(h[2:])
	return b, err
}
