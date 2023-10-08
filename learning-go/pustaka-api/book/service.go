package book

import "time"

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BooksInput) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s service) Create(bookRequest BooksInput) (Book, error) {
	var book = Book{
		Title:       bookRequest.Title,
		Price:       int(bookRequest.Price),
		Description: bookRequest.Description,
		Rating:      int(bookRequest.Rating),
		Discount:    int(bookRequest.Discount),
		UpdateAt:    time.Now(),
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}
