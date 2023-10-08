package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"
)

func main() {

	// CONNECT database dan automigrate
	dsn := "root:damai6543!@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db Connection Error")
	}

	fmt.Println("Connected")

	db.AutoMigrate(&book.Book{})

	//MENGGUNAKAN REPOSITORY
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	//bookRequest := book.BooksInput{
	//	Title:    "Harapan Bunga",
	//	Price:    25000,
	//	UpdateAt: time.Now(),
	//}
	//bookService.Create(bookRequest)

	//FINDALL
	//books, err := bookRepository.FindAll()
	//for _, book := range books {
	//	fmt.Println("Title:", book.Title)
	//}

	//FINDBY ID
	//book, err := bookRepository.FindByID(1)
	//fmt.Println("Title:", book.Title)

	//CREATE
	//book := book.Book{
	//	Title:       "100 Startup Digital",
	//	Price:       92000,
	//	Description: "Ini buku baru",
	//	Discount:    10,
	//	Rating:      4,
	//	UpdateAt:    time.Now(),
	//}
	//bookRepository.Create(book)

	// longchar itu biasanya lebih berat, lebih baik manual daripada automigrate

	// CREATE
	//book := book.Book{}
	//book.Title = "Atomic Habit"
	//book.Price = 105000
	//book.Discount = 5
	//book.Rating = 4
	//book.Description = "Ini buku ciptaan Lokalisasi"
	//book.UpdateAt = time.Now()
	//err = db.Create(&book).Error
	//if err != nil {
	//	fmt.Println("Error Created Book Record")
	//}

	// READ
	// SIAPKAN Data variabel dulu dan memunculkan query dengan debug
	//[2]SELECT 1 item
	//var book book.Book
	//[2]PARSING dengan ID
	//err = db.Debug().Last(&book).Error
	//[2]PARSING untuk primary key
	//err = db.Debug().Last(&book, 1).Error
	//if err != nil {
	//	fmt.Println("Show error finding a data")
	//}
	//fmt.Println("Title :", book.Title)
	//fmt.Println("Object data %v", book)
	//[3]Langkah selanjutnya
	//var books []book.Book
	//[3]Menggunakan FIND
	//err = db.Debug().Find(&books).Error
	//for _, b := range books {
	//	fmt.Println("title:", b.Title)
	//	fmt.Println("Object data %v", b)
	//}
	//[4]Menggunakan WHERE
	//err = db.Debug().Where("title = ?", "Atomic Habit").Find(&books).Error
	//for _, b := range books {
	//	fmt.Println("title:", b.Title)
	//	fmt.Println("Object data %v", b)
	//}
	//FIRST limit dan order id, TAKE hanya LIMIT 1, FIND tidak ada limit dan order
	// UPDATE
	//var book book.Book
	//err = db.Debug().Where("id=?", 1).Take(&book).Error
	//if err != nil {
	//	fmt.Println("Not Connected Data")
	//}
	//book.Title = "Man Tiger (Revised 2)"
	//err = db.Save(&book).Error
	//if err != nil {
	//	fmt.Println("Error Updating Data")
	//}

	// DELETE
	//var book book.Book
	//err = db.Debug().Where("id=?", 1).Take(&book).Error
	//if err != nil {
	//	fmt.Println("Not Connected Data")
	//}
	//err = db.Delete(&book).Error
	//if err != nil {
	//	fmt.Println("Error Deleting Data")
	//}

	// Inisiasi awal untuk menyiapkan router
	router := gin.Default()

	v1 := router.Group("/v1")

	// Ini root url
	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/hello", bookHandler.HelloHandler)
	v1.GET("/books/:id", bookHandler.BookHandler)
	v1.GET("/books/:id/:title", bookHandler.BooktitleHandler)
	v1.GET("/query", bookHandler.QueryHandler)
	v1.GET("/queryprice", bookHandler.QueryPriceHandler)
	v1.POST("/books", bookHandler.PostBooksHandler)

	router.Run("localhost:8080")
}

// Jika diluar package main, huruf funcnya harus kapital, sedangkan kalau kecil didalam satu package
// OBJECT DB ditempatkan dalam repository data (Layering)
// MAIN PROGRAM
// SERVICE (Berhubungan dengan Business Logic)
// REPOSITORY (Berhubungan dengan DB)
// DB
// MYSQL
