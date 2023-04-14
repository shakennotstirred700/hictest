package hictest_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "hictest/testutil/keeper"
	"hictest/testutil/nullify"
	"hictest/x/hictest"
	"hictest/x/hictest/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HictestKeeper(t)
	hictest.InitGenesis(ctx, *k, genesisState)
	got := hictest.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
