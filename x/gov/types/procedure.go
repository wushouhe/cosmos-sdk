package gov

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/go-crypto"
	"github.com/tendermint/go-wire"
)

//-----------------------------------------------------------

// Procedure
type Procedure struct {
	VotingPeriod      int64             //  Length of the voting period. Initial value: 2 weeks
	MinDeposit        int64             //  Minimum deposit for a proposal to enter voting period.
	OptionSet         []string          //  Options available to voters. {Yes, No, NoWithVeto, Abstain}
	ProposalTypes     []string          //  Types available to submitters. {PlainTextProposal, SoftwareUpgradeProposal}
	Threshold         rational.Rational //  Minimum propotion of Yes votes for proposal to pass. Initial value: 0.5
	Veto              rational.Rational //  Minimum value of Veto votes to Total votes ratio for proposal to be vetoed. Initial value: 1/3
	MaxDepositPeriod  int64             //  Maximum period for Atom holders to deposit on a proposal. Initial value: 2 months
	GovernancePenalty int64             //  Penalty if validator does not vote
	IsActive          bool              //  If true, procedure is active. Only one procedure can have isActive true.
}

func NewBaseAccountWithAddress(addr crypto.Address) BaseAccount {
	return BaseAccount{
		Address: addr,
	}
}

// Implements sdk.Account.
func (acc BaseAccount) Get(key interface{}) (value interface{}, err error) {
	panic("not implemented yet")
}

// Implements sdk.Account.
func (acc *BaseAccount) Set(key interface{}, value interface{}) error {
	panic("not implemented yet")
}

// Implements sdk.Account.
func (acc BaseAccount) GetAddress() crypto.Address {
	return acc.Address
}

// Implements sdk.Account.
func (acc *BaseAccount) SetAddress(addr crypto.Address) error {
	if len(acc.Address) != 0 {
		return errors.New("cannot override BaseAccount address")
	}
	acc.Address = addr
	return nil
}

// Implements sdk.Account.
func (acc BaseAccount) GetPubKey() crypto.PubKey {
	return acc.PubKey
}

// Implements sdk.Account.
func (acc *BaseAccount) SetPubKey(pubKey crypto.PubKey) error {
	if acc.PubKey != nil {
		return errors.New("cannot override BaseAccount pubkey")
	}
	acc.PubKey = pubKey
	return nil
}

// Implements sdk.Account.
func (acc *BaseAccount) GetCoins() sdk.Coins {
	return acc.Coins
}

// Implements sdk.Account.
func (acc *BaseAccount) SetCoins(coins sdk.Coins) error {
	acc.Coins = coins
	return nil
}

// Implements sdk.Account.
func (acc *BaseAccount) GetSequence() int64 {
	return acc.Sequence
}

// Implements sdk.Account.
func (acc *BaseAccount) SetSequence(seq int64) error {
	acc.Sequence = seq
	return nil
}

//----------------------------------------
// Wire

func RegisterWireBaseAccount(cdc *wire.Codec) {
	// Register crypto.[PubKey,PrivKey,Signature] types.
	crypto.RegisterWire(cdc)
}
