package gov

import (
	"encoding/json"
	"fmt"

	crypto "github.com/tendermint/go-crypto"
)

//----------------------------------------
// SubmitProposalMsg

type SubmitProposalMsg struct {
	Title          string         //  Title of the proposal
	Description    string         //  Description of the proposal
	ProposalType   string         //  Type of proposal. Initial set {PlainTextProposal, SoftwareUpgradeProposal}
	Proposer       crypto.address //  Address of the proposer
	InitialDeposit int64          //  Initial deposit paid by sender. Must be strictly positive.
}

func NewSubmitProposalMsg(title string, description string, proposalType string, initialDeposit int64) SendMsg {
	return SubmitProposalMsg{
		Title:          title,
		Description:    description,
		ProposalType:   proposalType,
		InitialDeposit: initialDeposit,
	}
}

// Implements Msg.
func (msg SubmitProposalMsg) Type() string { return "gov" } // TODO: "gov/submitproposal"

// Implements Msg.
func (msg SubmitProposalMsg) ValidateBasic() sdk.Error {
	// TODO: Add validation logic
	return nil
}

func (msg SubmitProposalMsg) String() string {
	return fmt.Sprintf("SubmitProposalMsg{%v, %v, %v, %v}", msg.Title, msg.Description, msg.ProposalType, msg.InitialDeposit)
}

// Implements Msg.
func (msg SubmitProposalMsg) Get(key interface{}) (value interface{}) {
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
func (msg SendMsg) GetSigners() []crypto.Address {
	addrs := make([]crypto.Address, len(msg.Inputs))
	for i, in := range msg.Inputs {
		addrs[i] = in.Address
	}
	return addrs
}

type DepositMsg struct {
	ProposalID int64 // ID of the proposal
	Deposit    int64 // Number of Atoms to add to the proposal's deposit
}

type ClaimDepositMsg struct {
	ProposalID int64
}

type VoteMsg struct {
	ProposalID      int64         //  proposalID of the proposal
	Option          string        //  option from OptionSet chosen by the voter
	ValidatorPubKey crypto.PubKey //  PubKey of the validator voter wants to tie its vote to
}
