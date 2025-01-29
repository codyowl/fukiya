package main

import (
	"fmt"
	"context"
	"fukiya/utilities"
	"fukiya/k8s_go_client"
	"fukiya/cli"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

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
		
		
		// testing kube go client
		kubeClientset, err := k8s_go_client.GetKubeConfig()
		if err != nil {
			fmt.Println("FAILED TO LOAD KUBECONFIG: %v", err)
		}

		// getting pod details for default namespace for now
		pods, err := kubeClientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			fmt.Println("FAILED TO GET PODS: %v", err)		
		}

		// checking pod status
		fmt.Println("Checking pod status master bruce ")
		k8s_go_client.GetPodsStatus(kubeClientset)

		for _, pod := range pods.Items {
			fmt.Println("fetched pod name:", pod.Name)
		}

		if err := cobrarootCmd.Execute(); err != nil {
			fmt.Println(err)
		}


	}else{
		fmt.Println("No it is not here")
	}
}