package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

type EmailService struct {
	from     string
	password string
	smtpHost string
	smtpPort string
}

func NewEmailService() *EmailService {
	return &EmailService{
		from:     os.Getenv("EMAIL_FROM"),
		password: os.Getenv("EMAIL_PASSWORD"),
		smtpHost: os.Getenv("SMTP_HOST"),
		smtpPort: os.Getenv("SMTP_PORT"),
	}
}

func (es *EmailService) SendPasswordResetEmail(to, name, token string) error {
	subject := "Recuperação de Senha"
	body := fmt.Sprintf(`
		<html>
		<body>
			<h2>Olá %s!</h2>
			<p>Você solicitou a recuperação de senha da sua conta.</p>
			<p>Seu código de verificação é: <strong>%s</strong></p>
			<p>Este código expira em 15 minutos.</p>
			<p>Se você não solicitou esta recuperação, ignore este email.</p>
			<br>
			<p>Atenciosamente,<br>Equipe de Suporte</p>
		</body>
		</html>
	`, name, token)

	message := fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n"+
		"\r\n"+
		"%s\r\n", to, subject, body)

	auth := smtp.PlainAuth("", es.from, es.password, es.smtpHost)
	addr := fmt.Sprintf("%s:%s", es.smtpHost, es.smtpPort)

	err := smtp.SendMail(addr, auth, es.from, []string{to}, []byte(message))
	if err != nil {
		log.Printf("Erro ao enviar email: %v", err)
		return err
	}

	log.Printf("Email de recuperação enviado para: %s", to)
	return nil
}

func (es *EmailService) SendWelcomeEmail(to, name string) error {
	subject := "Bem-vindo!"
	body := fmt.Sprintf(`
		<html>
		<body>
			<h2>Olá %s!</h2>
			<p>Bem-vindo ao nosso sistema!</p>
			<p>Sua conta foi criada com sucesso.</p>
			<br>
			<p>Atenciosamente,<br>Equipe de Suporte</p>
		</body>
		</html>
	`, name)

	message := fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n"+
		"\r\n"+
		"%s\r\n", to, subject, body)

	auth := smtp.PlainAuth("", es.from, es.password, es.smtpHost)
	addr := fmt.Sprintf("%s:%s", es.smtpHost, es.smtpPort)

	err := smtp.SendMail(addr, auth, es.from, []string{to}, []byte(message))
	if err != nil {
		log.Printf("Erro ao enviar email de boas-vindas: %v", err)
		return err
	}

	log.Printf("Email de boas-vindas enviado para: %s", to)
	return nil
}
