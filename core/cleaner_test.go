package core

import (
	"gocleaner/config"
	"gocleaner/internal"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// TestRemoveWithLog testa a função auxiliar removeWithLog
func TestRemoveWithLog(t *testing.T) {
	// Configurar logger para teste
	logFile, err := os.CreateTemp("", "log_*.txt")
	if err != nil {
		t.Fatalf("Falha ao criar arquivo de log temporário: %v", err)
	}
	logPath := logFile.Name()
	defer os.Remove(logPath)
	defer logFile.Close()
	
	oldLogger := internal.Logger
	oldLogFile := internal.LogFile
	internal.LogFile = logFile
	internal.Logger = log.New(logFile, "", log.LstdFlags)
	defer func() {
		internal.Logger = oldLogger
		internal.LogFile = oldLogFile
	}()
	
	// Criar diretório temporário para teste
	tempDir, err := os.MkdirTemp("", "test_*")
	if err != nil {
		t.Fatalf("Falha ao criar diretório temporário: %v", err)
	}
	defer os.RemoveAll(tempDir)
	
	// Testar remoção de arquivo
	testFile := filepath.Join(tempDir, "arquivo_test.txt")
	if err := os.WriteFile(testFile, []byte("conteúdo de teste"), 0644); err != nil {
		t.Fatalf("Falha ao criar arquivo de teste: %v", err)
	}
	
	// Executar remoção
	result := removeWithLog(testFile)
	
	// Verificar resultado
	if !result {
		t.Error("removeWithLog retornou false, esperado true")
	}
	
	// Verificar se o arquivo foi removido
	if _, err := os.Stat(testFile); !os.IsNotExist(err) {
		t.Error("Arquivo não foi removido")
	}
	
	// Testar com caminho inexistente
	result = removeWithLog("/caminho/inexistente/teste123456")
	if !result {
		t.Error("removeWithLog retornou false para caminho inexistente, esperado true")
	}
}

// TestRunCleaner testa a função principal RunCleaner
// Este é um teste de integração simplificado
func TestRunCleaner(t *testing.T) {
	// Criar diretório temporário para teste
	tempBaseDir, err := os.MkdirTemp("", "test_cleaner_*")
	if err != nil {
		t.Fatalf("Falha ao criar diretório temporário: %v", err)
	}
	defer os.RemoveAll(tempBaseDir)
	
	// Criar alguns arquivos e diretórios com datas variadas
	// Arquivos recentes (não devem ser removidos)
	recentDir := filepath.Join(tempBaseDir, "recent")
	if err := os.Mkdir(recentDir, 0755); err != nil {
		t.Fatalf("Falha ao criar diretório recente: %v", err)
	}
	recentFile := filepath.Join(recentDir, "recent.txt")
	if err := os.WriteFile(recentFile, []byte("arquivo recente"), 0644); err != nil {
		t.Fatalf("Falha ao criar arquivo recente: %v", err)
	}
	
	// Arquivos antigos (devem ser removidos)
	oldDir := filepath.Join(tempBaseDir, "old")
	if err := os.Mkdir(oldDir, 0755); err != nil {
		t.Fatalf("Falha ao criar diretório antigo: %v", err)
	}
	oldFile := filepath.Join(oldDir, "old.txt")
	if err := os.WriteFile(oldFile, []byte("arquivo antigo"), 0644); err != nil {
		t.Fatalf("Falha ao criar arquivo antigo: %v", err)
	}
	
	// Definir tempo de modificação para arquivos antigos (31 dias atrás)
	oldTime := time.Now().AddDate(0, 0, -31)
	if err := os.Chtimes(oldFile, oldTime, oldTime); err != nil {
		t.Fatalf("Falha ao modificar data do arquivo: %v", err)
	}
	if err := os.Chtimes(oldDir, oldTime, oldTime); err != nil {
		t.Fatalf("Falha ao modificar data do diretório: %v", err)
	}
	
	// Criar configuração para teste
	// Usar SMTP falso para não tentar enviar email real
	testConfig := &config.Config{
		Directory:     tempBaseDir,
		DaysThreshold: 30,
		SMTP: struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			To       string `yaml:"to"`
		}{
			Host:     "localhost",
			Port:     1025, // porta não utilizada normalmente
			Username: "test",
			Password: "test",
			To:       "test@localhost",
		},
		Schedule: config.Schedule{
			Enabled: false,
			Cron:    "",
		},
	}
	
	// Configurar arquivo de log temporário
	logFile, err := os.CreateTemp("", "cleaner_log_*.txt")
	if err != nil {
		t.Fatalf("Falha ao criar arquivo de log temporário: %v", err)
	}
	logPath := logFile.Name()
	defer os.Remove(logPath)
	
	// Backup logger original
	oldLogger := internal.Logger
	oldLogFile := internal.LogFile
	// Configurar novo logger
	internal.LogFile = logFile
	internal.Logger = log.New(logFile, "", log.LstdFlags)
	// Restaurar logger original ao final
	defer func() {
		internal.Logger = oldLogger
		internal.LogFile = oldLogFile
	}()
	
	// Substituir função LoadConfig temporariamente para retornar nossa configuração de teste
	oldLoadConfig := config.LoadConfig
	config.LoadConfig = config.LoadConfigFunc(func(string) (*config.Config, error) {
		return testConfig, nil
	})
	defer func() {
		config.LoadConfig = oldLoadConfig
	}()
	
	// Executar limpador
	RunCleaner(testConfig)
	
	// Verificar se os arquivos antigos foram removidos
	if _, err := os.Stat(oldFile); !os.IsNotExist(err) {
		t.Error("Arquivo antigo não foi removido")
	}
	if _, err := os.Stat(oldDir); !os.IsNotExist(err) {
		t.Error("Diretório antigo não foi removido")
	}
	
	// Verificar se os arquivos recentes não foram removidos
	if _, err := os.Stat(recentFile); os.IsNotExist(err) {
		t.Error("Arquivo recente foi removido incorretamente")
	}
}
