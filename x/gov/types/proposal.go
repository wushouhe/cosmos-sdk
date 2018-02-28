package gov

//-----------------------------------------------------------

// Proposal
type Proposal struct {
	Title                string //  Title of the proposal
	Description          string //  Description of the proposal
	Type                 string //  Type of proposal. Initial set {PlainTextProposal, SoftwareUpgradeProposal}
	Deposit              int64  //  Current deposit on this proposal. Initial value is set at InitialDeposit
	SubmitBlock          int64  //  Height of the block where TxGovSubmitProposal was included
	VotingStartBlock     int64  //  Height of the block where MinDeposit was reached. -1 if MinDeposit is not reached
	InitTotalVotingPower int64  //  Total voting power when proposal enters voting period (default 0)
	InitProcedureNumber  int16  //  Procedure number of the active procedure when proposal enters voting period (default -1)
}
