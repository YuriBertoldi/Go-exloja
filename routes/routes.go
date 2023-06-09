package routes

import (
	"net/http"

	controlller "github.com/YuriBertoldi/Go-exloja/Controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controlller.Index)
	http.HandleFunc("/newProduto", controlller.NewProduto)
	http.HandleFunc("/insertProduto", controlller.InsertProduto)
	http.HandleFunc("/deleteProduto", controlller.DeleteProduto)
	http.HandleFunc("/editProduto", controlller.EditProduto)
	http.HandleFunc("/updateProduto", controlller.UpdateProduto)
}
