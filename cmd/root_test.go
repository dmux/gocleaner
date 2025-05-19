package cmd

import (
	"gocleaner/config"
	"gocleaner/core"
	"os"
	"testing"
)

// Variáveis de pacote para mockagem de funções
var (
	originalLoadConfig config.LoadConfigFunc = config.LoadConfig
	originalRunCleaner core.RunCleanerFunc = core.RunCleaner
)

// TestExecute testa a execução do comando principal
func TestExecute(t *testing.T) {
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
	tempFile := "test_cmd_config.yaml"
	err := os.WriteFile(tempFile, []byte(tempConfig), 0644)
	if err != nil {
		t.Fatalf("Não foi possível criar arquivo de configuração de teste: %v", err)
	}
	defer os.Remove(tempFile)

	// Mockando as funções que seriam chamadas
	configCalled := false
	mockLoadConfig := func(path string) (*config.Config, error) {
		configCalled = true
		// Verificar se o caminho está correto
		if path != "config.yaml" {
			t.Errorf("Caminho de configuração incorreto, esperado 'config.yaml', obtido '%s'", path)
		}
		return &config.Config{
			Directory:     "/tmp/test",
			DaysThreshold: 7,
			Schedule: config.Schedule{
				Enabled: false,
				Cron:    "0 0 * * *",
			},
		}, nil
	}

	// Mock para RunCleaner
	cleanerCalled := false
	mockRunCleaner := func(cfg *config.Config) {
		cleanerCalled = true
		// Verificar se a configuração está correta
		if cfg.Directory != "/tmp/test" {
			t.Errorf("Directory incorreto, esperado '/tmp/test', obtido '%s'", cfg.Directory)
		}
		if cfg.DaysThreshold != 7 {
			t.Errorf("DaysThreshold incorreto, esperado 7, obtido %d", cfg.DaysThreshold)
		}
	}

	// Substituir função original
	config.LoadConfig = config.LoadConfigFunc(mockLoadConfig)
	core.RunCleaner = mockRunCleaner

	// Restaurar as funções originais ao final
	defer func() {
		config.LoadConfig = originalLoadConfig
		core.RunCleaner = originalRunCleaner
	}()

	// Testar execução com agendamento desabilitado
	Execute()

	// Verificar se as funções foram chamadas
	if !configCalled {
		t.Error("LoadConfig não foi chamada")
	}
	if !cleanerCalled {
		t.Error("RunCleaner não foi chamada")
	}
}
