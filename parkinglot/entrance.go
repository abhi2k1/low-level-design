package main

type Entrance interface{
	GetTicket(id int) Ticket
	ParkVehicle(vehicle Vehicle) Ticket
}


type entrance struct{
	id int
	parkingSpotManagers map[VehicleType]ParkingSpotManager
	ticket TicketManager
}

func NewEntrance(id int, parkingSpotManagers map[VehicleType]ParkingSpotManager, ticket TicketManager) *entrance{
	return &entrance{
		id: id,
		parkingSpotManagers: parkingSpotManagers,
		ticket: ticket,
	}
}

func (e *entrance) GetTicket(id int) Ticket{
	return e.ticket.GetTicket(id)
}

func (e *entrance) ParkVehicle(vehicle Vehicle) Ticket{
	// get the parking spot manager based on the vehicle type
	parkingSpotManager := e.getParkingSpotManager(vehicle)

	// find the parking spot
	parkingSpot := parkingSpotManager.FindParkingSpot()
	if parkingSpot == nil{
		return nil
	}

	// park the vehicle
	if !parkingSpot.ParkVehicle(){
		return nil
	}

	// create a new ticket
	ticket := e.ticket.IssueTicket(vehicle, parkingSpot)

	return ticket
}


func (e *entrance) getParkingSpotManager(vehicle Vehicle) ParkingSpotManager{
	switch vehicle.VehicleType{
	case Motorcycle:
		return e.parkingSpotManagers[Motorcycle]
	case Car:
		return e.parkingSpotManagers[Car]
	}
	return nil
}