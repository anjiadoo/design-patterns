package director

import abstractbuilder "design-patterns/builder/ibuilder"

//Director 导向器
type Director struct {
	Builder abstractbuilder.Builder
}
