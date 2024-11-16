package main

type VehicleType int

const (
	Motorcycle VehicleType = iota + 1
	Car 
)


type Vehicle struct{
	VehicleNum int
	VehicleType VehicleType
}

func NewVehicle(vehicleNum int, vehicleType VehicleType) Vehicle{
	return Vehicle{
		VehicleNum: vehicleNum,
		VehicleType: vehicleType,
	}
}