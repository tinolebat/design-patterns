package main

// Protection Proxy
import (
	"fmt"
)

type Driven interface {
	Drive()
}

type Car struct {
}

// Car struct{} is implementing the Driven interface with the Drive function
func (c *Car) Drive() {
	fmt.Println("Car is being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car    Car
	driver *Driver
}

func newCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{
		car:    Car{},
		driver: driver,
	}
}

func (cp *CarProxy) Drive() {
	if cp.driver.Age >= 16 {
		cp.car.Drive()
	} else {
		fmt.Println("Driver is too young")
	}
}

func main() {

	car := newCarProxy(&Driver{32})
	car.Drive()

}
