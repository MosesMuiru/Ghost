package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type BookRideRequest struct {
	UserId   string `json:"user_id"`
	RidersId string `json:"riders_Id"`
	Book     bool   `json:"book"`
}

type Start struct {
	StartYourShift bool `json:"start"`
}

type BookRide struct{}

var v any

// this is to access the data in V

func CUSTOMERDATA() *any {
	return &v

}

func (b *BookRide) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, _ := RequestDecoder[BookRideRequest](r)
	fmt.Println("this is the response", data)
	response := "response booked"
	_ = RequestEncoder[string](w, r, http.StatusOK, response)

}

func (s *Start) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	data, _ := RequestDecoder[BookRideRequest](r)
	fmt.Println("This is the data that has come back", data)

	// if its true, send your current location
	// it will sse to pool you current location
	// then after streaming your current location from the phone
	// a connection socket (hande shake is created j)
	_ = RequestEncoder[string](w, r, http.StatusOK, "Your streaming has started")

}

func RequestDecoder[T any](r *http.Request) (any, error) {

	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("Decode json: %w", err)
	}

	return v, nil
}

func RequestEncoder[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

type WebSocketHandler struct{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func WebSocketApi(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading", err)
		return
	}

	defer conn.Close()

	// listen for incoming locations sent
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Buda hii websocket imestop:---> %v", err)
			}

		}
		fmt.Println("This is the message----> ", message)

		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}
