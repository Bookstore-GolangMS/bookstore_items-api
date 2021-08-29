package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Bookstore-GolangMS/bookstore_items-api/domain/items"
	"github.com/Bookstore-GolangMS/bookstore_items-api/services"
	httputils "github.com/Bookstore-GolangMS/bookstore_items-api/utils/http_utils"
	"github.com/Bookstore-GolangMS/bookstore_oauth-go/oauth"
	"github.com/Bookstore-GolangMS/bookstore_utils-go/errors"
)

var (
	ItensController itensControllerInterface = &itensController{}
)

type itensControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type itensController struct {
}

func (i *itensController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// httputils.RespondJsonError(w, errors.NewRestError(err.Message, err.Status, err.Error, []))
		// return
	}

	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		httputils.RespondJsonError(w, errors.NewUnauthorizedError("unable to retrieve user information from given access_token"))
		return
	}

	var itemRequest items.Item

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputils.RespondJsonError(w, errors.NewBadRequestError("Error trying to read body json"))
		return
	}

	defer r.Body.Close()

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		httputils.RespondJsonError(w, errors.NewBadRequestError("Error trying to read body json"))
		return
	}

	itemRequest.Seller = sellerId

	result, errorService := services.ItemsService.Create(itemRequest)
	if errorService != nil {
		httputils.RespondJsonError(w, errorService)
		return
	}

	httputils.RespondJson(w, http.StatusCreated, result)
}

func Get(w http.ResponseWriter, r *http.Request) {

}
