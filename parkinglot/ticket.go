package main

import "time"

type Ticket interface{
	ParkingSpot() ParkingSpot
	Vehicle() Vehicle
	ID() int
	StartTime() time.Time

}

type ticket struct{
	id int
	startTime time.Time
	vehicle Vehicle
	parkingSpot ParkingSpot
}

func NewTicket(id int, vehicle Vehicle, parkingSpot ParkingSpot) Ticket{
	return &ticket{
		id: id,
		vehicle: vehicle,
		parkingSpot: parkingSpot,
		startTime: time.Now(),
	}
}

func (t *ticket) StartTime() time.Time{
	return t.startTime
}

func (t *ticket) Vehicle() Vehicle{
	return t.vehicle
}


func (t *ticket) ParkingSpot() ParkingSpot{
	return t.parkingSpot
}

func (t *ticket) ID() int{
	return t.id
}