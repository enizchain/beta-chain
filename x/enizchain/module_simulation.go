package enizchain

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"enizchain/testutil/sample"
	invochainsimulation "enizchain/x/enizchain/simulation"
	"enizchain/x/enizchain/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = invochainsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreatePost = "op_weight_msg_post"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePost int = 100

	opWeightMsgUpdatePost = "op_weight_msg_post"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePost int = 100

	opWeightMsgDeletePost = "op_weight_msg_post"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePost int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	invochainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PostList: []types.Post{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		PostCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&invochainGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePost int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreatePost, &weightMsgCreatePost, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePost = defaultWeightMsgCreatePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePost,
		invochainsimulation.SimulateMsgCreatePost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePost int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdatePost, &weightMsgUpdatePost, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePost = defaultWeightMsgUpdatePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePost,
		invochainsimulation.SimulateMsgUpdatePost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePost int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeletePost, &weightMsgDeletePost, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePost = defaultWeightMsgDeletePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePost,
		invochainsimulation.SimulateMsgDeletePost(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
