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

	printOccupancyLevel(computeOccupancyLevel(roomsOccupied, numberOfRooms))

	if roomsAvailable > 0 {
		for i := 0; i < roomsAvailable; i++ {
			printCurrentRoomStatus(startingRoomNumber, i)
		}
	} else {
		fmt.Println("No rooms are available tonight")
	}

}

func printCurrentRoomStatus(startingRoomNumber int, currentRoomNumber int) {
	roomNumber := startingRoomNumber + currentRoomNumber
	people := rand.Intn(10-1) + 1
	nights := rand.Intn(6-1) + 1
	nightText := " night"
	peopleText := " person "
	if nights > 1 {
		nightText = " nights"
	}
	if people > 1 {
		peopleText = " people "
	}
	fmt.Println(roomNumber, " : ", people, peopleText, nights, nightText)
}

func computeOccupancyLevel(roomsOccupied int, numberOfRooms int) float64 {
	return (float64(roomsOccupied) / float64(numberOfRooms)) * 100
}

func printOccupancyLevel(occupancy float64) {
	switch {
	case occupancy >= 60:
		fmt.Println("Occupancy Rate : ", "High")
	case occupancy >= 30:
		fmt.Println("Occupancy Rate : ", "Medium")
	default:
		fmt.Println("Occupancy Rate : ", "Low")
	}
	fmt.Printf("Occupancy Percentage : %0.2f %% \n", occupancy)
}
