package gov

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/iritachain-mod-parser/modules"
)

type GovClient struct {
}

func NewClient() GovClient {
	return GovClient{}
}

func (gov GovClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg := v.(type) {
	case *MsgSubmitProposal:
		docMsg := DocTxMsgSubmitProposal{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgVote:
		docMsg := DocTxMsgVote{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgDeposit:
		docMsg := DocTxMsgDeposit{}
		return docMsg.HandleTxMsg(msg), ok
	case *MsgVoteWeighted:
		docMsg := DocTxMsgVoteWeighted{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
