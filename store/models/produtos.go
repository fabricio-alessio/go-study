package models

import (
	"alessio.com/study/store/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {

	db := db.ConnectDB()
	query, err := db.Query("select * from product")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for query.Next() {
		var id int
		var nome string
		var descricao string
		var preco float64
		var quantidade int

		err := query.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CriarNovoPedido(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectDB()

	insert, err := db.Prepare("insert into product(name, description, price, quantity) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(productId int) {
	db := db.ConnectDB()

	delete, err := db.Prepare("delete from product where id = $1")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(productId)

	defer db.Close()
}
