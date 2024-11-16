package main

type ParkingSpotManager interface{
	FindParkingSpot() ParkingSpot
	AddParkingSpot(ps ParkingSpot) 
	RemoveParkingSpot(ps ParkingSpot)
	ParkVehicle(p ParkingSpot) bool
	RemoveVehicle(p ParkingSpot) bool
}

type TwoWheelerParkingSpotManager struct{
	parkingSpot []ParkingSpot
	parkingStrategy ParkingStrategy
}

func NewTwoWheelerParkingSpotManager(spots []ParkingSpot, parkingStrategy ParkingStrategy) *TwoWheelerParkingSpotManager{
	return &TwoWheelerParkingSpotManager{
		parkingSpot: spots,
		parkingStrategy: parkingStrategy,
	}
}

func (t *TwoWheelerParkingSpotManager) FindParkingSpot() ParkingSpot{
	// todo: create a stategy to find the parking spot
	return t.parkingStrategy.FindParkingSpot(t.parkingSpot)
}

func (t *TwoWheelerParkingSpotManager) AddParkingSpot(ps ParkingSpot){
	t.parkingSpot = append(t.parkingSpot, ps)
}

func (t *TwoWheelerParkingSpotManager) RemoveParkingSpot(ps ParkingSpot){
	for i, spot := range t.parkingSpot{
		if spot == ps{
			t.parkingSpot = append(t.parkingSpot[:i], t.parkingSpot[i+1:]...)
		}
	}
}

func (t *TwoWheelerParkingSpotManager) ParkVehicle(p ParkingSpot) bool{
	return p.ParkVehicle()
}

func (t *TwoWheelerParkingSpotManager) RemoveVehicle(p ParkingSpot) bool{
	return p.RemoveVehicle()
}

type FourWheelerParkingSpotManager struct{
	parkingSpot []ParkingSpot
	parkingStrategy ParkingStrategy
}

func NewFourWheelerParkingSpotManager(spots []ParkingSpot, parkingStrategy ParkingStrategy) *FourWheelerParkingSpotManager{
	return &FourWheelerParkingSpotManager{
		parkingSpot: spots,
		parkingStrategy: parkingStrategy,
	}
}

func (f *FourWheelerParkingSpotManager) FindParkingSpot() ParkingSpot{
	return f.parkingStrategy.FindParkingSpot(f.parkingSpot)
}

func (f *FourWheelerParkingSpotManager) AddParkingSpot(ps ParkingSpot){
	f.parkingSpot = append(f.parkingSpot, ps)
}


func (f *FourWheelerParkingSpotManager) RemoveParkingSpot(ps ParkingSpot){
	for i, spot := range f.parkingSpot{
		if spot == ps{
			f.parkingSpot = append(f.parkingSpot[:i], f.parkingSpot[i+1:]...)
		}
	}
}

func (f *FourWheelerParkingSpotManager) ParkVehicle(p ParkingSpot) bool{
	return p.ParkVehicle()
}

func (f *FourWheelerParkingSpotManager) RemoveVehicle(p ParkingSpot) bool{
	return p.RemoveVehicle()
}



