package product

import "fmt"

//Car 成型的产品
type Car struct {
	Name        string
	Color       string
	Engine      string
	Wheel       string
	CircleWheel string
}

//PrintDescribe 驾驶
func (c *Car) PrintDescribe() {
	fmt.Printf("the Car name is %v.\n", c.Name)
	fmt.Printf("the Car Color is %v.\n", c.Color)
	fmt.Printf("the Car Wheel is %v.\n", c.Wheel)
	fmt.Printf("the Car Engine is %v.\n", c.Engine)
	fmt.Printf("the Car CircleWheel is %v.\n", c.CircleWheel)
	fmt.Println("***************************************")
}
