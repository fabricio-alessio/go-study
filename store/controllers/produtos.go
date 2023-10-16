package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"alessio.com/study/store/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	produtos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CriarNovoPedido(nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProdutoStr := r.URL.Query().Get("id")
	idProduto, err := strconv.Atoi(idProdutoStr)
	if err != nil {
		panic("Impossible to convert product id using" + idProdutoStr)
	}

	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProdutoStr := r.URL.Query().Get("id")
	idProduto, err := strconv.Atoi(idProdutoStr)
	if err != nil {
		panic("Impossible to convert product id using" + idProdutoStr)
	}
	produto := models.RetornaProduto(idProduto)

	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			panic(err.Error())
		}
		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			panic(err.Error())
		}
		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			panic(err.Error())
		}

		models.EditProduto(idConvertido, nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301)
}
