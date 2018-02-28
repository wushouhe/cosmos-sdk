
import (
	"encoding/json"
	"fmt"
)

type ClaimDepositMsg struct {
	ProposalID int64
}

func NewClaimDepositMsg(proposalId int64) SendMsg {
	return ClaimDepositMsg{
		ProposalId: proposalId,
	}
}

// Implements Msg.
func (msg ClaimDepositMsg) Type() string { return "gov" } // TODO: "gov/claimdeposit"

// Implements Msg.
func (msg ClaimDepositMsg) ValidateBasic() sdk.Error {
	// TODO: Add validation logic
	return nil
}

func (msg ClaimDepositMsg) String() string {
	return fmt.Sprintf("ClaimDepositMsg{%v}", msg.ProposalId)
}

// Implements Msg.
func (msg ClaimDepositMsg) Get(key interface{}) (value interface{}) {
	return nil
}

// Implements Msg.
func (msg ClaimDepositMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg) // XXX: ensure some canonical form
	if err != nil {
		panic(err)
	}
	return b
}

// Implements Msg.
func (msg ClaimDepositMsg) GetSigners() []crypto.Address {
	addrs := make([]crypto.Address, len(msg.Inputs))
	for i, in := range msg.Inputs {
		addrs[i] = in.Address
	}
	return addrs
}