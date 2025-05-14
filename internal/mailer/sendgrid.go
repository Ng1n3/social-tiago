package mailer

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"time"

	sib "github.com/sendinblue/APIv3-go-library/v2/lib"
)

type BrevoMailer struct {
	fromEmail string
	// apiKey    string
	client *sib.APIClient
}

func NewBrevo(apiKey, fromEmail string) *BrevoMailer {
	cfg := sib.NewConfiguration()
	cfg.AddDefaultHeader("api-key", apiKey)
	client := sib.NewAPIClient(cfg)

	return &BrevoMailer{
		fromEmail: fromEmail,
		client:    client,
	}
}

func (m *BrevoMailer) Send(templateFile, username, email string, data any, isSandbox bool) (int, error) {


	// template parsing and building
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return -1, fmt.Errorf("failed to parse template: %w", err)
	}

	// Render subject
	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return -1, fmt.Errorf("failed to render subject: %w", err)
	}

	//Render HTML body
	body := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		return -1, fmt.Errorf("failed to render body: %w", err)
	}

	emailRequest := sib.SendSmtpEmail{
		Sender: &sib.SendSmtpEmailSender{
			Name:  FromName,
			Email: m.fromEmail,
		},
		To: []sib.SendSmtpEmailTo{
			{
				Email: email,
				Name:  username,
			},
		},
		Subject:     subject.String(),
		HtmlContent: body.String(),
	}

	if isSandbox {
		log.Printf("[SANDBOX] Would send email to %s with subject: %s", email, subject.String())
	}

	var lastErr error
	for i := 0; i < maxRetries; i++ {
		_, response, lastErr := m.client.TransactionalEmailsApi.SendTransacEmail(context.Background(), sib.SendSmtpEmail(emailRequest))
		if lastErr != nil {
			time.Sleep(time.Second * time.Duration(i+1)) // exponential backoff
			continue
		}
		return response.StatusCode, nil

	}

	return -1,  fmt.Errorf("failed to send email after %d attempts: %w", maxRetries, lastErr)

}
