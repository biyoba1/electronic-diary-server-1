package repository

import (
	diary "github.com/biyoba1/electronic-diary"
	"github.com/jmoiron/sqlx"
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

type Repository struct {
	Student
	Subject
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Student: NewStudentPostgres(db),
		Subject: NewSubjectPostgres(db),
	}
}
