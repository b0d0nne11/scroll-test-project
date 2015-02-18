package handlers

import (
	"fmt"
	"net/http"

	"github.com/mailgun/scroll"
)

func HelloWorld(w http.ResponseWriter, r *http.Request, params map[string]string) (interface{}, error) {
	return scroll.Response{
		"message": fmt.Sprintf("Hello world!"),
	}, nil
}

func ReplyNotImplemented(w http.ResponseWriter, r *http.Request) {
	scroll.Reply(w, scroll.Response{"message": "Not Implemented"}, http.StatusNotImplemented)
}
