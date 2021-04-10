package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Initializing Constants
	const hotelName string = "Gopher Paris Inn"
	const numberOfRooms int = 134
	const startingRoomNumber int = 110
	fmt.Println(hotelName)
	// Seeding Time
	rand.Seed(time.Now().UnixNano())
	roomsOccupied := rand.Intn(134)
	roomsAvailable := numberOfRooms - roomsOccupied
	var occupancy float64 = (float64(roomsOccupied) / float64(numberOfRooms)) * 100
	if occupancy >= 60 {
		fmt.Println("Occupancy Rate : ", "High")
	} else if occupancy >= 30 {
		fmt.Println("Occupancy Rate : ", "Medium")
	} else {
		fmt.Println("Occupancy Rate : ", "Low")
	}
	fmt.Println("Occupancy Percentage : ", occupancy)

	if roomsAvailable > 0 {
		for i := 0; i < roomsAvailable; i++ {
			roomNumber := startingRoomNumber + i
			people := rand.Intn(10-1) + 1
			nights := rand.Intn(6-1) + 1
			fmt.Println(roomNumber, " : ", people, " people ", nights, " nights")
		}
	} else {
		fmt.Println("No rooms are available tonight")
	}

}
