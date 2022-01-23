package services

import (
	repositories "Simple-API/database/repositories"
	model "Simple-API/domain/model"
)

type AlbumService struct {
	albumRepository *repositories.AlbumRepository
}

func NewService() (*AlbumService, error) {
	repository, err := repositories.NewRepository()
	if err != nil {
		return &AlbumService{}, err
	}
	return &AlbumService{
		albumRepository: repository,
	}, nil
}

func (service *AlbumService) FindAll() ([]model.Album, error) {
	service.albumRepository.Lock()
	defer service.albumRepository.Unlock()
	albums, err := service.albumRepository.FindAll()
	return albums, err
}

func (service *AlbumService) Insert(album model.Album) (string, error) {
	service.albumRepository.Lock()
	defer service.albumRepository.Unlock()
	id, err := service.albumRepository.Insert(album)
	return id, err
}

func (service *AlbumService) FindById(id string) (model.Album, error) {
	service.albumRepository.Lock()
	defer service.albumRepository.Unlock()
	album, err := service.albumRepository.FindById(id)
	return album, err
}

func (service *AlbumService) UpdateById(album model.Album, id string) error {
	service.albumRepository.Lock()
	defer service.albumRepository.Unlock()
	err := service.albumRepository.UpdateById(album, id)
	return err
}

func (service *AlbumService) DeleteById(id string) error {
	service.albumRepository.Lock()
	defer service.albumRepository.Unlock()
	err := service.albumRepository.DeleteById(id)
	return err
}
