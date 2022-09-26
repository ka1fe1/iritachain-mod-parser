package opb

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritachain-mod-parser/modules"
)

type DocMsgReclaim struct {
	Denom     string `bson:"amount"`
	Recipient string `bson:"recipient"`
	Operator  string `bson:"operator"`
}

func (m *DocMsgReclaim) GetType() string {
	return MsgTypeReclaim
}

func (m *DocMsgReclaim) BuildMsg(v interface{}) {
	msg := v.(*MsgReclaim)
	m.Denom = msg.Denom
	m.Recipient = msg.Recipient
	m.Operator = msg.Operator
}

func (m *DocMsgReclaim) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgReclaim)
	addrs = append(addrs, msg.Operator)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
