package gov

import (
	"github.com/tendermint/go-wire"
)

func RegisterWire(cdc *wire.Codec) {
	// TODO include option to always include prefix bytes.
	cdc.RegisterConcrete(ClaimDepositMsg{}, "cosmos-sdk/ClaimDepositMsg", nil)
	cdc.RegisterConcrete(DepositMsg{}, "cosmos-sdk/DepositMsg", nil)
	cdc.RegisterConcrete(SubmitProposalMsg{}, "cosmos-sdk/SubmitProposalMsg", nil)
	cdc.RegisterConcrete(VoteMsg{}, "cosmos-sdk/VoteMsg", nil)
}
