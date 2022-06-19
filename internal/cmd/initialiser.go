package cmd

func InitialiseRootCmd() *RootCmd {
	rootCmd := NewRootCmd()
	deployCmd := NewDeployCmd()
	rootCmd.Command.AddCommand(
		deployCmd.command)
	return rootCmd
}
