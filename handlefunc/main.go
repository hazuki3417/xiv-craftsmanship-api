package handlefunc

import (
	"log"
	"net/http"
	"os"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func GetOpenApi(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("api/openapi.yaml")

	if err != nil {
		log.Printf("failed to read openapi.yaml: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
