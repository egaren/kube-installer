package cmd

func InitialiseRootCmd() *RootCmd {
	rootCmd := NewRootCmd()
	deployCmd := NewDeployCmd()
	upgradeCmd := NewUpgradeCmd()
	rootCmd.Command.AddCommand(
		deployCmd.command,
		upgradeCmd.command,
	)
	return rootCmd
}
