package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            // obter os valores dos campos do formulário
            nome := r.FormValue("nome")
            email := r.FormValue("email")
            telefone := r.FormValue("telefone")

            // exibir os valores enviados
            fmt.Fprintf(w, "Nome: %s\n", nome)
            fmt.Fprintf(w, "E-mail: %s\n", email)
            fmt.Fprintf(w, "Telefone: %s\n", telefone)
        } else {
            // exibir o formulário
            http.ServeFile(w, r, "formulario.html")
        }
    })

    // iniciar o servidor
    fmt.Println("Servidor rodando na porta 8080...")
    http.ListenAndServe(":8080", nil)
}
