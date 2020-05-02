package fefinedmessage

import bridge "design-patterns/bridge/imessage"

//SpecialUrgencyMessage 扩充高层次抽象类的接口
type SpecialUrgencyMessage struct {
	*bridge.AbstractMessage
}

//SendMessage zz
func (s *SpecialUrgencyMessage) SendMessage(msg string, user string) {
	msg = "【加急】" + msg
	s.AbstractMessage.SendMessage(msg, user)
}

//ReceiveMessage xx
func (s *SpecialUrgencyMessage) ReceiveMessage(msg string, user string) {
	msg = "【加急】" + msg
	s.AbstractMessage.ReceiveMessage(msg, user)
}
