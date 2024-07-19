package service

import (
	diary "github.com/biyoba1/electronic-diary"
	"github.com/biyoba1/electronic-diary/pkg/repository"
)

type SubjectService struct {
	repo repository.Subject
}

func NewSubjectService(repo repository.Subject) *SubjectService {
	return &SubjectService{repo: repo}
}

func (s *SubjectService) Create(subject diary.SubjectDiary) (int, error) {
	return s.repo.Create(subject)
}

func (s *SubjectService) GetAll(userId int) ([]diary.StudentDiary, []diary.SubjectDiary) {
	return s.repo.GetAll(userId)
}

func (s *SubjectService) Update(subjectId int, input diary.UpdateSubjectInput) error {
	return s.repo.Update(subjectId, input)
}

func (s *SubjectService) Delete(subjectId int) error {
	return s.repo.Delete(subjectId)
}
