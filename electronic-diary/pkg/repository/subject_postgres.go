package repository

import (
	"fmt"
	diary "github.com/biyoba1/electronic-diary"
	"github.com/jmoiron/sqlx"
	"strings"
)

type SubjectPostgres struct {
	db *sqlx.DB
}

func NewSubjectPostgres(db *sqlx.DB) *SubjectPostgres {
	return &SubjectPostgres{db: db}
}

func (r *SubjectPostgres) Create(subject diary.SubjectDiary) (int, error) {
	var id int
	createSubjectQuery := fmt.Sprintf("INSERT INTO %s (user_id, subject, mark, date) VALUES ($1, $2, $3, $4) RETURNING id", subjectsTable)
	row := r.db.QueryRow(createSubjectQuery, subject.UserId, subject.Subject, subject.Mark, subject.Date)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *SubjectPostgres) GetAll(userId int) ([]diary.StudentDiary, []diary.SubjectDiary) {
	var students []diary.StudentDiary
	var subjects []diary.SubjectDiary

	getAllQuery := fmt.Sprintf("SELECT st.id, st.name, st.last_name, st.patronymic, subj.id, subj.user_id, subj.subject, subj.mark, subj.date FROM %s st INNER JOIN %s subj ON st.id = subj.user_id WHERE subj.user_id = $1", studentsTable, subjectsTable)
	rows, err := r.db.Query(getAllQuery, userId)
	if err != nil {
		return nil, nil
	}
	var student diary.StudentDiary
	for rows.Next() {
		var subject diary.SubjectDiary
		err := rows.Scan(&student.Id, &student.Name, &student.LastName, &student.Patronymic, &subject.Id, &subject.UserId, &subject.Subject, &subject.Mark, &subject.Date)
		if err != nil {
			return nil, nil
		}
		subjects = append(subjects, subject)
	}
	students = append(students, student)
	return students, subjects
}

func (r *SubjectPostgres) Update(subjectId int, input diary.UpdateSubjectInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.UserId != nil {
		setValues = append(setValues, fmt.Sprintf("user_id=$%d", argId))
		args = append(args, *input.UserId)
		argId++
	}

	if input.Subject != nil {
		setValues = append(setValues, fmt.Sprintf("subject=$%d", argId))
		args = append(args, *input.Subject)
		argId++
	}

	if input.Mark != nil {
		setValues = append(setValues, fmt.Sprintf("mark=$%d", argId))
		args = append(args, *input.Mark)
		argId++
	}

	if input.Date != nil {
		setValues = append(setValues, fmt.Sprintf("date=$%d", argId))
		args = append(args, *input.Date)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	fmt.Println(setQuery)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", subjectsTable, setQuery, argId)
	args = append(args, subjectId)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *SubjectPostgres) Delete(subjectId int) error {
	deleteSubjectQuery := fmt.Sprintf("DELETE FROM %s WHERE id=$1", subjectsTable)
	_, err := r.db.Exec(deleteSubjectQuery, subjectId)
	return err
}
