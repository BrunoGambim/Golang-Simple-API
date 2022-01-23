package main

import (
	controllers "Simple-API/controllers"

	"log"
	"net/http"
)

func main() {
	controllers.NewAlbumController().StartHandling()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
