package models

import (
	"alura/loja/db"
)

type Produto struct {
	ID         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectBd()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConnectBd()

	deletarProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarProduto.Exec(id)
	defer db.Close()
}

func BuscaTodosOsProdutos() []Produto {
	// Obtenha a conexão com o banco de dados
	conexao := db.ConnectBd()
	defer conexao.Close()

	selectProdutos, err := conexao.Query("SELECT * FROM produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}
	defer selectProdutos.Close()

	var produtos []Produto

	for selectProdutos.Next() {
		var produto Produto
		err := selectProdutos.Scan(&produto.ID, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtos = append(produtos, produto)
	}
	return produtos
}

func EditaProduto(id string) Produto {
	db := db.ConnectBd()

	produtoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	produtoParaAtualizar := Produto{}

	for produtoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.ID = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}

	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectBd()

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
