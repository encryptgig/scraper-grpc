package model

type Key struct {
	Model   `json:"model"`
	KeyByte []byte `json:"key_byte,omitempty"`
	IV      []byte `json:"iv,omitempty"`
	Algo    string `json:"algo,omitempty"`
	Size    int    `json:"size,omitempty"`
	Meta    string `json:"meta,omitempty"`
	OwnerID string `json:"owner_id"`
}
