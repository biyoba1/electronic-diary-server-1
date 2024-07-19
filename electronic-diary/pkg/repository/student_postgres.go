package repository

import (
	"fmt"
	diary "github.com/biyoba1/electronic-diary"
	"github.com/jmoiron/sqlx"
	"strings"
)

type StudentPostgres struct {
	db *sqlx.DB
}

func NewStudentPostgres(db *sqlx.DB) *StudentPostgres {
	return &StudentPostgres{db: db}
}

func (r *StudentPostgres) CreateStudent(student diary.StudentDiary) (int, error) {
	var id int
	createStudentQuery := fmt.Sprintf("INSERT INTO %s (name, last_name, patronymic) VALUES ($1, $2, $3) RETURNING id", studentsTable)

	row := r.db.QueryRow(createStudentQuery, student.Name, student.LastName, student.Patronymic)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *StudentPostgres) GetAll() ([]diary.StudentDiary, error) {
	var students []diary.StudentDiary
	getAllStudentQuery := fmt.Sprintf("SELECT id, name, last_name, patronymic FROM %s", studentsTable)
	rows, err := r.db.Query(getAllStudentQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var student diary.StudentDiary
		err := rows.Scan(&student.Id, &student.Name, &student.LastName, &student.Patronymic)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (r *StudentPostgres) GetById(id int) (diary.StudentDiary, error) {
	var student diary.StudentDiary
	getByIdStudentQuery := fmt.Sprintf("SELECT id, name, last_name, patronymic FROM %s  WHERE id=$1", studentsTable)
	row := r.db.QueryRow(getByIdStudentQuery, id)
	err := row.Scan(&student.Id, &student.Name, &student.LastName, &student.Patronymic)
	if err != nil {
		return student, err
	}
	return student, nil
}

func (r *StudentPostgres) Update(id int, input diary.UpdateStudentInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.LastName != nil {
		setValues = append(setValues, fmt.Sprintf("last_name=$%d", argId))
		args = append(args, *input.LastName)
		argId++
	}

	if input.Patronymic != nil {
		setValues = append(setValues, fmt.Sprintf("patronymic=$%d", argId))
		args = append(args, *input.Patronymic)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	fmt.Println(setQuery)

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", studentsTable, setQuery, argId)
	args = append(args, id)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *StudentPostgres) Delete(id int) error {
	deleteStudentQuery := fmt.Sprintf("DELETE FROM %s WHERE id=$1 ", studentsTable)
	_, err := r.db.Exec(deleteStudentQuery, id)
	return err
}
