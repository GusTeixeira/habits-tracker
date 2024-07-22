package responses

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

type ErrorMessage struct {
	Error string `json:"error" docs:"description:Descricao do error"`
}

type RespostaPaginada struct {
	Pagina    int         `json:"pagina"`
	Paginas   int         `json:"total_paginas"`
	Registros int         `json:"total_registros"`
	Dados     interface{} `json:"dados"`
}

func HTTPError(w http.ResponseWriter, message string, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorMessage{
		Error: message,
	})
}

func HTTPSuccessPaginate(w http.ResponseWriter, dados interface{}, page, items, registros int) {
	var resposta = RespostaPaginada{
		Pagina:    page,
		Paginas:   int(math.Ceil(float64(registros) / float64(items))),
		Registros: registros,
		Dados:     dados,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resposta)
}

func HTTPSuccess(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)

}

func HTTPSuccessFile(w http.ResponseWriter, r *http.Request, filename string, status int) {
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "updated.xlsx"))
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.WriteHeader(status)
	http.ServeFile(w, r, filename)
}

func HTTPError2(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)

}
