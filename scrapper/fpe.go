package scrapper

import (
	"encoding/hex"
	"github.com/capitalone/fpe/ff1"
)

func f(in rune) bool {
	return !((in >= '0' && in <= '9') || (in >= 'a' && in <= 'z') || (in >= 'A' && in <= 'Z'))
}

func Tokenize(inData string, key []byte, tweak []byte) (string, error) {

	if key == nil {
		key, _ = hex.DecodeString("123456ABCDEF7890123456ABCDEF7890")
	}

	if tweak == nil {
		tweak, _ = hex.DecodeString("123456ABCDEF7890")
	}

	strout := []string{}

	acc := make([]byte, len(inData))

	out := make([]byte, len(inData))

	cip, _ := ff1.NewCipher(36, 1000, key, tweak)

	offset := 0
	offsetacc := 0

	for l, i := range inData {

		if l == len(inData)-1 {
			if !f(i) {
				acc[offsetacc] = byte(i)
				offsetacc++
			}
		}

		if f(i) || l == len(inData)-1 {
			if offsetacc == 0 {
				strout = append(strout)
				out[offset] = byte(i)
				offset++

			} else if offsetacc <= 2 {
				copy(out[offset:], acc[:offsetacc])
				offset += offsetacc
				if f(i) {
					out[offset] = byte(i)
					offset++
				}
			} else {

				o, err := cip.Encrypt(string(acc[:offsetacc]))
				if err != nil {
					o = string(acc[:offsetacc])
				}
				copy(out[offset:], o)
				offset += len(o)
				if f(i) {
					out[offset] = byte(i)
					offset++
				}
			}
			offsetacc = 0
		} else {
			acc[offsetacc] = byte(i)
			offsetacc++
		}
	}

	return string(out), nil
}

func TokenizeInt(inData string, key []byte, tweak []byte) (string, error) {
	if key == nil {
		key, _ = hex.DecodeString("123456ABCDEF7890123456ABCDEF7890")
	}

	if tweak == nil {
		tweak, _ = hex.DecodeString("123456ABCDEF7890")
	}
	cip, _ := ff1.NewCipher(10, 1000, key, tweak)
	return cip.Encrypt(inData)
}

func Detokenize(inData string, key []byte, tweak []byte) (string, error) {

	if key == nil {
		key, _ = hex.DecodeString("123456ABCDEF7890123456ABCDEF7890")
	}

	if tweak == nil {
		tweak, _ = hex.DecodeString("123456ABCDEF7890")
	}

	strout := []string{}

	acc := make([]byte, len(inData))

	out := make([]byte, len(inData))

	cip, _ := ff1.NewCipher(36, 1000, key, tweak)

	offset := 0
	offsetacc := 0

	for l, i := range inData {

		if l == len(inData)-1 {
			if !f(i) {
				acc[offsetacc] = byte(i)
				offsetacc++
			}
		}

		if f(i) || l == len(inData)-1 {
			if offsetacc == 0 {
				strout = append(strout)
				out[offset] = byte(i)
				offset++

			} else if offsetacc <= 2 {
				copy(out[offset:], acc[:offsetacc])
				offset += offsetacc
				if f(i) {
					out[offset] = byte(i)
					offset++
				}
			} else {

				o, err := cip.Decrypt(string(acc[:offsetacc]))
				if err != nil {
					o = string(acc[:offsetacc])
				}
				copy(out[offset:], o)
				offset += len(o)
				if f(i) {
					out[offset] = byte(i)
					offset++
				}
			}
			offsetacc = 0
		} else {
			acc[offsetacc] = byte(i)
			offsetacc++
		}
	}

	return string(out), nil
}

func DetokenizeInt(inData string, key []byte, tweak []byte) (string, error) {
	if key == nil {
		key, _ = hex.DecodeString("123456ABCDEF7890123456ABCDEF7890")
	}

	if tweak == nil {
		tweak, _ = hex.DecodeString("123456ABCDEF7890")
	}
	cip, _ := ff1.NewCipher(10, 1000, key, tweak)
	return cip.Decrypt(inData)
}
