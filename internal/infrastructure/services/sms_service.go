package services

import (
	"fmt"
	"log"
	"os"
)

type SMSService struct {
	apiKey    string
	apiSecret string
	from      string
	baseURL   string
}

func NewSMSService() *SMSService {
	return &SMSService{
		apiKey:    os.Getenv("SMS_API_KEY"),
		apiSecret: os.Getenv("SMS_API_SECRET"),
		from:      os.Getenv("SMS_FROM"),
		baseURL:   os.Getenv("SMS_BASE_URL"),
	}
}

func (ss *SMSService) SendPasswordResetSMS(to, token string) error {
	message := fmt.Sprintf("Seu código de verificação é: %s. Expira em 15 minutos.", token)

	// Simulação de envio de SMS - em produção, use uma API real como Twilio, AWS SNS, etc.
	log.Printf("SMS enviado para %s: %s", to, message)

	// Exemplo de implementação com API HTTP (comentado para simulação)
	/*
		data := url.Values{}
		data.Set("to", to)
		data.Set("from", ss.from)
		data.Set("message", message)

		req, err := http.NewRequest("POST", ss.baseURL+"/messages", strings.NewReader(data.Encode()))
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", "Bearer "+ss.apiKey)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("SMS API returned status: %d", resp.StatusCode)
		}
	*/

	return nil
}

func (ss *SMSService) SendWelcomeSMS(to, name string) error {
	message := fmt.Sprintf("Olá %s! Bem-vindo ao nosso sistema!", name)

	log.Printf("SMS de boas-vindas enviado para %s: %s", to, message)

	return nil
}
