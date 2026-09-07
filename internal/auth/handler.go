package auth

import (
	"net/http"

	"github.com/orca-infrastructures/orca/internal/jsonhelper"
)

type LoginRequest struct {
	id       string
	password string
}

func Login(w http.ResponseWriter, r *http.Request) {
	loginRequest, err := jsonhelper.Decode[LoginRequest](r)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

}
