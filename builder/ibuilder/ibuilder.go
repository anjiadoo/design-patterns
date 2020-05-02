package abstractbuilder

import "design-patterns/builder/product"

// Builder 抽象建造者
type Builder interface {
	BuilderName(name string) Builder
	BuilderColor(color string) Builder
	BuilderEngine(engine string) Builder
	BuilderWheel(wheel string) Builder
	BuilderCircleWheel(circleWheel string) Builder
	GetCar() *product.Car
}
