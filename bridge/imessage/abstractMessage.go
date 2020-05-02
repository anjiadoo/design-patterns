package abstractmessage

import implementor "design-patterns/bridge/iImplementor"

//AbstractMessage 定义高层次抽象类的接口，维护一个指向Implementor对象的指针
type AbstractMessage struct {
	Impl implementor.MessageImplementor
}

//SendMessage 方法实现
func (a *AbstractMessage) SendMessage(msg string, user string) {
	a.Impl.Send(msg, user)
}

//ReceiveMessage 方法实现
func (a *AbstractMessage) ReceiveMessage(msg string, user string) {
	a.Impl.Receive(msg, user)
}
