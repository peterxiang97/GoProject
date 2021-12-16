package swagger

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestCalculateLoan(t *testing.T) {

	t.Run("Return Interest Only", func(t *testing.T) {

		Details := CalculateloanBody{
			LoanAmount:  int32(350000),
			LoanType:    "Interest Only",
			PaymentFrequency: "Monthly",
			InterestRate: float64(0.035),
			LoanTerm: int32(3),
	}

		body, _ := json.Marshal(Details)

		req:= httptest.NewRequest("POST", "/calculate-loan", bytes.NewBuffer(body))
        response := httptest.NewRecorder()

       	CalculateLoan(response, req)
		requestBodyBytes, _ := ioutil.ReadAll(response.Body)
		var  CalculateLoan1 LoanRepayments
		json.Unmarshal(requestBodyBytes, &CalculateLoan1)
		got := CalculateLoan1.MonthlyRepayments
        want := int32(20)

        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }
    })

	t.Run("Verify validation for invalid body - Negative", func(t *testing.T) {

		postBody := []byte("test")

		responseBody := bytes.NewBuffer(postBody)

		req := httptest.NewRequest("POST", "/calculate-loan", responseBody)
		w := httptest.NewRecorder()
		CalculateLoan(w, req)
		requestBodyBytes, err := ioutil.ReadAll(w.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var modelError ModelError
		json.Unmarshal(requestBodyBytes, &modelError) //nolint

		if result, _ := strconv.Atoi(modelError.Code); result != http.StatusBadRequest {
			t.Errorf("got: %s want: %d", modelError.Code, http.StatusBadRequest)
		}

	})
}