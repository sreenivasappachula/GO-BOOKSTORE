package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-gorm/gorm/pkg/models"
	"github.com/go-gorm/gorm/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book


func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks:=models.GetAllBooks()
	res, _:json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	bookId:=vars["bookId"]
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("error while parsing")
	}
	bookDetails, _:=models.GetBookById(ID)
	res,_:=json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook:=&modesl.Book{}
	utils.ParseBody(r, CreateBook)
	b:=CreateBook.CreateBook()
	res, _:=json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	bookId:=vars["bookId"]
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("error while parsing")
	}
	book:=models.DeleteBook(ID)
	res, _:=json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOk)
	w.Write(res)
}


func UpdateBook(w http.ResponseWriter, r *http.Request){
	var UpdateBook =&models.Book{}
	utils.ParseBody(r, UpdateBook)
	vars:=mux.Vars(r)
	bookId:vars["bookId"]
	ID, err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("error while parsing")
	}
	bookDeetails, db:=models.GetBookById(ID)

	if updateBook.Name!=""{
		bookDeetails.Name= UpdateBook.Name
	}
	if updateBook.Author!=""{
		bookDeetails.Author= UpdateBook.Author
	}

	if updateBook.Publication!=""{
		bookDeetails.Publication= UpdateBook.Publication
	}

	db.Save(&bookDeetails)
	res,_:=json.Marshal(bookDeetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOk)
	w.Write(res)
}

