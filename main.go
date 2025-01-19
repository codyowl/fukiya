package main

import (
	"fmt"
	"fukiya/utilities"
	"fukiya/cli"
	"github.com/spf13/cobra"
)

// cobra root command
var cobrarootCmd = &cobra.Command{
	Use:     "Fukiya",
	Version: "0.1.0",
	Short:   "A hassle free CLI for DEVOPS to make their life easier",
}


func main(){
	cobrarootCmd.AddCommand(cli.WatchCmd())

	if utilities.IsKubePresent(){
		fmt.Println("Yes kubectl is present")
		
		// 
		if err := cobrarootCmd.Execute(); err != nil {
			fmt.Println(err)
		}


	}else{
		fmt.Println("No it is not here")
	}
}