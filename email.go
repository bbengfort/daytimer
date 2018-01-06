package daytimer

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
)

//===========================================================================
// Send Agenda
//===========================================================================

// Send the agenda to the specified email address, loading the email config
// from the configuration directory and sending the HTML message.
func (a *Agenda) Send(to string) error {

	// Load email configuration
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	// Create Agenda HTML message
	buffer := new(bytes.Buffer)
	template := MustLoadTemplate("templates/agenda.html")
	if err := template.Execute(buffer, &a); err != nil {
		return err
	}

	// Create new email to send
	email := NewEmail(a.Title, buffer.String(), config.Email)
	return email.Send([]string{to})
}

//===========================================================================
// Email Message
//===========================================================================

// Email message allows you to create and send SMTP requests.
type Email struct {
	From    string
	To      string
	Subject string
	Body    string
	config  *SMTPConfig
}

// NewEmail creates a new email message
func NewEmail(subject string, body string, config *SMTPConfig) *Email {
	return &Email{
		From:    fmt.Sprintf("Daytimer Agenda <%s>", config.From),
		Subject: subject,
		Body:    body,
		config:  config,
	}
}

// Send the message using the config.
func (e *Email) Send(to []string) error {
	e.To = strings.Join(to, ",")

	// Generate email template
	buffer := new(bytes.Buffer)
	template := MustLoadTemplate("templates/email.txt")
	if err := template.Execute(buffer, &e); err != nil {
		return err
	}

	err := smtp.SendMail(
		e.config.Addr(),
		e.config.Auth(),
		e.config.From, to,
		buffer.Bytes(),
	)

	return err
}

//===========================================================================
// SMTP Configuration
//===========================================================================

// SMTPConfig loads the email configuration from JSON.
type SMTPConfig struct {
	Host     string `json:"host"`     // SMTP Hostname
	Port     int    `json:"port"`     // SMTP Port usually 25, 465, or 587
	User     string `json:"user"`     // Username for SMTP authentication
	From     string `json:"from"`     // From email address (must be valid)
	Password string `json:"password"` // Password for SMTP authentication
	UseTLS   bool   `json:"use_tls"`  // Use TLS for secure email
}

// Auth creates an SMTP authentication struct
func (c *SMTPConfig) Auth() smtp.Auth {
	return smtp.PlainAuth("", c.User, c.Password, c.Host)
}

// Addr returns the SMTP server address combining host and port
func (c *SMTPConfig) Addr() string {
	if c.Port > 0 {
		return fmt.Sprintf("%s:%d", c.Host, c.Port)
	}
	return c.Host
}
