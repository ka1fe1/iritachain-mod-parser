package iritachain_mod_parser

import (
	. "github.com/kaifei-bianjie/common-parser"
	"github.com/kaifei-bianjie/iritachain-mod-parser/codec"
	"github.com/kaifei-bianjie/iritachain-mod-parser/modules/evm"
	"github.com/kaifei-bianjie/iritachain-mod-parser/modules/opb"
)

type MsgClient struct {
	Evm Client
	Opb Client
}

func NewMsgClient() MsgClient {
	codec.MakeEncodingConfig()
	return MsgClient{
		Evm: evm.NewClient(),
		Opb: opb.NewClient(),
	}
}
