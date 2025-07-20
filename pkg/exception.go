package pkg

import (
	"fmt"
	"net/http"
	"strings"
)

func HandleError(w http.ResponseWriter, err interface{}) {

	str := fmt.Sprintln(err)
	arr := strings.Split(str, "~")

	key := arr[0]
	msg := arr[len(arr)-1]

	switch key {
	case "400":
		WriteError(w, http.StatusBadRequest, msg)
	default:
		WriteError(w, http.StatusInternalServerError, msg)
	}

}

func PanicException(status int, message string) {
	err := fmt.Errorf("%v~%v", status, message)
	panic(err)
}

func BadRequestError(messages ...string) {
	var msg string = "Bad Request"

	if len(messages) > 0 {
		msg = strings.Join(messages, "`")
	}

	PanicException(http.StatusBadRequest, msg)
}

func InternalServerError(messages ...string) {
	var msg string = "Internal Server Error"

	if len(messages) > 0 {
		msg = strings.Join(messages, ";")
	}

	PanicException(http.StatusInternalServerError, msg)

}
