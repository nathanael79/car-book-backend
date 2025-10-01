package service

import (
	"book-car/pkg/pagination"
	"book-car/repository"
)

type Service[T any] struct {
	repository *repository.Repo[T]
}

func NewService[T any](repository *repository.Repo[T]) *Service[T] {
	return &Service[T]{repository: repository}
}

func (s *Service[T]) FindOneByID(ID string) (*T, error) {
	result, err := s.repository.FindOneByID(ID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service[T]) FindAll(pagination *pagination.Pagination) (*[]T, error) {
	result, err := s.repository.FindAll(pagination)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service[T]) Create(request *T) (*T, error) {

	result, err := s.repository.Create(request)

	if err != nil {
		return nil, err
	}

	return result, nil
}
