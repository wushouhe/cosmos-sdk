package gov

import (
	"encoding/json"
	"fmt"
)

type DepositMsg struct {
	ProposalId    int64          `json:"proposal_id"`    // ID of the proposal
	Depositer     crypto.address `json:"depositer"`      // Address of the depositer
	DepositAmount sdk.Coins      `json:"deposit_amount"` // Coins to add to the proposal's deposit
}

func NewDepositMsgMsg(proposalId int64, depositAmount Deposit) DepositMsg {
	return DepositMsg{
		ProposalId:    proposalId,
		DepositAmount: depositAmount,
	}
}

// Implements Msg.
func (msg DepositMsg) Type() string { return "gov" }

// Implements Msg.
func (msg DepositMsg) ValidateBasic() sdk.Error {
	if len(msg.Depositer) == 0 {
		return ErrInvalidAddress(msg.Depositer.String())
	}
	if !in.Coins.IsValid() {
		return ErrInvalidCoins(in.Coins.String())
	}
	if !in.Coins.IsPositive() {
		return ErrInvalidCoins(in.Coins.String())
	}
	return nil
}

func (msg DepositMsg) String() string {
	return fmt.Sprintf("DepositMsg{%v=>%v: %v}", msg.Depositer, msg.ProposerId, msg.DepositAmount)
}

// Implements Msg.
func (msg DepositMsg) Get(key interface{}) (value interface{}) {
	return nil
}

// Implements Msg.
func (msg SubmitProposalMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg) // XXX: ensure some canonical form
	if err != nil {
		panic(err)
	}
	return b
}

// Implements Msg.
func (msg SubmitProposalMsg) GetSigners() []crypto.Address {
	return [1]crypto.Address{msg.Depositer}
}
