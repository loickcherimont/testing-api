package service

import (
	"github.com/loickcherimont/testing-api/model"
	"github.com/loickcherimont/testing-api/repository"
)

func GetAllTests() []model.Test {
	return repository.FindAll()
}

func GetTestById(id uint) model.Test {
	return repository.FindById(id)
}
