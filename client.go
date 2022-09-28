package iritachain_mod_parser

import (
	. "github.com/kaifei-bianjie/common-parser"
	"github.com/kaifei-bianjie/iritachain-mod-parser/codec"
	"github.com/kaifei-bianjie/iritachain-mod-parser/modules/evm"
)

type MsgClient struct {
	Evm Client
}

func NewMsgClient() MsgClient {
	codec.MakeEncodingConfig()
	return MsgClient{
		Evm: evm.NewClient(),
	}
}
