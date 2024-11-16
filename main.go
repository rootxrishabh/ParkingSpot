package main

import (
	"fmt"
	"time"
)

/* Designing a parking system, here are the requirements -
1. Multiple Parking levels.
2. Different types of parking spots (small, medium large).
3. Vehicles like motorcycles, cars, buses.
4. Real-time tracking of available spots.
5. Ability to calculate parking fees.
*/

type ParkingSpot struct {
	ParkingLevel           ParkelLevel
	AvailableParkingLevels AvailableParkingLevels
	ParkingSpotType        ParkingSpotType
	VehiclesParked         []Vehicle
}

type Vehicle struct {
	RegistrationNumber string
	TimeOfArrial       time.Time
	VehicleType        string
}

type ParkelLevel int

type AvailableParkingLevels int

type ParkingSpotType string

// Charges are per-hour based
const (
	motorcycleParkingCost = 50
	carParkingCost        = 100
	busParkingCost        = 150

	car        = "car"
	motorcycle = "motorcycle"
	bus        = "bus"
)

var (
	ParkingSpots []ParkingSpot
)

func calculateCheckoutCost(vehicle *Vehicle) (int) {
	timeParked :=  time.Now().Nanosecond() - vehicle.TimeOfArrial.Nanosecond()

	if vehicle.VehicleType == car {
		return timeParked * carParkingCost
	} else if vehicle.VehicleType == bus {
		return timeParked * busParkingCost
	} else if vehicle.VehicleType == motorcycle {
		return timeParked * motorcycleParkingCost
	}

	return -1
}

func ParkVehicle(vehicle *Vehicle) int {
	spot := checkAvailability(vehicle)
	if spot == -1 {
		return spot
	}

	return spot
}

func checkAvailability(vehicle *Vehicle) int {
	for spot, parkingSpot := range ParkingSpots {
		if vehicle.VehicleType == string(parkingSpot.ParkingSpotType) && parkingSpot.AvailableParkingLevels > 0 {
			vehicle.TimeOfArrial = time.Now()
			parkingSpot.VehiclesParked = append(parkingSpot.VehiclesParked, *vehicle)
			parkingSpot.AvailableParkingLevels = parkingSpot.AvailableParkingLevels - 1
			ParkingSpots[spot] = parkingSpot
			return spot
		}
	}
	return -1
}

func checkout(vehicle *Vehicle) (int) {
	for spot, parkingSpot := range ParkingSpots{
		for i, v := range parkingSpot.VehiclesParked {
			if v.RegistrationNumber == vehicle.RegistrationNumber {
				parkingSpot.VehiclesParked = append(parkingSpot.VehiclesParked[:i], parkingSpot.VehiclesParked[i+1:]...)
				parkingSpot.AvailableParkingLevels++
				ParkingSpots[spot] = parkingSpot
				return calculateCheckoutCost(vehicle)
			}
		}
	}
	return -1
}

// init will populate the parking spots
func init() {
	s1 := ParkingSpot{
		ParkingLevel:           3,
		AvailableParkingLevels: 3,
		ParkingSpotType:        motorcycle,
	}

	s2 := ParkingSpot{
		ParkingLevel:           3,
		AvailableParkingLevels: 3,
		ParkingSpotType:        car,
	}

	s3 := ParkingSpot{
		ParkingLevel:           3,
		AvailableParkingLevels: 3,
		ParkingSpotType:        bus,
	}

	ParkingSpots = append(ParkingSpots, s1, s2, s3)
}

func main() {
	var parked int

	v1 := Vehicle{
		RegistrationNumber: "RJ36CA1924",
		VehicleType:        "car",
	}

	v2 := Vehicle{
		RegistrationNumber: "RJ36CA1925",
		VehicleType:        "motorcycle",
	}

	v3 := Vehicle{
		RegistrationNumber: "RJ36CA1926",
		VehicleType:        "bus",
	}

	// Test cases

	parked = ParkVehicle(&v1)
	if parked == -1 {
		fmt.Println("No spots avaiable")
	} else {
		fmt.Println(parked)
	}

	parked = ParkVehicle(&v2)
	if parked == -1 {
		fmt.Println("No spots avaiable")
	} else {
		fmt.Println(parked)
	}

	parked = ParkVehicle(&v3)
	if parked == -1 {
		fmt.Println("No spots avaiable")
	} else {
		fmt.Println(parked)
	}
	parked = ParkVehicle(&v3)
	if parked == -1 {
		fmt.Println("No spots avaiable")
	} else {
		fmt.Println(parked)
	}
	parked = ParkVehicle(&v3)
	if parked == -1 {
		fmt.Println("No spots avaiable")
	} else {
		fmt.Println(parked)
	}
	parked = ParkVehicle(&v3)
	if parked == -1 {
		fmt.Println("No spots avaiable")
	} else {
		fmt.Println(parked)
	}
	parked = ParkVehicle(&v3)
	if parked == -1 {
		fmt.Println("No spots avaiable")
	} else {
		fmt.Println(parked)
	}

	fmt.Println(checkout(&v3))

}