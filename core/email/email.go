package email

import (
	"fmt"
	"net/smtp"
	"strings"

	"gocleaner/config"

	mail "github.com/jordan-wright/email"
)

func SendReport(cfg *config.Config, items []string) error {
	e := mail.NewEmail()
	e.From = fmt.Sprintf("Limpador <%s>", cfg.SMTP.Username)
	e.To = []string{cfg.SMTP.To}
	e.Subject = "Relatório de exclusão de arquivos e pastas"

	if len(items) == 0 {
		e.Text = []byte("Nenhum item foi excluído.")
	} else {
		body := "Os seguintes arquivos e diretórios foram excluídos:\n\n"
		body += strings.Join(items, "\n")
		e.Text = []byte(body)
	}

	return e.Send(fmt.Sprintf("%s:%d", cfg.SMTP.Host, cfg.SMTP.Port),
		smtp.PlainAuth("", cfg.SMTP.Username, cfg.SMTP.Password, cfg.SMTP.Host))
}
