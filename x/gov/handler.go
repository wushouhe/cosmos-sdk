package gov

import (
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Handle all "gov" type messages.
func NewHandler(gm governanceMapper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case DepositMsg:
			return handleDepositMsg(ctx, gm, msg)
		case ClaimDepositMsg:
			return handleClaimDepositMsg(ctx, gm, msg)
		case SubmitProposalMsg:
			return handleSubmitProposalMsg(ctx, gm, msg)
		case VoteMsg:
			return handleVoteMsg(ctx, gm, msg)
		default:
			errMsg := "Unrecognized gov Msg type: " + reflect.TypeOf(msg).Name()
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle SendMsg.
func handleDepositMsg(ctx sdk.Context, gm GovernanceMapper, msg DepositMsg) sdk.Result {

	proposal := gm.getProposal(ctx, msg.ProposalId)

	if proposal == nil {
		// throw
	}

	if ctx.isCheckTx() {
		return sdk.Result{} // TODO
	}

	return sdk.Result{} // TODO
}

// Handle SendMsg.
func handleClaimDepositMsg(ctx sdk.Context, gm GovernanceMapper, msg ClaimDepositMsg) sdk.Result {
	return sdk.Result{} // TODO
}

// Handle SendMsg.
func handleSubmitProposalMsg(ctx sdk.Context, gm GovernanceMapper, msg SubmitProposalMsg) sdk.Result {

	proposal := Proposal{
		Title:                msg.Title,
		Description:          msg.Description,
		ProposalType:         msg.ProposalType,
		Deposit:              msg.InitialDeposit,
		SubmitBlock:          ctx.BlockHeight(),
		VotingStartBlock:     -1,
		InitTotalVotingPower: 0,         // TODO: Get total voting power from stake kvstore
		Procedure:            Procedure, // TODO: Get active Procedure from params kvstore
	}

	return sdk.Result{} // TODO
}

// Handle SendMsg.
func handleVoteMsg(ctx sdk.Context, ck CoinKeeper, msg SendMsg) sdk.Result {
	return sdk.Result{} // TODO
}
