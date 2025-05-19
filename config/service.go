// Package config contém funções e estruturas relacionadas à configuração do aplicativo
package config

// Este arquivo serve como um serviço de configuração para o aplicativo
// Contém funções auxiliares para manipulação de configurações

// GetConfigPath retorna o caminho padrão do arquivo de configuração
func GetConfigPath() string {
	return "config.yaml"
}