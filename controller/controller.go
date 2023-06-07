package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AlmasOrazgaliev/assignment3/model"
	"github.com/AlmasOrazgaliev/assignment3/repository"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

//var repo, err = repository.NewDB()

type Controller struct {
	repository *repository.Repository
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{
		repository: repository.NewDB(db),
	}
}

func (c *Controller) HandleBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		res, err := c.repository.GetBooks()

		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusOK, res)
	} else if r.Method == "POST" {
		book := model.Book{}
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			response(w, http.StatusBadRequest, err)
			return
		}
		err = c.repository.CreateBook(&book)
		if err != nil {
			response(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusCreated, nil)
	}
}

func (c *Controller) HandleBookById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		errResponse(w, http.StatusBadRequest, err)
		return
	}
	book, err := c.repository.GetById(id)
	if err != nil {
		errResponse(w, http.StatusBadRequest, err)
		return
	}
	if book == nil {
		errResponse(w, http.StatusNoContent, errors.New("no such book"))
	}
	if r.Method == "GET" {
		response(w, http.StatusOK, book)
	} else if r.Method == "DELETE" {
		err = c.repository.DeleteBook(book)
		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusNoContent, nil)
	} else if r.Method == "PUT" {
		var updatedBook model.Book
		err = json.NewDecoder(r.Body).Decode(&updatedBook)
		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		err = c.repository.UpdateBook(book, &updatedBook)
		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusCreated, nil)
	}
}

func (c *Controller) HandleSearch(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	fmt.Println(r.URL.Query())
	fmt.Println(title)
	books, err := c.repository.SelectByTitle(title)
	if len(*books) == 0 {
		err = errors.New("no such book")
	}
	if err != nil {
		errResponse(w, http.StatusInternalServerError, err)
		return
	}
	response(w, http.StatusOK, books)
}

func (c *Controller) HandleOrder(w http.ResponseWriter, r *http.Request) {
	order := mux.Vars(r)["order"]
	fmt.Println(order)
	books, err := c.repository.OrderBy(order)
	if err != nil {
		errResponse(w, http.StatusInternalServerError, err)
		return
	}
	response(w, http.StatusOK, books)
}

func errResponse(w http.ResponseWriter, code int, err error) {
	response(w, code, map[string]string{"error": err.Error()})
}

func response(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
