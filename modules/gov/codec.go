package gov

import (
	spartangov "github.com/bianjieai/spartan-cosmos/module/gov/module"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		distribution.AppModuleBasic{},
		spartangov.AppModuleBasic{},
	)
}
