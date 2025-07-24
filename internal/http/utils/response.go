package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

type WriteResponseModel struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

// the func writerespons is created to set the header status code and return the response
// there are two ways of creating dynmaic data - any and interface

func WriteResponse(w http.ResponseWriter, status int, data interface{}) error {

	// TODO:
	// set header
	// set status code
	// send response

	// content type is a header just to tell what type of data we are sending
	// application/json means we are sending data of type json..
	w.Header().Set("Content-Type", "application/json") // Done

	// writeHeaders are used to make the status codes
	w.WriteHeader(status) // Done

	// encode the data
	return json.NewEncoder(w).Encode(data)
}

// to send via object model..
func GeneralError(err error) WriteResponseModel {
	return WriteResponseModel{
		Status: StatusError,
		Error:  err.Error(),
	}
}

// next step is to work with validators
func CheckValidation(err validator.ValidationErrors) WriteResponseModel {
	var validtionFieldErrors []string
	for _, e := range err {
		switch e.ActualTag() { // actualTag error gives the actual tag
		case "required":
			validtionFieldErrors = append(validtionFieldErrors, fmt.Sprintf("The field %s is required", e.Field()))
		default:
			validtionFieldErrors = append(validtionFieldErrors, fmt.Sprintf("The field %s is invalid", e.Field()))
		}
	}

	return WriteResponseModel{
		Status: StatusError,
		Error:  strings.Join(validtionFieldErrors, ", "),
	}
}
