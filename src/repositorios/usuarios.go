package repositorios

import (
	"database/sql"
	"social-media-go/src/modelos"
)

// Usuarios representa um repositorio de usuarios
type usuarios struct {
	db *sql.DB
}

func (u *usuarios) Criar(usuario modelos.Usuario) {
	panic("unimplemented")
}

// NovoRepositorioDeUsuarios cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Criar insere um usuario no banco de dados
func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertID()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido)
}
