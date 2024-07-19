package service

import (
	diary "github.com/biyoba1/electronic-diary"
	"github.com/biyoba1/electronic-diary/pkg/repository"
)

type StudentService struct {
	repo repository.Student
}

func NewStudentService(repo repository.Student) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) CreateStudent(student diary.StudentDiary) (int, error) {
	return s.repo.CreateStudent(student)
}

func (s *StudentService) GetAll() ([]diary.StudentDiary, error) {
	return s.repo.GetAll()
}

func (s *StudentService) GetById(id int) (diary.StudentDiary, error) {
	return s.repo.GetById(id)
}

func (s *StudentService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *StudentService) Update(id int, input diary.UpdateStudentInput) error {
	return s.repo.Update(id, input)
}
