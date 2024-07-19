package service

import (
	diary "github.com/biyoba1/electronic-diary"
	"github.com/biyoba1/electronic-diary/pkg/repository"
)

type Student interface {
	CreateStudent(student diary.StudentDiary) (int, error)
	GetAll() ([]diary.StudentDiary, error)
	GetById(id int) (diary.StudentDiary, error)
	Update(id int, input diary.UpdateStudentInput) error
	Delete(id int) error
}

type Subject interface {
	Create(subject diary.SubjectDiary) (int, error)
	GetAll(userId int) ([]diary.StudentDiary, []diary.SubjectDiary)
	Update(subjectId int, input diary.UpdateSubjectInput) error
	Delete(subjectId int) error
}

type Service struct {
	Student
	Subject
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Student: NewStudentService(repos.Student),
		Subject: NewSubjectService(repos.Subject),
	}
}
