package controllers

import (
	errorHandler "Simple-API/controllers/error"
	model "Simple-API/domain/model"
	services "Simple-API/services"

	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type AlbumController struct {
	albumService *services.AlbumService
}

func NewAlbumController() *AlbumController {
	service, err := services.NewService()
	if err != nil {
		log.Fatal(err)
	}

	return &AlbumController{
		albumService: service,
	}
}

func (controller *AlbumController) StartHandling() {
	http.HandleFunc("/albums", controller.handleBasicURI)
	http.HandleFunc("/albums/", controller.handleURIWithId)
}

func (controller *AlbumController) handleBasicURI(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Print(http.MethodGet)
		controller.findAll(w, r)
	case http.MethodPost:
		log.Print(http.MethodPost)
		controller.insert(w, r)
	default:
		errorHandler.HandleNotAllowedError(w)
	}
}

func (controller *AlbumController) handleURIWithId(w http.ResponseWriter, r *http.Request) {
	URIParts := strings.Split(r.RequestURI, "/")
	if len(URIParts) != 3 {
		errorHandler.HandleInvalidURIPathError(w)
		return
	}
	id := URIParts[2]

	switch r.Method {
	case http.MethodGet:
		log.Print(http.MethodGet)
		controller.findById(w, r, id)
	default:
		errorHandler.HandleNotAllowedError(w)
	}
}

func (controller *AlbumController) findAll(w http.ResponseWriter, r *http.Request) {
	albums, err := controller.albumService.FindAll()

	if err != nil {
		errorHandler.HandleInternalError(w, err)
		return
	}

	jsonBytes, err := json.Marshal(albums)
	if err != nil {
		errorHandler.HandleParsingJsonError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (controller *AlbumController) insert(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		errorHandler.HandleParsingJsonError(w, err)
		return
	}

	var contentType = r.Header.Get("Content-Type")
	if contentType != "application/json" {
		errorHandler.HandleUsupportedMediaTypeError(w)
		return
	}

	var album model.Album
	err = json.Unmarshal(bodyBytes, &album)

	if err != nil {
		errorHandler.HandleParsingJsonError(w, err)
		return
	}

	id, err := controller.albumService.Insert(album)

	if err != nil {
		errorHandler.HandleInternalError(w, err)
		return
	}

	findNewAlbumURL := r.Host + r.RequestURI + "/" + id
	w.Header().Add("Location", findNewAlbumURL)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (controller *AlbumController) findById(w http.ResponseWriter, r *http.Request, id string) {

	album, err := controller.albumService.FindById(id)

	if err != nil {
		errorHandler.HandleNotFoundTypeError(w, err)
		return
	}

	jsonBytes, err := json.Marshal(album)
	if err != nil {
		errorHandler.HandleParsingJsonError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
