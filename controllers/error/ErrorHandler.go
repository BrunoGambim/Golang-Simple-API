package errorHandler

import (
	"log"
	"net/http"
)

func HandleNotAllowedError(w http.ResponseWriter) {
	log.Print("Not Allowed")
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func HandleInvalidURIPathError(w http.ResponseWriter) {
	log.Print("Invalid URI path")
	w.WriteHeader(http.StatusBadRequest)
}

func HandleInternalError(w http.ResponseWriter, err error) {
	log.Print("Internal error")
	log.Print(err.Error())
	w.WriteHeader(http.StatusInternalServerError)
}

func HandleParsingJsonError(w http.ResponseWriter, err error) {
	log.Print("Parsing json error")
	log.Print(err.Error())
	w.WriteHeader(http.StatusInternalServerError)
}

func HandleUsupportedMediaTypeError(w http.ResponseWriter) {
	log.Print("Usupported media type")
	w.WriteHeader(http.StatusUnsupportedMediaType)
}

func HandleNotFoundTypeError(w http.ResponseWriter, err error) {
	log.Print("Not found")
	log.Print(err.Error())
	w.WriteHeader(http.StatusNotFound)
}
