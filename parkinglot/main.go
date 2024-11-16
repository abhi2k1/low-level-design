package main

import "fmt"


func main() {
	
	// parking spot - 6 two wheeler and 4 four wheeler
	twoWheelerSpots := make([]ParkingSpot, 0)
	fourWheelerSpots := make([]ParkingSpot, 0)

	for i := 0; i < 6; i++{
		twoWheelerSpots = append(twoWheelerSpots, NewTwoWheelerParkingSpot(i, 10, 0))
	}

	for i:= 0; i<4;i++{
		fourWheelerSpots = append(fourWheelerSpots, NewFourWheelerParkingSpot(i, 20, 0))
	}

	// parking strategy
	parkingStategy := NewNearestParkingStrategy()


	// parking spot manager
	twoWheelerParkingSpotManager := NewTwoWheelerParkingSpotManager(twoWheelerSpots, parkingStategy)

	fourWheelerParkingSpotManager := NewFourWheelerParkingSpotManager(fourWheelerSpots, parkingStategy)

	parkingSpotManager := map[VehicleType]ParkingSpotManager{
		Motorcycle: twoWheelerParkingSpotManager,
		Car: fourWheelerParkingSpotManager,
	}

	// ticket manager

	ticketManager := NewTicketManager()

	// cost computation

	twoWheelerCostComputation := NewTwoWheelerCostComputation()
	fourWheelerCostComputation := NewFourWheelerCostComputation()

	entrance := NewEntrance(1,parkingSpotManager, ticketManager)

	exist := NewExit(1, ticketManager, map[VehicleType]CostComputation{
		Motorcycle: twoWheelerCostComputation,
		Car: fourWheelerCostComputation,
	}, parkingSpotManager)


	// entrance vehicle 
	t1 := entrance.ParkVehicle(NewVehicle(1, Motorcycle))
	t2 := entrance.ParkVehicle(NewVehicle(2, Motorcycle))

	t3 := entrance.ParkVehicle(NewVehicle(3, Car))
	t4 := entrance.ParkVehicle(NewVehicle(4, Car))


	// calculate cost
	c1 := exist.CalculateCost(t1.ID())
	fmt.Println("Cost of vehicle 1: ", c1)

	c2 :=  exist.CalculateCost(t2.ID())
	fmt.Println("Cost of vehicle 2: ", c2)

	c3 := exist.CalculateCost(t3.ID())
	fmt.Println("Cost of vehicle 3: ", c3)

	c4 := exist.CalculateCost(t4.ID())
	fmt.Println("Cost of vehicle 4: ", c4)
	
	
	// remove vehicle
	exist.ExitVehicle(t1.ID())
	exist.ExitVehicle(t2.ID())
	exist.ExitVehicle(t3.ID())
	exist.ExitVehicle(t4.ID())	
}