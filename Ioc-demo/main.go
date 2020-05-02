package main

import (
	"fmt"
	"reflect"
)

var inj *Injector

//Injector 注入器
type Injector struct {
	mappers map[reflect.Type]reflect.Value // 根据类型映射实际的值
}

//PushParameter xx
func (inj *Injector) PushParameter(value interface{}) *Injector {
	inj.mappers[reflect.TypeOf(value)] = reflect.ValueOf(value)
	return inj
}

//PopParameter xx
func (inj *Injector) PopParameter(t reflect.Type) reflect.Value {
	return inj.mappers[t]
}

//Call xx
func (inj *Injector) Call(cb interface{}) interface{} {
	t := reflect.TypeOf(cb)
	if t.Kind() != reflect.Func {
		panic("Should invoke a function!")
	}

	inValueSlice := make([]reflect.Value, t.NumIn())
	for k := 0; k < t.NumIn(); k++ {
		inValueSlice[k] = inj.PopParameter(t.In(k))
	}
	return reflect.ValueOf(cb).Call(inValueSlice)
}

//CallBack 注入的依赖
func CallBack(a int, b string) string {
	fmt.Printf("依赖注入:CallBack args1:%v, args2:%v\n", a, b)
	return `依赖注入:Done.`
}

func main() {
	// 创建注入器
	inj = &Injector{make(map[reflect.Type]reflect.Value)}

	ret := inj.PushParameter(8090).PushParameter("anjiadoo").Call(CallBack)
	fmt.Println(ret)
	fmt.Println("**************************************************")

	ret = inj.PushParameter(1019).PushParameter("wanlili").Call(CallBack)
	fmt.Println(ret)
	fmt.Println("**************************************************")
}
