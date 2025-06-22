package car1

import "fmt"

type Car struct {
	Name         string
	CurrOilLevel float32
	MaxCapacity  float32
}

type Station struct {
	Name string
}

func newCar(name string, currOilLevel float32, maxCapacity float32) *Car {
	return &Car{
		Name:         name,
		CurrOilLevel: 0,
		MaxCapacity:  maxCapacity,
	}
}
func newStation(name string) *Station {
	return &Station{
		Name: name,
	}
}
func (c *Car) AddFuel(amount float32) float32 {
	availableSpace := c.MaxCapacity - c.CurrOilLevel
	if availableSpace <= 0 {
		return 0
	}
	if amount > availableSpace {
		amount = availableSpace
	}
	c.CurrOilLevel += amount
	return amount
}
func (s Station) Refuel(car *Car, amount float32) {

	if car.CurrOilLevel >= car.MaxCapacity {
		fmt.Printf("Бак полный ")
		return
	}
	added := car.AddFuel(amount)
	fmt.Printf("Добавлено %.2f литров\n", added)

}

func main() {
	car := newCar("Lada", 0, 60)
	station := newStation("Moscow")
	station.Refuel(car, 10)
}
