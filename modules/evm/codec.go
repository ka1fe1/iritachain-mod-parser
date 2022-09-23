package evm

import (
	"github.com/kaifei-bianjie/common-parser/codec"
	"github.com/tharsis/ethermint/x/evm"
)

func init() {
	codec.RegisterAppModules(
		evm.AppModuleBasic{},
	)
}
