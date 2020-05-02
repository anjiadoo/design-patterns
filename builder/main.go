package main

import (
	"design-patterns/builder/builderimpl"
	"design-patterns/builder/director"
)

func main() {
	dr := director.Director{&concretebuilder.ConcreteBuilder{}}
	VOVCar := dr.Builder.BuilderName("沃尔沃").BuilderColor("红色").BuilderEngine("V-X2b.101").BuilderWheel("VoV-30b-25").BuilderCircleWheel("V-1X").GetCar()
	VOVCar.PrintDescribe()

	BMWCar := dr.Builder.BuilderName("宝马").BuilderColor("太空灰").BuilderEngine("B-X5g.325").BuilderWheel("BMW-56r-12").BuilderCircleWheel("B-4A").GetCar()
	BMWCar.PrintDescribe()
}
