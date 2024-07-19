package diary

import "errors"

type StudentDiary struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" binding:"required"`
	LastName   string `json:"LastName" binding:"required"`
	Patronymic string `json:"patronymic" binding:"required"`
}

type UpdateStudentInput struct {
	Name       *string `json:"name"`
	LastName   *string `json:"LastName"`
	Patronymic *string `json:"patronymic"`
}

func (i UpdateStudentInput) Validate() error {
	if i.Name == nil && i.LastName == nil && i.Patronymic == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
