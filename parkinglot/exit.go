package main


type Exit interface{
	CalculateCost(id int) float64
	ExitVehicle(id int)
}

type exit struct{
	id int
	ticket TicketManager
	costComputation map[VehicleType]CostComputation
	ParkingSpotManager map[VehicleType]ParkingSpotManager
}

func NewExit(id int, ticket TicketManager, costComputation map[VehicleType]CostComputation, ParkingSpotManager map[VehicleType]ParkingSpotManager) Exit{
	return &exit{
		id: id,
		ticket: ticket,
		costComputation: costComputation,
		ParkingSpotManager: ParkingSpotManager,
	}
}

func (e *exit) CalculateCost(id int) float64{
	ticket := e.ticket.GetTicket(id)
	if ticket == nil{
		return 0
	}

	// get the cost computation based on the vehicle type
	vehicle  := ticket.Vehicle()
	costComputation := e.getCostComputation(vehicle)

	// calculate the cost
	ps := ticket.ParkingSpot()

	startTime := ticket.StartTime()
	cost := costComputation.Price(startTime, ps)

	return cost
}

func (e *exit) getCostComputation(vehicle Vehicle) CostComputation{
	switch vehicle.VehicleType{
	case Motorcycle:
		return e.costComputation[Motorcycle]
	case Car:
		return e.costComputation[Car]
	}
	return NewDefaultCostComputation()
}

func (e *exit) ExitVehicle(id int){
	ticket := e.ticket.GetTicket(id)
	if ticket == nil{
		return
	}

	// get the parking spot manager based on the vehicle type
	vehicle := ticket.Vehicle()
	psm := e.getParkingSpotManager(vehicle)

	// free the parking spot
	psm.RemoveVehicle(ticket.ParkingSpot())

	// remove the ticket or mark it as paid or something
	e.ticket.RemoveTicket(id)
}

func (e *exit) getParkingSpotManager(vehicle Vehicle) ParkingSpotManager{
	switch vehicle.VehicleType{
	case Motorcycle:
		return e.ParkingSpotManager[Motorcycle]
	case Car:
		return e.ParkingSpotManager[Car]
	}
	return nil
}