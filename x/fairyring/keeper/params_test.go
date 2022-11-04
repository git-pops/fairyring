package keeper_test

import (
	"testing"

	testkeeper "fairyring/testutil/keeper"
	"fairyring/x/fairyring/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FairyringKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}