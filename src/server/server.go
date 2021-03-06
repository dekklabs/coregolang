package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dekklabs/apirest/src/draw"
	"github.com/dekklabs/apirest/src/router"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Running enciende el servidor
func Running() {
	r := mux.NewRouter()

	router.Routers(r)

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	draw.Drawing(port)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	handler := c.Handler(r)

	err := http.ListenAndServe(":"+port, handler)

	if err != nil {
		fmt.Println("Error servidor", err)
	}
}
