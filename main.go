package main

import (
	"fmt"
	"net/http"

	"swagger-sample/controller"
	_ "swagger-sample/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Swagger Example API (HTTPS)
// @version 1.0
// @description Example API server using only net/http and HTTPS.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8443
// @BasePath /

func main() {
	mux := http.NewServeMux()

	// Redirect root "/" to Swagger UI
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusFound)
	})

	// Swagger endpoint
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("https://localhost:8080/swagger/doc.json"),
	))

	// API endpoints
	mux.HandleFunc("/sum", controller.SumHandler)
	mux.HandleFunc("/div", controller.DivisionHandler)

	// Configure CORS

	fmt.Println("🚀 HTTPS server running at https://localhost:8080")
	err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", mux)
	if err != nil {
		panic(err)
	}
}
