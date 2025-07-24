package students

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rahulvarma07/goo_backend/internal/http/models"
	response "github.com/rahulvarma07/goo_backend/internal/http/utils"
)

func CreateStudent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// the w here (response writer)
		// is used to send the data as response
		// some important things to remember while sending the response
		// 1.content-type (Header)
		// 2.Status code ()
		// 3.actual reponse

		// the r (request)
		// is used to get the user response
		// in response we can have the body sent by the user

		// TODO:
		// make a student model
		// decode it using json.Decode
		// check for any error
		// according to that return the response

		var studentModel models.Student

		// adding validators

		err := json.NewDecoder(r.Body).Decode(&studentModel)

		if errors.Is(err, io.EOF) {
			response.WriteResponse(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		if err != nil {
			response.WriteResponse(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// checking for validations
		err = validator.New().Struct(studentModel)
		if err != nil {
			validationError := err.(validator.ValidationErrors)
			response.WriteResponse(w, http.StatusBadRequest, response.CheckValidation(validationError))
			return
		}

		// if success
		response.WriteResponse(w, http.StatusCreated, map[string]string{"message": "created"})
	}
}
