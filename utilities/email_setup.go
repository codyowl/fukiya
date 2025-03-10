package utilities

import (
	"fmt"
	"os"
)

const emailConfigFilePath = "fukiya/email_config.json"

// email config struct
type emailConfig struct {
	SMTPServer	string `json:"smtp_server"`
	SMTPPort	int   `json:"smtp_port"`
	SenderEmail string `json:"sender_email"`
	SenderPass string `json: "sender_pass"`
}

// helper function to get user input
func GetUserInput(prompt string)string {

}

// helper function to configure email and save it in json
func ConfigureEmail() error {

}

