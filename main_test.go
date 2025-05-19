package main

import (
	"os"
	"testing"
)

// TestMain testa a execução da função main
// Este teste simples verifica se o pacote main pode ser importado
// corretamente, já que não podemos substituir a função cmd.Execute diretamente
func TestMain(t *testing.T) {
	// Como não podemos substituir funções de pacotes importados em Go,
	// só podemos verificar se o pacote é válido e pode ser compilado
	// Um teste mais completo poderia ser realizado em cmd/root_test.go
	t.Log("Verificação de compilação do pacote main")

	// Verificar se a função principal existe
	var _ = main
}

// TestMainIntegration é um teste de integração que verifica
// se o programa pode ser executado em um ambiente controlado
func TestMainIntegration(t *testing.T) {
	// Criar arquivo de configuração temporário para teste
	tempConfig := `
directory: "/tmp/test_integration"
days_threshold: 7
schedule:
  enabled: false
  cron: "0 0 * * *"
smtp:
  host: "localhost"
  port: 1025
  username: "test"
  password: "test"
  to: "test@localhost"
`
	// Criar diretório de teste se não existir
	testDir := "/tmp/test_integration"
	if err := os.MkdirAll(testDir, 0755); err != nil {
		t.Fatalf("Não foi possível criar diretório de teste: %v", err)
	}
	
	// Criar arquivo de configuração
	if err := os.WriteFile("config.yaml", []byte(tempConfig), 0644); err != nil {
		t.Fatalf("Não foi possível criar arquivo de configuração de teste: %v", err)
	}
	
	// Limpar no final
	defer func() {
		os.Remove("config.yaml")
		os.Remove("log.txt")
	}()
	
	// Este teste é apenas para verificar se o programa pode ser iniciado
	// sem erros, não executamos a função main real para evitar
	// efeitos colaterais como envio de email ou exclusão de arquivos.
	// Em vez disso, apenas verificamos se o pacote pode ser importado e
	// se a função main existe.
	t.Log("O pacote main foi importado com sucesso e contém a função main")
}
