package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			username := r.FormValue("username")
			password := r.FormValue("email")
			if username == "" || password == "" {
				fmt.Fprintf(w, "Por favor, preencha todos os campos.")
				return
			}
			fmt.Fprintf(w, "Login bem-sucedido! Bem-vindo, %s.", username)
		}
		fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Formulário de Contato</title>
		  <link rel="stylesheet" href="Tela.css">
		</head>
		<body>
			<h1>Formulário de Contato</h1>
			<form id="formulario">
				<label for="nome">Nome:</label>
				<input type="text" id="nome" name="nome"><br>
		
				<label for="email">E-mail:</label>
				<input type="email" id="email" name="email"><br>
		
				<label for="telefone">Telefone:</label>
				<input type="tel" id="telefone" name="telefone"><br>
		
				<button type="submit">Enviar</button>
			</form>
		
			<script src="info.js"></script>
		</body>
		</html>
		
		`)
	})
	http.ListenAndServe(":8080", nil)
}
