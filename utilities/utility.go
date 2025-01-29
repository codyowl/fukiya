package utilities

import (
	"os/exec"
	"strings"
	"fmt"
	v1 "k8s.io/api/core/v1"
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



// helper function to send email
func SendEmailAlert(pod *v1.Pod){
	fmt.Printf("sending email alert: Pod %s belongs to the namespace %s is down!\n", pod.Name, pod.Namespace)
}
