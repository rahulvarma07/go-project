package response

import (
	"encoding/json"
	"net/http"
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
