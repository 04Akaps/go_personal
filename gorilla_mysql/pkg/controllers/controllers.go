package controllers

import (
	"encoding/json"
	"fmt"
	model "gorilla_mysql/pkg/models"
	utils "gorilla_mysql/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook model.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := model.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	bookMethod := &model.Book{}
	utils.ParseBody(r, bookMethod) // 해당 부분을 통해서 body값과 Book Struct형태가 같은지를 검증
	b := bookMethod.CreateBook()

	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write((res))
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	// parseInt는 bookId라는 값을 파싱하는 역할을 수행 합니다.

	if err != nil {
		fmt.Println("error while parsing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bookDetails, _ := model.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &model.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book, db := model.GetBookById(ID)

	if book.Name != "" {
		book.Name = updateBook.Name
	}

	if book.Author != "" {
		book.Name = updateBook.Author
	}

	if book.Publication != "" {
		book.Publication = updateBook.Publication
	}

	db.Save(&book)
	res, _ := json.Marshal(book)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book := model.DeleteBookById(ID)
	res, _ := json.Marshal(book)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
