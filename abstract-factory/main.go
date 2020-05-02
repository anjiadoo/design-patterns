package main

import (
	"fmt"

	concretefactory "design-patterns/abstract-factory/factoryimpl"
)

func main() {
	fmt.Println("**************************************************")
	fmt.Println("创建X1系列产品")
	X1 := concretefactory.GetFactoryX1
	x1pa := X1().CreateProductA()
	x1pb := X1().CreateProductB()
	fmt.Printf("X1系列产品%v的id是%v，价格是$%v.\n", x1pa.GetProductName(), x1pa.GetProductID(), x1pa.GetProductPrice())
	fmt.Printf("X1系列产品%v的id是%v，价格是$%v.\n", x1pb.GetProductName(), x1pb.GetProductID(), x1pb.GetProductPrice())

	fmt.Println("**************************************************")
	fmt.Println("创建X2系列产品")
	X2 := concretefactory.GetFactoryX2
	x2pa := X2().CreateProductA()
	x2pb := X2().CreateProductB()
	fmt.Printf("X2系列产品%v的id是%v，价格是$%v.\n", x2pa.GetProductName(), x2pa.GetProductID(), x2pa.GetProductPrice())
	fmt.Printf("X2系列产品%v的id是%v，价格是$%v.\n", x2pb.GetProductName(), x2pb.GetProductID(), x2pb.GetProductPrice())
}
