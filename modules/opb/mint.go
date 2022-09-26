package opb

import (
	"fmt"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritachain-mod-parser/modules"
)

type DocMsgMint struct {
	Amount    string `bson:"amount"`
	Recipient string `bson:"recipient"`
	Operator  string `bson:"operator"`
}

func (m *DocMsgMint) GetType() string {
	return MsgTypeMint
}

func (m *DocMsgMint) BuildMsg(v interface{}) {
	msg := v.(*MsgMint)
	m.Amount = fmt.Sprint(msg.Amount)
	m.Recipient = msg.Recipient
	m.Operator = msg.Operator
}

func (m *DocMsgMint) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgMint)
	addrs = append(addrs, msg.Operator)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
