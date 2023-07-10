package cli

import "github.com/spf13/cobra"

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "jwtcli",
		Short: "JWT handling in cli",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	rootCmd.AddCommand(createCmd())
	return rootCmd
}
