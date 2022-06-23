package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Feride3d/payment-service-emulator/internal/router"
	"github.com/stretchr/testify/require"
)

func TestNotFound(t *testing.T) {
	r := router.New(nil)
	srv := httptest.NewServer(r.RootHandler())
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/unknown")
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

/*
func TestCreatePayment(t *testing.T) {
	r := router.New(memstorage.New)
	srv := httptest.NewServer(r.RootHandler())
	defer srv.Close()

	respPayment := createPayment(t, srv.URL, "test payment")
	checkPaymentList(t, srv.URL, respPayment)
}

func TestCancelPayment(t *testing.T) {
	r := router.New(memstorage.New())
	srv := httptest.NewServer(r.RootHandler())
	defer srv.Close()

	respPayment := createPayment(t, srv.URL, "test payment")
	deletePayment(t, srv.URL, respPayment.ID)
	checkPaymentList(t, srv.URL)
}

func createPayment(t *testing.T, url string, text string) models.Payment {
	payment := models.Payment{
		Text: text,
	}
	data, err := json.Marshal(&payment)
	require.NoError(t, err)

	resp, err := http.Post(url+"/payment", "application/json", bytes.NewReader(data))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var respPayment models.Payment
	jsDecoder := json.NewDecoder(resp.Body)
	err = jsDecoder.Decode(&respPayment)
	require.NoError(t, err)

	return respPayment
} */

func deletePayment(t *testing.T, url string, id string) {
	req, err := http.NewRequest(http.MethodDelete, url+"/payment/"+id, nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
