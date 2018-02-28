package gov

import (
	wire "github.com/tendermint/go-wire"
)

type governanceMapper struct {

	// The (unexposed) key used to access the store from the Context.
	key sdk.StoreKey

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

// Implements sdk.AccountMapper.
func (am accountMapper) GetProposal(ctx sdk.Context, addr crypto.Address) sdk.Account {
	store := ctx.KVStore(am.key)
	bz := store.Get(addr)
	if bz == nil {
		return nil
	}
	acc := am.decodeAccount(bz)
	return acc
}

// Implements sdk.AccountMapper.
func (am accountMapper) SetAccount(ctx sdk.Context, acc sdk.Account) {
	addr := acc.GetAddress()
	store := ctx.KVStore(am.key)
	bz := am.encodeAccount(acc)
	store.Set(addr, bz)
}
