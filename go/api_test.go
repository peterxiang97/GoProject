package swagger

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestCalculateLoan(t *testing.T) {

	t.Run("Return Interest Only Check", func(t *testing.T) {

		Details := CalculateloanBody{
			LoanAmount:  int32(35000),
			LoanType:    "Interest Only",
			PaymentFrequency: "Monthly",
			InterestRate: float64(0.035),
			LoanTerm: int32(3),
	}

		Expected := LoanRepayments{
			MonthlyRepayments: int32(102),
			TotalInterestPayable:  int32(3675),
			AmountOwing: []LoanRepaymentsAmountOwing{
				LoanRepaymentsAmountOwing{
				
					Year: int32(0),
					Interest: int32(3675),
					Principal: int32(35000),
					Total: int32(38675),
				},
				LoanRepaymentsAmountOwing{
				
					Year: int32(1),
					Interest: int32(2450),
					Principal: int32(35000),
					Total: int32(37450),
				},
				LoanRepaymentsAmountOwing{
				
					Year: int32(2),
					Interest: int32(1125),
					Principal: int32(35000),
					Total: int32(36125),
				},
				LoanRepaymentsAmountOwing{
				
					Year: int32(3),
					Interest: int32(0),
					Principal: int32(35000),
					Total: int32(35000),
				},
			},
		}

		body, _ := json.Marshal(Details)

		req:= httptest.NewRequest("POST", "/calculate-loan", bytes.NewBuffer(body))
        response := httptest.NewRecorder()

       	CalculateLoan(response, req)
		requestBodyBytes, _ := ioutil.ReadAll(response.Body)
		var  CalculateLoan1 LoanRepayments
		json.Unmarshal(requestBodyBytes, &CalculateLoan1)  //nolint
		

        if CalculateLoan1.MonthlyRepayments != Expected.MonthlyRepayments {
            t.Errorf("wrong monthly repayments")
        }  
		if CalculateLoan1.TotalInterestPayable != Expected.TotalInterestPayable {
            t.Errorf("wrong total interest payable")
        }
		
		for i := 0; i <= 3; i++{
			if CalculateLoan1.AmountOwing[i] != Expected.AmountOwing[i]{
				t.Errorf("wrong amount Owing '%d'", i)
			}
			
		}
    })

	t.Run("Return Principal & Interest Check", func(t *testing.T) {

		Details := CalculateloanBody{
			LoanAmount:  int32(35000),
			LoanType:    "Principal & Interest",
			PaymentFrequency: "Monthly",
			InterestRate: float64(0.035),
			LoanTerm: int32(3),
	}

		Expected := LoanRepayments{
			MonthlyRepayments: int32(1026),
			TotalInterestPayable:  int32(1921),
			AmountOwing: []LoanRepaymentsAmountOwing{
				LoanRepaymentsAmountOwing{
				
					Year: int32(0),
					Interest: int32(1921),
					Principal: int32(35000),
					Total: int32(36921),
				},
				LoanRepaymentsAmountOwing{
				
					Year: int32(1),
					Interest: int32(875),
					Principal: int32(23739),
					Total: int32(24614),
				},
				LoanRepaymentsAmountOwing{
				
					Year: int32(2),
					Interest: int32(230),
					Principal: int32(12077),
					Total: int32(12307),
				},
				LoanRepaymentsAmountOwing{
				
					Year: int32(3),
					Interest: int32(0),
					Principal: int32(0),
					Total: int32(0),
				},
			},
		}

		body, _ := json.Marshal(Details)

		req:= httptest.NewRequest("POST", "/calculate-loan", bytes.NewBuffer(body))
        response := httptest.NewRecorder()

       	CalculateLoan(response, req)
		requestBodyBytes, _ := ioutil.ReadAll(response.Body)
		var  CalculateLoan1 LoanRepayments
		json.Unmarshal(requestBodyBytes, &CalculateLoan1)  //nolint
		

        if CalculateLoan1.MonthlyRepayments != Expected.MonthlyRepayments {
            t.Errorf("wrong monthly repayments")
        }  
		if CalculateLoan1.TotalInterestPayable != Expected.TotalInterestPayable {
            t.Errorf("wrong total interest payable")
        }
		
		for i := 0; i <= 3; i++{
			if CalculateLoan1.AmountOwing[i] != Expected.AmountOwing[i]{
				t.Errorf("wrong amount Owing '%d'", i)
			}
			
		}
    })

	
}
