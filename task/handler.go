package task

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"todo-list/db"
	user "todo-list/user"
)

type Handler struct{}

func NewTaskHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var userID = r.Context().Value("userID").(uint)

	var currentUser user.User
	userRes := db.Db.First(&currentUser, userID)

	fmt.Println(currentUser)
	if userRes.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("User not found")
		return
	}

	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)

	result := db.Db.Create(&task)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result.Error)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid task ID")
		return
	}

	var task Task
	result := db.Db.First(&task, taskID)

	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task not found")
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	taskID, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid task ID")
		return
	}

	var task Task

	result := db.Db.Delete(&task, taskID)

	if result.Error != nil || result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task not found")
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode("Task Deleted Successfully")
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	taskID, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid task ID")
		return
	}

	var task Task
	result := db.Db.First(&task, taskID)

	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task not found")
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&task)

	db.Db.Save(&task)

	json.NewEncoder(w).Encode(task)
}

func (h *Handler) MarkTaskAs(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	taskID, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Task ID")
		return
	}

	var task Task
	result := db.Db.First(&task, taskID)

	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task not found")
		return
	}

	var completed bool
	_ = json.NewDecoder(r.Body).Decode(&completed)

	task.Completed = completed
	db.Db.Save(&task)

	json.NewEncoder(w).Encode(task)

}

func (h *Handler) updatePriority() {

}
