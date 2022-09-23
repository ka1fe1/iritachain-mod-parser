package iritachain_mod_parser

import (
	. "github.com/kaifei-bianjie/common-parser"
	"github.com/kaifei-bianjie/iritachain-mod-parser/codec"
	"github.com/kaifei-bianjie/iritachain-mod-parser/modules/evm"
	"github.com/kaifei-bianjie/iritachain-mod-parser/modules/gov"
	"github.com/kaifei-bianjie/iritachain-mod-parser/modules/opb"
	"github.com/kaifei-bianjie/iritachain-mod-parser/modules/tibc"
)

type MsgClient struct {
	Tibc Client
	Evm  Client
	Opb  Client
	Gov  Client
}

func NewMsgClient() MsgClient {
	codec.MakeEncodingConfig()
	return MsgClient{
		Tibc: tibc.NewClient(),
		Evm:  evm.NewClient(),
		Opb:  opb.NewClient(),
		Gov:  gov.NewClient(),
	}
}
