package keeper

import (
	"hictest/x/dex/types"
)

var _ types.QueryServer = Keeper{}
