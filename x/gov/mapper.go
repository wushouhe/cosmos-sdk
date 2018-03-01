package gov

import (
	wire "github.com/tendermint/go-wire"
)

type governanceMapper struct {

	// The (unexposed) key used to access the store from the Context.
	proposalStoreKey sdk.StoreKey

	// The (unexposed) key used to access the store from the Context.
	proposalStoreey sdk.StoreKey

	// The wire codec for binary encoding/decoding of accounts.
	cdc *wire.Codec
}

// NewGovernanceMapper returns a mapper that uses go-wire to (binary) encode and decode gov types.
func NewGovernanceMapper(key sdk.StoreKey, proto sdk.Account) accountMapper {
	cdc := wire.NewCodec()
	return accountMapper{
		key: key,
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
	store := ctx.KVStore(am.key)
	bz := store.Get(proposalId)
	if bz == nil {
		return nil
	}
	proposal := gm.decodeProposal(bz)
	return acc
}

// Implements sdk.AccountMapper.
func (gm governanceMapper) SetProposal(ctx sdk.Context, proposal Proposal) {
	proposalId := proposal.ProposalId
	store := ctx.KVStore(am.key)
	bz := gm.encodeProposal(proposal)
	store.Set(proposalId, bz)
}

func (gm governanceMapper) encodeProposal(proposal Proposal) []byte {
	bz, err := gm.cdc.MarshalBinary(proposal)
	if err != nil {
		panic(err)
	}
	return bz
}

func (gm governanceMapper) decodeProposal(bz []byte) Proposal {
	// TODO

	proposal := am.clonePrototypePtr()
	err := gm.cdc.UnmarshalBinary(bz, proposal)
	if err != nil {
		panic(err)
	}

	return proposal
}
