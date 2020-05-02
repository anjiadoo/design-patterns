package fefinedmessage

import abstractmessage "design-patterns/bridge/imessage"

//OrdinaryMessage 扩充高层次抽象类的接口
type OrdinaryMessage struct {
	*abstractmessage.AbstractMessage
}

//SendMessage xx
func (o *OrdinaryMessage) SendMessage(msg string, user string) {
	o.AbstractMessage.SendMessage(msg, user)
}

//ReceiveMessage xx
func (o *OrdinaryMessage) ReceiveMessage(msg string, user string) {
	o.AbstractMessage.ReceiveMessage(msg, user)
}
