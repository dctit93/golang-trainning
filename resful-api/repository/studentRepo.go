package repository

import (
	"database/sql"
	"resful-api/models"
)

func GetStudents(db *sql.DB, student models.Student, students []models.Student) ([]models.Student, error) {
	rows, err := db.Query("select * from students")
	if err != nil {
		return []models.Student{}, err
	}
	for rows.Next() {
		err = rows.Scan(&student.Id, &student.Name, &student.Address)
		students = append(students, student)
	}
	if err != nil {
		return []models.Student{}, err
	}
	return students, err
}

func GetStudentsByID(db *sql.DB, student models.Student, id int) (models.Student, error) {
	rows := db.QueryRow("select * from students where id=$1", id)
	err := rows.Scan(&student.Id, &student.Name, &student.Address)
	return student, err
}

func AddStudent(db *sql.DB, student models.Student) (int, error) {
	err := db.QueryRow("insert into students (id, name, address) values($1, $2, $3) RETURNING id;",
		student.Id, student.Name, student.Address).Scan(&student.Id)

	if err != nil {
		return 0, err
	}

	return student.Id, nil
}
