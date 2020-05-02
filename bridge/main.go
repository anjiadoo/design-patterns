package main

import (
	"fmt"

	bridge "design-patterns/bridge/imessage"
	sendmode "design-patterns/bridge/implementor"
	fefinedmessage "design-patterns/bridge/messageimpl"
)

func main() {
	SMSImpl := &sendmode.MessageSMS{}
	m := &fefinedmessage.OrdinaryMessage{&bridge.AbstractMessage{SMSImpl}}
	m.ReceiveMessage("晚上一起看电影吗?", "wanlili")

	um := &fefinedmessage.UrgencyMessage{&bridge.AbstractMessage{SMSImpl}}
	um.ReceiveMessage("晚上一起看电影吗?", "wanlili")

	sm := &fefinedmessage.SpecialUrgencyMessage{&bridge.AbstractMessage{SMSImpl}}
	sm.ReceiveMessage("晚上一起看电影吗?", "wanlili")

	fmt.Println("*******************************************************")

	MobileImpl := &sendmode.MessageMobile{}
	m = &fefinedmessage.OrdinaryMessage{&bridge.AbstractMessage{MobileImpl}}
	m.ReceiveMessage("晚上一起看电影吗?", "wanlili")

	um = &fefinedmessage.UrgencyMessage{&bridge.AbstractMessage{MobileImpl}}
	um.ReceiveMessage("晚上一起看电影吗?", "wanlili")

	sm = &fefinedmessage.SpecialUrgencyMessage{&bridge.AbstractMessage{MobileImpl}}
	sm.ReceiveMessage("晚上一起看电影吗?", "wanlili")

	fmt.Println("*******************************************************")

	EmailImpl := &sendmode.MessageEmail{}
	m = &fefinedmessage.OrdinaryMessage{&bridge.AbstractMessage{EmailImpl}}
	m.ReceiveMessage("晚上一起看电影吗?", "wanlili")

	um = &fefinedmessage.UrgencyMessage{&bridge.AbstractMessage{EmailImpl}}
	um.ReceiveMessage("晚上一起看电影吗?", "wanlili")

	sm = &fefinedmessage.SpecialUrgencyMessage{&bridge.AbstractMessage{EmailImpl}}
	sm.ReceiveMessage("晚上一起看电影吗?", "wanlili")
}
