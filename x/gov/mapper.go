package gov

import (
	"github.com/cosmos/cosmos-sdk/x/bank"
	wire "github.com/tendermint/go-wire"
)

type governanceMapper struct {
	ck bank.CoinKeeper

	// The (unexposed) key used to access the store from the Context.
	proposalStoreKey sdk.StoreKey

	// The (unexposed) key used to access the store from the Context.
	validatorInfoStoreKey sdk.StoreKey

	// The (unexposed) key used to access the store from the Context.
	votesStoreKey sdk.StoreKey

	// The wire codec for binary encoding/decoding of accounts.
	cdc *wire.Codec
}

// NewGovernanceMapper returns a mapper that uses go-wire to (binary) encode and decode gov types.
func NewGovernanceMapper(key sdk.StoreKey, ck bank.CoinKeeper) accountMapper {
	cdc := wire.NewCodec()
	return accountMapper{
		key: key,
		ck:  ck,
		cdc: cdc,
	}
}

// Returns the go-wire codec.  You may need to register interfaces
// and concrete types here, if your app's sdk.Account
// implementation includes interface fields.
// NOTE: It is not secure to expose the codec, so check out
// .Seal().
func (gm governanceMapper) WireCodec() *wire.Codec {
	return gm.cdc
}

func (gm governanceMapper) GetProposal(ctx sdk.Context, proposalId int64) sdk.Account {
	store := ctx.KVStore(am.proposalStoreKey)
	bz := store.Get(proposalId)
	if bz == nil {
		return nil
	}

	proposal := Proposal{}
	err := gm.cdc.UnmarshalBinary(bz, proposal)
	if err != nil {
		panic(err)
	}

	return acc
}

// Implements sdk.AccountMapper.
func (gm governanceMapper) SetProposal(ctx sdk.Context, proposal Proposal) {
	proposalId := proposal.ProposalId
	store := ctx.KVStore(am.proposalStoreKey)

	bz, err := gm.cdc.MarshalBinary(proposal)
	if err != nil {
		panic(err)
	}

	store.Set(proposalId, bz)
}

func (gm governanceMapper) GetValidatorInfo(ctx sdk.Context, proposalId int64, validatorAddr crypto.address) sdk.Account {
	store := ctx.KVStore(am.validatorInfoStoreKey)

	bz := store.Get(proposalId)
	if bz == nil {
		return nil
	}

	vote := Vote{}
	err := gm.cdc.UnmarshalBinary(bz, vote)
	if err != nil {
		panic(err)
	}

	return vote
}

// Implements sdk.AccountMapper.
func (gm governanceMapper) SetVote(ctx sdk.Context, vote Vote) {
	proposalId := proposal.ProposalId
	store := ctx.KVStore(am.votesStoreKey)

	bz, err := gm.cdc.MarshalBinary(vote)
	if err != nil {
		panic(err)
	}

	store.Set(proposalId, bz)
}

func (gm governanceMapper) GetVote(ctx sdk.Context, proposalId int64, voter crypto.address) sdk.Account {
	store := ctx.KVStore(am.votesStoreKey)
	bz := store.Get(proposalId)
	if bz == nil {
		return nil
	}

	vote := Vote{}
	err := gm.cdc.UnmarshalBinary(bz, vote)
	if err != nil {
		panic(err)
	}

	return vote
}

// Implements sdk.AccountMapper.
func (gm governanceMapper) SetVote(ctx sdk.Context, vote Vote) {
	proposalId := proposal.ProposalId
	store := ctx.KVStore(am.votesStoreKey)

	bz, err := gm.cdc.MarshalBinary(vote)
	if err != nil {
		panic(err)
	}

	store.Set(proposalId, bz)
}
