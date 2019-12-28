package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"resful-api/models"
	"resful-api/repository"
	"resful-api/utils"

	"strconv"

	"github.com/gorilla/mux"
)

var students []models.Student

func GetAllStudent(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student models.Student
		students, err := repository.GetStudents(db, student, students)
		if err != nil {
			utils.PreponseData(w, 400, err)
		}
		utils.PreponseData(w, 200, students)
	}
}
func GetStudentByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student models.Student
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		students, err := repository.GetStudentsByID(db, student, id)
		if err != nil {
			utils.PreponseData(w, 400, err)
		}
		utils.PreponseData(w, 200, students)
	}
}
func AddStudent(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student models.Student
		fmt.Println(&r.Body)
		fmt.Println(json.NewDecoder(r.Body).Decode(&student))
		json.NewDecoder(r.Body).Decode(&student)

		if string(student.Id) == "" || student.Name == "" || student.Address == "" {
			utils.PreponseData(w, 400, utils.SendWithMessage("Thiếu thông tin đầu vào"))
			return
		}

		id, err := repository.AddStudent(db, student)
		fmt.Println(id)
		if err != nil {
			utils.PreponseData(w, 400, err)
		}
		utils.PreponseData(w, 400, utils.SendWithMessage("Thêm thành công"))
		return
	}
}
