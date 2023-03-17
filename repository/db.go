package repository

import (
	"database/sql"
	"github.com/AlmasOrazgaliev/assignment3/model"
	"gorm.io/gorm"
	"strings"
)

type Repository struct {
	DB *gorm.DB
}

func NewDB(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) GetBooks() (*[]model.Book, error) {
	rows, err := r.DB.Find(&model.Book{}).Rows()
	if err != nil {
		return nil, err
	}
	return rowIterator(rows)
}

func (r *Repository) GetById(id int) (*model.Book, error) {
	res := r.DB.First(&model.Book{}, id).Row()
	var book model.Book
	err := res.Scan(&book.Id, &book.Title, &book.Description, &book.Cost)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
func (r *Repository) CreateBook(book *model.Book) error {
	res := r.DB.Model(&model.Book{}).Create(&book)
	return res.Error
}

func (r *Repository) UpdateBook(book *model.Book, updatedBook *model.Book) error {
	res := r.DB.Model(&book).Updates(updatedBook)
	res.Save(&book)
	return res.Error
}

func (r *Repository) DeleteBook(book *model.Book) error {
	res := r.DB.Delete(&book, book.Id)
	return res.Error
}

func (r *Repository) SelectByTitle(title string) (*[]model.Book, error) {
	rows, err := r.DB.Model(&model.Book{}).Where("LOWER(title) LIKE ?", "%"+title+"%").Rows()
	if err != nil {
		return nil, err
	}
	return rowIterator(rows)
}

func (r *Repository) OrderBy(order string) (*[]model.Book, error) {
	rows, err := r.DB.Model(&model.Book{}).Order("cost " + strings.ToUpper(order)).Rows()
	if err != nil {
		return nil, err
	}
	return rowIterator(rows)
}

func rowIterator(rows *sql.Rows) (*[]model.Book, error) {
	var books []model.Book
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.Id, &book.Title, &book.Description, &book.Cost)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return &books, nil
}
