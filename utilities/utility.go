package utilities

import (
	"os/exec"
	"strings"
	"fmt"
)

const (
	kubectlVersionCommand = " kubectl version"
)

// helper function to check if kubectl present in your system or not
func IsKubePresent() bool{
	kubectlVersionCommandExec := exec.Command(kubectlVersionCommand)
	output, err := kubectlVersionCommandExec.Output()
	
	if err != nil {
		fmt.Println("kubectl not installed on this system")
		return false
	}
	return strings.Contains(string(output),"Client Version:")
}

