package abstractproduct

// ProductB 产品B接口
type ProductB interface {
	GetProductName() string
	GetProductID() string
	GetProductPrice() int32
	PrintProductPrice()
	PrintProductPrices()
}
