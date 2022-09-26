package msgs

import (
	opb "github.com/bianjieai/irita/modules/opb/types"
	evm "github.com/tharsis/ethermint/x/evm/types"
)

const (
	MsgTypeEthereumTx = "ethereum_tx"

	MsgTypeMint    = "mint"
	MsgTypeReclaim = "reclaim"
)

type (
	//evm
	MsgEthereumTx = evm.MsgEthereumTx

	//opb
	MsgMint    = opb.MsgMint
	MsgReclaim = opb.MsgReclaim
)
