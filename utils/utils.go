package utils

import (
	"encoding/json"
	"io/ioutil"
	"my-book-api/models"

	"os"
)

var Books[] models.Book

const filePath ="/Users/samithadilshan/Documents/Randoli_BookProject/Book-store/data/books.json"


func LoadBooks() error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
			if os.IsNotExist(err) {
					Books = []models.Book{}
					return SaveBooks()
			}
			return err
	}
	return json.Unmarshal(data, &Books)
}

func SaveBooks() error {
	data, err := json.MarshalIndent(Books, "", "  ")
	if err != nil {
			return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}