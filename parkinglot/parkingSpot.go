package main


type ParkingSpot interface{
	IsSpotAvailable() bool
	ParkVehicle() bool
	RemoveVehicle() bool
	Price() int
}


type TwoWheelerParkingSpot struct{
	id int
	isAvailable bool
	price int
	noOfFloor int
}

func NewTwoWheelerParkingSpot(id int, price int, noOfFloor int) *TwoWheelerParkingSpot{
	return &TwoWheelerParkingSpot{
		id: id,
		isAvailable: true,
		price: price,
		noOfFloor: noOfFloor,
	}
}

func (t *TwoWheelerParkingSpot) NoOfFloor() int{
	return t.noOfFloor
}

func (t *TwoWheelerParkingSpot) IsSpotAvailable() bool{
	return t.isAvailable
}

func (t *TwoWheelerParkingSpot) ParkVehicle() bool{
	if t.isAvailable{
		t.isAvailable = false
		return true
	}
	return false
}

func (t *TwoWheelerParkingSpot) RemoveVehicle() bool{
	t.isAvailable = true
	return true
}

func (t *TwoWheelerParkingSpot) Price() int{
	return t.price
}


type FourWheelerParkingSpot struct{
	id int
	isAvailable bool
	price int
	noOfFloor int
}

func NewFourWheelerParkingSpot(id int, price int, noOfFloor int) *TwoWheelerParkingSpot{
	return &TwoWheelerParkingSpot{
		id: id,
		isAvailable: true,
		price: price,
		noOfFloor: noOfFloor,
	}
}

func (t *FourWheelerParkingSpot) NoOfFloor() int{
	return t.noOfFloor
}

func (t *FourWheelerParkingSpot) IsSpotAvailable() bool{
	return t.isAvailable
}

func (t *FourWheelerParkingSpot) ParkVehicle() bool{
	if t.isAvailable{
		t.isAvailable = false
		return true
	}
	return false
}

func (t *FourWheelerParkingSpot) RemoveVehicle() bool{
	t.isAvailable = true
	return true
}

func (t *FourWheelerParkingSpot) Price() int{
	return t.price
}