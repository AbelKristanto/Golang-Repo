package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pustaka-api/book"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (handler *bookHandler) RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Budi",
		"age":  10,
	})
}

func (handler *bookHandler) HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"content": "Informasi Project",
		"bio":     "hi There!",
	})
}

func (handler *bookHandler) BookHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (handler *bookHandler) BooktitleHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")
	ctx.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func (handler *bookHandler) PostBooksHandler(ctx *gin.Context) {
	// title dan price
	var bookInput book.BooksInput

	err := ctx.ShouldBindJSON(&bookInput)
	//if err != nil {
	//	errorMessages := []string{}
	//	for _, e := range err.(validator.ValidationErrors) {
	//		errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
	//		// Menampilkan lebih dari satu
	//		errorMessages = append(errorMessages, errorMessage)
	//		// ======
	//		//ctx.JSON(http.StatusBadRequest, errorMessage)
	//		//fmt.Println(err)
	//		// biar tidak mengeksuksi dibawahnya
	//		//return
	//	}
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"errors": errorMessages,
	//	})
	//	return
	//}
	bookRequest, err := handler.bookService.Create(bookInput)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookRequest,
		//"author": bookInput.Author,
	})
}

// query?title=bumi manusia (Untuk proses query data dengan type)
func (handler *bookHandler) QueryHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
	})
}
func (handler *bookHandler) QueryPriceHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	price := ctx.Query("price")
	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

// BEFORE
//func RootHandler(ctx *gin.Context) {
//	ctx.JSON(http.StatusOK, gin.H{
//		"name": "Budi",
//		"age":  10,
//	})
//}
//
//func HelloHandler(ctx *gin.Context) {
//	ctx.JSON(http.StatusOK, gin.H{
//		"content": "Informasi Project",
//		"bio":     "hi There!",
//	})
//}
//
//func BookHandler(ctx *gin.Context) {
//	id := ctx.Param("id")
//	ctx.JSON(http.StatusOK, gin.H{
//		"id": id,
//	})
//}
//
//func BooktitleHandler(ctx *gin.Context) {
//	id := ctx.Param("id")
//	title := ctx.Param("title")
//	ctx.JSON(http.StatusOK, gin.H{
//		"id":    id,
//		"title": title,
//	})
//}
//
//func PostBooksHandler(ctx *gin.Context) {
//	// title dan price
//	var bookInput book.BooksInput
//
//	err := ctx.ShouldBindJSON(&bookInput)
//	if err != nil {
//		errorMessages := []string{}
//		for _, e := range err.(validator.ValidationErrors) {
//			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
//			// Menampilkan lebih dari satu
//			errorMessages = append(errorMessages, errorMessage)
//			// ======
//			//ctx.JSON(http.StatusBadRequest, errorMessage)
//			//fmt.Println(err)
//			// biar tidak mengeksuksi dibawahnya
//			//return
//		}
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"errors": errorMessages,
//		})
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{
//		"title": bookInput.Title,
//		"price": bookInput.Price,
//		//"author": bookInput.Author,
//	})
//}
//
//// query?title=bumi manusia (Untuk proses query data dengan type)
//func QueryHandler(ctx *gin.Context) {
//	title := ctx.Query("title")
//	ctx.JSON(http.StatusOK, gin.H{
//		"title": title,
//	})
//}
//func QueryPriceHandler(ctx *gin.Context) {
//	title := ctx.Query("title")
//	price := ctx.Query("price")
//	ctx.JSON(http.StatusOK, gin.H{
//		"title": title,
//		"price": price,
//	})
//}
