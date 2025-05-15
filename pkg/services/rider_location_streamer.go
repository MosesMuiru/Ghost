package services

import (
	"github.com/ghost/pkg/handler"
)

// this will be streaming the location of the rider
// if the set location is trur then the rider is streamed
// there are two

type RidersData struct {
	Log     string
	Lat     string
	RiderId string
}

// just a custuctor to access the detials

func NewDetails() *any {
	h := handler.CUSTOMERDATA()
	return h
}

func StreamLocation(booked bool, rider_id string) {
	//	v := handler.CUSTOMERDATA()
	///	ridersChannel := make(chan RidersData)

	// create a channel tunnel where the infomation of the user can flow easily
	// we need to make sure that the rider and the customer are linked well

	// using the riders detials, make sure here belongs to a certain channel

	// a user can use the detials (eg the rider_id) to book
	// rider of this ride id ------will be linked with use of this id----

	// through a channel ---location will be streamed to the customer-----

	// check if the rider is already in another ride or free
	//isFree := IfRiderIsFree(rider_id)

	//RidersData := make(chan RidersData)

	// from here i can pass to

}

func IfRiderIsFree(rirder_id string) bool {
	return true
}

// for location

// i should have a way to which the customer only streams his location
// just the location of the rider is important
