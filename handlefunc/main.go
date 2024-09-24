package handlefunc

import (
	"net/http"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func GetOpenApi(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "api/openapi.yaml")
}
