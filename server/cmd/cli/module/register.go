package module

import "github.com/spf13/cobra"

func RegisterCommand(parent *cobra.Command) {
	gCmd := newGenerateCmd()
	rCmd := newRemoveCmd()
	parent.AddCommand(gCmd)
	parent.AddCommand(rCmd)
}
