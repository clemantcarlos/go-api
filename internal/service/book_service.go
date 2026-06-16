package service

import (
	"errors"
	"go-api-test/internal/model"
	"go-api-test/internal/store"
)
type Logger interface {
	Log(msg string, error string)
}
type Service struct {
	store store.Store
	logger Logger
}

func New(s store.Store) *Service {
	return &Service{
		store: s,
		logger: nil,
	}
}

func (s *Service) GetAllBooks() ([]*model.Book, error){
	books, err := s.store.GetAll()
	if err != nil {
		s.logger.Log("El error es %v\n", err.Error())
		return nil, err
	}
	return books, nil
}
func (s *Service) GetBookById(id int) (*model.Book, error){
	return s.store.GetById(id)
}
func (s *Service) CreateBook(book model.Book) (*model.Book, error){
	if book.Titulo == "" {
		return nil, errors.New("Necesitamos el titulo")
	}
	
	return s.store.Create(&book)
}
func (s *Service) UpdateBook(id int, book model.Book) (*model.Book, error){
	if book.Titulo == "" {
		return nil, errors.New("Necesitamos el titulo")
	}

	return s.store.Update(id, &book)
}
func (s *Service) DeleteBook(id int) error{
	return s.store.Delete(id)
}