package main

import (
	controllers "Simple-API/controllers"

	"log"
	"net/http"
)

func main() {
	controllers.NewAlbumController().StartHandling()
	//testando
	//teste
	//teste 2
	log.Fatal(http.ListenAndServe(":8080", nil))
}
