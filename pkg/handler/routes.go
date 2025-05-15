package handler

import (
	"net/http"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/book", &BookRide{})
	mux.HandleFunc("/websocket", WebSocketApi)

	return mux

}
