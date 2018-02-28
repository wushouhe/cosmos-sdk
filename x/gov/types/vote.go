import crypto "github.com/tendermint/go-crypto"

type Vote struct {
	ProposalID      int64         //  proposalID of the proposal
	Option          string        //  option from OptionSet chosen by the voter
	ValidatorPubKey crypto.PubKey //  PubKey of the validator voter wants to tie its vote to
}