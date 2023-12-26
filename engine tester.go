func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the redeemInfo
	for _, elem := range genState.RedeemInfoList {
		k.SetRedeemInfo(ctx, elem)
	}


	// Set googlIAPOrder count
	k.SetGoogleIAPOrderCount(ctx, genState.GoogleIapOrderCount)

	// Set all the execution
	for _, elem := range genState.ExecutionList {
		k.SetExecution(ctx, elem)
	}

	// Set all the item
	for _, elem := range genState.ItemList {
		k.SetItem(ctx, elem)
	}
