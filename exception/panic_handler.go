package exception

import (
	"encoding/json"
	"github.com/Ikhlashmulya/golang-restful-api/model"
	"net/http"
)

func PanicHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if exception, ok := err.(*ErrorNotFound); ok {
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(model.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT_FOUND",
			Error:  exception.Error(),
		})

		return
	}

	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(writer).Encode(model.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err,
	})

}
