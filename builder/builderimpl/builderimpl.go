package concretebuilder

import (
	abstractbuilder "design-patterns/builder/ibuilder"
	"design-patterns/builder/product"
)

//ConcreteBuilder 具体建造者
type ConcreteBuilder struct {
	WowCar *product.Car
}

//BuilderName 车名
func (c *ConcreteBuilder) BuilderName(name string) abstractbuilder.Builder {
	if c.WowCar == nil {
		c.WowCar = new(product.Car)
	}
	c.WowCar.Name = name
	return c
}

//BuilderColor 车身颜色
func (c *ConcreteBuilder) BuilderColor(color string) abstractbuilder.Builder {
	if c.WowCar == nil {
		c.WowCar = new(product.Car)
	}
	c.WowCar.Color = color
	return c
}

//BuilderEngine 装配引擎
func (c *ConcreteBuilder) BuilderEngine(Engine string) abstractbuilder.Builder {
	if c.WowCar == nil {
		c.WowCar = new(product.Car)
	}
	c.WowCar.Engine = Engine
	return c
}

//BuilderWheel 装配车轮
func (c *ConcreteBuilder) BuilderWheel(Wheel string) abstractbuilder.Builder {
	if c.WowCar == nil {
		c.WowCar = new(product.Car)
	}
	c.WowCar.Wheel = Wheel
	return c
}

//BuilderCircleWheel 装配方向盘
func (c *ConcreteBuilder) BuilderCircleWheel(CircleWheel string) abstractbuilder.Builder {
	if c.WowCar == nil {
		c.WowCar = new(product.Car)
	}
	c.WowCar.CircleWheel = CircleWheel
	return c
}

//GetCar 获取一辆小车
func (c *ConcreteBuilder) GetCar() *product.Car {
	return c.WowCar
}
