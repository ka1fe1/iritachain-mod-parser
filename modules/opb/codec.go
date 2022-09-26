package opb

import (
	"github.com/bianjieai/irita/modules/opb"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		opb.AppModuleBasic{},
	)
}
