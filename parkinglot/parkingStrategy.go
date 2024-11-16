package main

type ParkingStrategy interface{
	FindParkingSpot(ps []ParkingSpot) ParkingSpot
}

type NearestParkingStrategy struct{}

func NewNearestParkingStrategy() ParkingStrategy{
	return &NearestParkingStrategy{}
}

func (n *NearestParkingStrategy) FindParkingSpot(ps []ParkingSpot) ParkingSpot{
	for _, spot := range ps{
		if spot.IsSpotAvailable(){
			return spot
		}
	}
	return nil
}

type NearestToEntranceParkingStrategy struct{}

func NewNearestToEntranceParkingStrategy() ParkingStrategy{
	return &NearestToEntranceParkingStrategy{}
}

func (n *NearestToEntranceParkingStrategy) FindParkingSpot(ps []ParkingSpot) ParkingSpot{
	// business logic to find the nearest parking spot to the entrance
	// todo: implement the logic
	panic("implement me")
}

type NearestToExitParkingStrategy struct{}

func NewNearestToExitParkingStrategy() ParkingStrategy{
	return &NearestToExitParkingStrategy{}
}

func (n *NearestToExitParkingStrategy) FindParkingSpot(ps []ParkingSpot) ParkingSpot{
	// business logic to find the nearest parking spot to the exit
	panic("implement me")
}