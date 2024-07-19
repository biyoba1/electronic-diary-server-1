package diary

import "errors"

type SubjectDiary struct {
	Id      int    `json:"id" db:"id"`
	UserId  int    `json:"userId" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	Mark    int    `json:"mark" binding:"required"`
	Date    string `json:"date" binding:"required"`
}

type UpdateSubjectInput struct {
	UserId  *int    `json:"userId"`
	Subject *string `json:"subject"`
	Mark    *int    `json:"mark"`
	Date    *string `json:"date"`
}

func (u UpdateSubjectInput) Validate() error {
	if u.Mark == nil && u.Date == nil && u.Subject == nil && u.UserId == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
