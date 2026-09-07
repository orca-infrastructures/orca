package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/orca-infrastructures/orca/internal/auth"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portEnv := os.Getenv("PORT")
	domainEnv := os.Getenv("DOMAIN")

	mux := http.NewServeMux()
	mux.HandleFunc("/login", auth.Login)

	log.Fatal(http.ListenAndServe(domainEnv+portEnv, mux))
}
