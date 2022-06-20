package cmd

func InitialiseRootCmd() *RootCmd {
	rootCmd := NewRootCmd()
	deployCmd := NewDeployCmd()
	upgradeCmd := NewUpgradeCmd()
	uninstallCmd := NewUninstallCmd()
	rootCmd.Command.AddCommand(
		deployCmd.command,
		upgradeCmd.command,
		uninstallCmd.command,
	)
	return rootCmd
}
