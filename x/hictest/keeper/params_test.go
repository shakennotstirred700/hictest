package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "hictest/testutil/keeper"
	"hictest/x/hictest/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HictestKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
