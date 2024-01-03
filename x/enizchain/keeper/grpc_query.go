package keeper

import (
	"enizchain/x/enizchain/types"
)

var _ types.QueryServer = Keeper{}
