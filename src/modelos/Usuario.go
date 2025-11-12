package modelos

import "time"

// Usuario representa um usuario utilizando a rede social
type Usuario struct {
	ID uint64 `json:"id,omitempty"`
	Nome string `json:"nome,omitempty"`
	Nick string `nick:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}