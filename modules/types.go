package msgs

import (
	opb "github.com/bianjieai/irita/modules/opb/types"
	tibctranfer "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/types"
	tibcclient "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	tibcpacket "github.com/bianjieai/tibc-go/modules/tibc/core/04-packet/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	evm "github.com/tharsis/ethermint/x/evm/types"
)

const (
	//tibc
	MsgTypeTIBCNftTransfer     = "tibc_nft_transfer"
	MsgTypeTIBCRecvPacket      = "tibc_recv_packet"
	MsgTypeTIBCUpdateClient    = "tibc_update_client"
	MsgTypeTIBCAcknowledgement = "tibc_acknowledge_packet"
	MsgTypeTIBCCleanPacket     = "clean_packet"
	MsgTypeTIBCRecvCleanPacket = "recv_clean_packet"

	MsgTypeEthereumTx = "ethereum_tx"

	MsgTypeMint    = "mint"
	MsgTypeReclaim = "reclaim"

	MsgTypeSubmitProposal = "submit_proposal"
	MsgTypeDeposit        = "deposit"
	MsgTypeVote           = "vote"
	MsgTypeVoteWeighted   = "vote_weighted"
)

type (
	//tibc
	TIBCAcknowledgement        = tibcpacket.Acknowledgement
	NonFungibleTokenPacketData = tibctranfer.NonFungibleTokenPacketData
	MsgTIBCNftTransfer         = tibctranfer.MsgNftTransfer
	MsgTIBCUpdateClient        = tibcclient.MsgUpdateClient
	MsgTIBCRecvPacket          = tibcpacket.MsgRecvPacket
	MsgTIBCAcknowledgement     = tibcpacket.MsgAcknowledgement
	MsgCleanPacket             = tibcpacket.MsgCleanPacket
	MsgRecvCleanPacket         = tibcpacket.MsgRecvCleanPacket

	//evm
	MsgEthereumTx = evm.MsgEthereumTx

	//opb
	MsgMint    = opb.MsgMint
	MsgReclaim = opb.MsgReclaim

	//gov
	MsgDeposit        = gov.MsgDeposit
	MsgSubmitProposal = gov.MsgSubmitProposal
	TextProposal      = gov.TextProposal
	MsgVote           = gov.MsgVote
	Proposal          = gov.Proposal
	SdkVote           = gov.Vote
	GovContent        = gov.Content
	MsgVoteWeighted   = gov.MsgVoteWeighted
)
