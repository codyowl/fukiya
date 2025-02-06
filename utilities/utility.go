package utilities

import (
	"os/exec"
	"strings"
	"fmt"
	"os"
	"syscall"
	v1 "k8s.io/api/core/v1"
)

const (
	kubectlCmd = "kubectl" 
	pidFilePath = "/tmp/fukiya_watch.pid"
	logFilePath = "/tmp/fukiya.log"
)

// kubectl sub commands
var VersionCmd = []string{"version", "--client"}


// helper function to run the process in detached mode
func RunInBackground(){
	cmd := exec.Command(os.Args[0], "watch", "--daemon")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	// saving the logs in a file 
	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	cmd.Stdout = logFile
	cmd.Stderr = logFile

	err = cmd.Start()
	if err != nil{
		fmt.Println("Failed to start monitoring:", err)
		return
	}

	// writing process id to a file 
	err = os.WriteFile(pidFilePath, []byte(fmt.Sprintf("%d", cmd.Process.Pid)), 0644)
	if err != nil {
		fmt.Println("Failed to write PID file:", err)
	}
	fmt.Printf("Pod monitoring started in background (PID: %d)\n", cmd.Process.Pid)
	os.Exit(0)

}

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
