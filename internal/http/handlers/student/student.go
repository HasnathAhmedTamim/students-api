package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/hasnathahmedtamim/students-api/internal/storage"
	"github.com/hasnathahmedtamim/students-api/internal/types"
	"github.com/hasnathahmedtamim/students-api/internal/utils/response"
)

// New returns a http.HandlerFunc that creates a new student
func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating a student")
		// w.Write([]byte("Welcome to Students API!!"))

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("request body is empty")))
			return
		}

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
		}

		// Request validation

		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJSON(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)
		slog.Info("Student created successfully", slog.String("userId", fmt.Sprint(lastId)))
		if err != nil {
			response.WriteJSON(w, http.StatusInternalServerError, err)

			return
		}

		response.WriteJSON(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

// GetById returns a http.HandlerFunc that gets a student by id
func GetById(storage storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")

		slog.Info("Getting a student by id", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("error parsing id", slog.String("id", id))
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		student, err := storage.GetStudentById(intId)
		if err != nil {
			slog.Error("error getting student by id", slog.String("id", id))
			response.WriteJSON(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		response.WriteJSON(w, http.StatusOK, student)
	}
}

// GetAll returns a http.HandlerFunc that gets all students
func GetAll(storage storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Getting all students")
		students, err := storage.GetAllStudents()
		if err != nil {
			slog.Error("error getting all students", slog.String("error", err.Error()))
			response.WriteJSON(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		response.WriteJSON(w, http.StatusOK, students)
	}
}

// ...existing code...

// UpdateById returns a http.HandlerFunc that updates a student by id
func UpdateById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("Updating a student by id", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("error parsing id", slog.String("id", id))
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		var student types.Student
		err = json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("request body is empty")))
			return
		}

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// Request validation
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJSON(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		err = storage.UpdateStudent(intId, student.Name, student.Email, student.Age)
		if err != nil {
			slog.Error("error updating student", slog.String("id", id))
			response.WriteJSON(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		slog.Info("Student updated successfully", slog.String("id", id))
		response.WriteJSON(w, http.StatusOK, map[string]string{"success": "OK"})
	}
}

// DeleteById returns a http.HandlerFunc that deletes a student by id
func DeleteById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("Deleting a student by id", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("error parsing id", slog.String("id", id))
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		err = storage.DeleteStudent(intId)
		if err != nil {
			slog.Error("error deleting student", slog.String("id", id))
			response.WriteJSON(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		slog.Info("Student deleted successfully", slog.String("id", id))
		response.WriteJSON(w, http.StatusOK, map[string]string{"success": "OK"})
	}
}
