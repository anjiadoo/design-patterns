package abstractproduct

// ProductA 产品A接口
type ProductA interface {
	GetProductName() string
	GetProductID() string
	GetProductPrice() int32
	PrintProductPrice()
}
