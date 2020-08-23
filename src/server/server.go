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

	if port != "" {
		port = "5000"
	}

	draw.Drawing()

	handler := cors.Default().Handler(r)

	err := http.ListenAndServe(":"+port, handler)

	if err != nil {
		fmt.Println("Error ", err)
	}
}
