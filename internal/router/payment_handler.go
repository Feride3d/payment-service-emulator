package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Feride3d/payment-service-emulator/internal/storage"
	"github.com/Feride3d/payment-service-emulator/internal/storage/models"
)

type paymentHandler struct {
	store storage.Storage
}

// конструктор для пэйментхэндлера, принимает стор и возвращает стор
func newPaymentHandler(store storage.Storage) paymentHandler {
	return paymentHandler{
		store: store,
	}
}

func (h paymentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handleCreatePayment(w, r)
		/* 	case http.MethodPut:
		h.handlePaymentStatus(w, r) */
		/* 	case http.MethodDelete:
		h.handleDeletePayment(w, r) */
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h paymentHandler) handleCreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	jsDecoder := json.NewDecoder(r.Body)
	err := jsDecoder.Decode(&payment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed to decode body: %v", err)
		return
	}

	respPayment, err := h.store.CreatePayment(r.Context(), payment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to create payment in store: %v", err)
		return
	}

	writeJSON(w, &respPayment)
}

/* func (h paymentHandler) handlePaymentStatus(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	jsDecoder := json.NewDecoder(r.Body)
	err := jsDecoder.Decode(&payment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed to decode body: %v", err)
		return
	}

	respPayment, err := h.store.PaymentStatus(r.Context(), payment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to get status of payment from store: %v", err)
		return
	}

	writeJSON(w, &respPayment)
} */

/* func (h paymentHandler) handleDeletePayment(w http.ResponseWriter, r *http.Request) {
	id, _ := shiftPath(r.URL.Path)
	if id == "" {
		http.NotFound(w, r)
		return
	}

	err := h.store.DeletePayment(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed delete payment from store: %v", err)
		return
	}
}
*/
