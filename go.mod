module github.com/kaifei-bianjie/iritachain-mod-parser

go 1.16

require (
	github.com/bianjieai/irita v0.1.1-rc1.0.20220809071819-204fd6357364
	github.com/bianjieai/spartan-cosmos v1.0.2
	github.com/bianjieai/tibc-go v0.3.1-0.20220412124234-9b75094ff868
	github.com/cosmos/cosmos-sdk v0.45.5-0.20220523154235-2921a1c3c918
	github.com/kaifei-bianjie/common-parser v0.0.0-20220923023138-65dfc81a8ff5
	github.com/stretchr/testify v1.7.5
	github.com/tharsis/ethermint v0.10.3
)

replace (
	github.com/CosmWasm/wasmd => github.com/bianjieai/wasmd v0.19.1-0.20211215102105-45e28c7c896c
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.45.1-irita-20220816.0.20220816095307-845547d9c19e
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.8-irita-210413.0.20211012090339-cee6e09e8ae3
	github.com/tharsis/ethermint => github.com/bianjieai/ethermint v0.10.2-irita-20220826
)
