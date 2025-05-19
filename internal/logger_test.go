package internal

import (
	"os"
	"testing"
)

func TestSetupLogger(t *testing.T) {
	// Definir nome de arquivo de log para teste
	logFileName := "test_log.txt"
	
	// Remover o arquivo de log de teste, caso exista
	_ = os.Remove(logFileName)
	
	// Substituir a variável LogFile
	oldLogFile := LogFile
	defer func() {
		// Restaurar LogFile original e fechar o arquivo de teste
		if LogFile != nil && LogFile != oldLogFile {
			LogFile.Close()
		}
		LogFile = oldLogFile
		
		// Limpar arquivo de log de teste
		_ = os.Remove(logFileName)
	}()
	
	// Executar a inicialização do logger com um arquivo específico
	SetupLogger()
	
	// Verificar se o logger foi inicializado corretamente
	if Logger == nil {
		t.Error("Logger não foi inicializado")
	}
	
	// Verificar se o arquivo de log foi criado
	if LogFile == nil {
		t.Error("Arquivo de log não foi criado")
	}
	
	// Testar gravação no log
	testMessage := "Mensagem de teste do logger"
	Logger.Println(testMessage)
	
	// Fechar o arquivo para garantir que o conteúdo seja gravado
	LogFile.Close()
	
	// Verificar se a mensagem foi gravada no arquivo
	content, err := os.ReadFile("log.txt")
	if err != nil {
		t.Errorf("Erro ao ler arquivo de log: %v", err)
	}
	
	if len(content) == 0 {
		t.Error("Arquivo de log está vazio")
	}
	
	// Limpeza
	_ = os.Remove("log.txt")
}
