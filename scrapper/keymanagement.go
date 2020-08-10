package scrapper

import (
	"crypto/aes"
	"encoding/hex"
	"errors"
	"github.com/encryptgig/scraper-grpc/scrapperpb"
	"github.com/google/uuid"
	"sync"
)

type Key struct {
	ID      string `json:"id,omitempty"`
	KeyByte []byte `json:"key_byte,omitempty"`
	IV      []byte `json:"iv,omitempty"`
	Algo    string `json:"algo,omitempty"`
	Size    int    `json:"size,omitempty"`
	Meta    string `json:"meta,omitempty"`
}

func (k *Key) ValidateKey() (err error) {
	switch k.Algo {
	case "aes":
		_, err = aes.NewCipher(k.KeyByte)
		return
	default:
		err = errors.New("unsupported algorithm")
	}
	return
}

type InputKeyParameters struct {
	Key       string      `json:"key"`
	IV        string      `json:"iv"`
	Algorithm string      `json:"algorithm"`
	Size      int         `json:"size"`
	Meta      interface{} `json:"meta"`
}

func (in *InputKeyParameters) ValidateKeyParams() (err error) {
	if len(in.Key) == 0 {
		err = errors.New("invalid key bytes in request")
	}
	if len(in.IV) == 0 {
		err = errors.New("invalid IV bytes in request")
	}
	if len(in.Algorithm) == 0 {
		err = errors.New("invalid algorithm in request")
	}
	return
}

type KeyMap map[string]*Key

var Keys = KeyMap{}

func (k KeyMap) Add(in *Key) {
	RwMutex.Lock()
	defer RwMutex.Unlock()
	k[in.ID] = in
}

func (k KeyMap) Get(id string) (*Key, bool) {
	RwMutex.RLock()
	defer RwMutex.RUnlock()
	v, ok := k[id]
	return v, ok
}

var RwMutex sync.RWMutex

func getKeyBytes(key string) ([]byte, error) {
	f, err := hex.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func init() {
	kb, _ := getKeyBytes("C1AB55FE70ACC6EEB61FCDB0DF48FF799C35E1C1B3835AB86203FF8582734D09")
	iv, _ := getKeyBytes("F162046AFC1F2302FE7D11E4629068F1")
	Keys["default_key"] = &Key{KeyByte: kb, IV: iv, ID: "default_key", Size: 256, Algo: "aes"}
}

func ConfigureKey(in *scrapperpb.Key) (kid string, err error) {

	kb, err := getKeyBytes(in.KeyBytes)
	if err != nil {
		err = errors.New("invalid key bytes in request")
		return
	}

	iv, err := getKeyBytes(in.KeyIv)
	if err != nil {
		err = errors.New("invalid IV bytes in request")
		return
	}

	id := in.ID

	if len(id) == 0 {
		id = uuid.New().String()
	}

	key := Key{KeyByte: kb, IV: iv, Algo: in.KeyAlgorithm, Size: int(in.KeySize), Meta: in.Meta, ID: id}

	err = key.ValidateKey()
	if err != nil {
		return
	}
	Keys.Add(&key)

	return id, nil

}

func ListKeys() *scrapperpb.KeyList {

	RwMutex.RLock()
	defer RwMutex.RUnlock()

	out := &scrapperpb.KeyList{}

	for _, v := range Keys {
		kl := &scrapperpb.KeyInfo{}
		kl.ID = v.ID
		kl.Meta = v.Meta
		kl.KeySize = int32(v.Size)
		kl.KeyAlgorithm = v.Algo
		out.Key = append(out.Key, kl)
	}
	return out
}

func DeleteKey(id string) string {

	RwMutex.Lock()
	defer RwMutex.Unlock()

	in := struct {
		ID string `json:"id"`
	}{}

	delete(Keys, in.ID)

	return id

}
