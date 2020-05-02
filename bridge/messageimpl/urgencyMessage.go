package fefinedmessage

import bridge "design-patterns/bridge/imessage"

//UrgencyMessage 扩充高层次抽象类的接口
type UrgencyMessage struct {
	*bridge.AbstractMessage
}

//SendMessage zz
func (u *UrgencyMessage) SendMessage(msg string, user string) {
	msg = "【紧急】" + msg
	u.AbstractMessage.SendMessage(msg, user)
}

//ReceiveMessage xx
func (u *UrgencyMessage) ReceiveMessage(msg string, user string) {
	msg = "【紧急】" + msg
	u.AbstractMessage.ReceiveMessage(msg, user)
}
