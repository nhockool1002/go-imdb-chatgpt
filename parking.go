package main

import (
	"fmt"
)

type Car struct {
	PlateNumber string
	Brand       string
	Color       string
}

type ParkingLot struct {
	Capacity int
	Cars     []Car
}

func (p *ParkingLot) parkCar(car Car) {
	if len(p.Cars) < p.Capacity {
		p.Cars = append(p.Cars, car)
		fmt.Printf("Xe có biển số %s đã được đậu trong bãi đậu xe.\n", car.PlateNumber)
	} else {
		fmt.Println("Bãi đậu xe đã đầy.")
	}
}

func (p *ParkingLot) removeCar(plateNumber string) {
	for i, car := range p.Cars {
		if car.PlateNumber == plateNumber {
			p.Cars = append(p.Cars[:i], p.Cars[i+1:]...)
			fmt.Printf("Xe có biển số %s đã rời khỏi bãi đậu xe.\n", plateNumber)
			return
		}
	}
	fmt.Printf("Không tìm thấy xe có biển số %s trong bãi đậu xe.\n", plateNumber)
}

func (p ParkingLot) showCars() {
	if len(p.Cars) == 0 {
		fmt.Println("Bãi đậu xe hiện đang trống.")
		return
	}
	fmt.Println("Danh sách xe trong bãi đậu xe:")
	for _, car := range p.Cars {
		fmt.Printf("Biển số: %s, Hãng: %s, Màu sắc: %s\n", car.PlateNumber, car.Brand, car.Color)
	}
}

func main() {
	var myParkingLot ParkingLot
	myParkingLot.Capacity = 3

	myParkingLot.parkCar(Car{PlateNumber: "ABC123", Brand: "Toyota", Color: "Black"})
	myParkingLot.parkCar(Car{PlateNumber: "XYZ456", Brand: "Honda", Color: "Red"})
	myParkingLot.showCars()

	myParkingLot.removeCar("ABC123")
	myParkingLot.showCars()
}
