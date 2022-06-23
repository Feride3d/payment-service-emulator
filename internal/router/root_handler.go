package router

import (
	"net/http"

	"github.com/Feride3d/payment-service-emulator/internal/storage"
)

type rootHandler struct {
	paymentHandler paymentHandler
}

func newRootHandler(store storage.Storage) rootHandler {
	return rootHandler{
		paymentHandler: newPaymentHandler(store),
	}
}

func (h rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	// take the url from the request, take the path, cut off the tail and put everything that is left in the head
	head, r.URL.Path = shiftPath(r.URL.Path)

	switch head {
	case "payment":
		h.paymentHandler.ServeHTTP(w, r)

	default:
		http.NotFound(w, r)
	}
}
