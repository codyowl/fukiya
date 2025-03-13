package utilities

import (
	"fmt"
	"os"
	"encoding/json"
	"bufio"
	"strings"
)

const emailConfigFilePath = "fukiya/email_config.json"

// email config struct
type EmailConfig struct {
	SMTPServer	string `json:"smtp_server"`
	SMTPPort	int   `json:"smtp_port"`
	SenderEmail string `json:"sender_email"`
	SenderPass string `json:"sender_pass"`
	EmailEnabled bool   `json:"email_enabled"`
}

// helper function to get user input
func GetUserInput(prompt string)string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// helper function to configure email and save it in json
func ConfigureEmail() error {
	// email config struct
	config := EmailConfig{}
	config.SMTPServer = GetUserInput("SMTP SERVER details : ex gmail, outlook")
	config.SMTPPort = 587
	config.SenderEmail = GetUserInput("Sender Email: ")
	config.SenderPass = GetUserInput("Sender Email Password: ")
	config.EmailEnabled = true

	file, err := os.Create(emailConfigFilePath)
	if err != nil {
		return err
	}
	defer file.close()

	// json encoder to cosume the struct
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		return err
	}

	fmt.Println("Email configuration saved !!")

}

