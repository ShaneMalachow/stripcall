package stripcall

import (
	"fmt"
	"net/http"
)

func HandleError(w *http.ResponseWriter, message string, code int) {
	fmt.Println(message)
	http.Error(*w, message, code)
	return
}
