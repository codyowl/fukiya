package utilities

import (
	"os/exec"
	"strings"
	"fmt"
	"path/filepath"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	kubectlCmd = "kubectl" 
)

// kubectl sub commands
var VersionCmd = []string{"version", "--client"}

// helper function to check if kubectl present in your system or not
func IsKubePresent() bool{
	kubectlVersionCommandExec := exec.Command(kubectlCmd, VersionCmd...)
	output, err := kubectlVersionCommandExec.Output()
	
	
	if err != nil {
		fmt.Println("kubectl not installed on this system")
		return false
	}
	return strings.Contains(string(output),"Client Version:")
}


// helper function to get kubeconfig out of kubeconfig path
func GetKubeConfig()(*kubernetes.Clientset, error){
	kubeconfig := filepath.Join(homeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("error loading kubeconfig: %v", err)
	}

	// creating the kube clientset
	kubeclient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error loading kube client: %v", err)
	}

	return kubeclient, nil
} 

// helper function to get home dir
func homeDir() string {
	return "/root"
}

// helper function to set the pod monitor configuration
func SetConfig(){

}
