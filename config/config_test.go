package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Criar arquivo de configuração temporário para teste
	tempConfig := `
directory: "/tmp/test"
days_threshold: 7
schedule:
  enabled: false
  cron: "0 0 * * *"
smtp:
  host: "smtp.example.com"
  port: 587
  username: "test@example.com"
  password: "password"
  to: "to@example.com"
`
	tempFile := "test_config.yaml"
	err := os.WriteFile(tempFile, []byte(tempConfig), 0644)
	if err != nil {
		t.Fatalf("Não foi possível criar arquivo de configuração de teste: %v", err)
	}
	defer os.Remove(tempFile)

	// Testar carregamento de configuração
	cfg, err := LoadConfig(tempFile)
	if err != nil {
		t.Fatalf("Erro ao carregar configuração: %v", err)
	}

	// Verificar valores
	if cfg.Directory != "/tmp/test" {
		t.Errorf("Directory incorreto, esperado '/tmp/test', obtido '%s'", cfg.Directory)
	}
	if cfg.DaysThreshold != 7 {
		t.Errorf("DaysThreshold incorreto, esperado 7, obtido %d", cfg.DaysThreshold)
	}
	if cfg.Schedule.Enabled != false {
		t.Errorf("Schedule.Enabled incorreto, esperado false, obtido %v", cfg.Schedule.Enabled)
	}
	if cfg.Schedule.Cron != "0 0 * * *" {
		t.Errorf("Schedule.Cron incorreto, esperado '0 0 * * *', obtido '%s'", cfg.Schedule.Cron)
	}
	if cfg.SMTP.Host != "smtp.example.com" {
		t.Errorf("SMTP.Host incorreto, esperado 'smtp.example.com', obtido '%s'", cfg.SMTP.Host)
	}
	if cfg.SMTP.Port != 587 {
		t.Errorf("SMTP.Port incorreto, esperado 587, obtido %d", cfg.SMTP.Port)
	}
	if cfg.SMTP.Username != "test@example.com" {
		t.Errorf("SMTP.Username incorreto, esperado 'test@example.com', obtido '%s'", cfg.SMTP.Username)
	}
	if cfg.SMTP.Password != "password" {
		t.Errorf("SMTP.Password incorreto, esperado 'password', obtido '%s'", cfg.SMTP.Password)
	}
	if cfg.SMTP.To != "to@example.com" {
		t.Errorf("SMTP.To incorreto, esperado 'to@example.com', obtido '%s'", cfg.SMTP.To)
	}
}

func TestLoadConfigError(t *testing.T) {
	// Testar carregamento de configuração com arquivo inexistente
	_, err := LoadConfig("arquivo_inexistente.yaml")
	if err == nil {
		t.Error("Deveria retornar erro ao tentar carregar arquivo inexistente")
	}

	// Criar arquivo de configuração inválido
	invalidConfig := `
directory: "/tmp/test"
days_threshold: "not_a_number" # Isso deve causar erro
`
	invalidFile := "invalid_config.yaml"
	err = os.WriteFile(invalidFile, []byte(invalidConfig), 0644)
	if err != nil {
		t.Fatalf("Não foi possível criar arquivo de configuração inválido: %v", err)
	}
	defer os.Remove(invalidFile)

	// Testar carregamento de configuração inválida
	_, err = LoadConfig(invalidFile)
	if err == nil {
		t.Error("Deveria retornar erro ao tentar carregar configuração inválida")
	}
}
