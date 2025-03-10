package cli

import (
    "fmt"
	"github.com/spf13/cobra"
	"fukiya/utilities"
)


// function to watch all pods logic
func WatchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "watch",
		Short: "Watch all pods in all namespaces",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("This will watch all the pods create a function and import it here")
			// running in detached state
			utilities.RunInBackground()
			fmt.Println("Pod monitoring started successfully!")
			return nil
		},		
	}
}

// function to config email and phone 
func EmailSetupCmd() *cobra.Command{
	return &cobra.Command{
		Use:   "email-setup",
		Short: "email notifications setup for pod monitoring alerts",
		Run: func(cmd *cobra.Command, args []string) {
			err := utilities.ConfigureEmail()
			if err != nil {
				fmt.Println("❌ Email setup failed:", err)
			}
		},
	}
}