package cli

import (
    "fmt"
	"github.com/spf13/cobra"
)


// function to watch all pods logic
func WatchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "watch",
		Short: "Watch all pods in all namespaces",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("This will watch all the pods create a function and import it here")
			return nil
		},
	}
}