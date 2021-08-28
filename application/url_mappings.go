package app

import (
	"net/http"

	"github.com/Bookstore-GolangMS/bookstore_items-api/controllers"
)

func mapUrl() {
	router.HandleFunc("/items", controllers.ItensController.Create).Methods(http.MethodPost)
}
