package opb

import (
	"github.com/cosmos/cosmos-sdk/types"
	common "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritachain-mod-parser/modules"
)

type OpbClient struct {
}

func NewClient() OpbClient {
	return OpbClient{}
}

func (opb OpbClient) HandleTxMsg(v types.Msg) (common.MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgMint:
		docMsg := DocMsgMint{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgReclaim:
		docMsg := DocMsgReclaim{}
		return docMsg.HandleTxMsg(msg), true
	default:
		return common.MsgDocInfo{}, false
	}

}
