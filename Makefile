# Nome do projeto
APP_NAME=gocleaner

# Diretório de saída dos binários
DIST_DIR=dist

# Versão do build
VERSION=1.0.0

# Go build flags
LDFLAGS=-s -w -X main.version=$(VERSION)

# Plataformas de destino
PLATFORMS=darwin/amd64 linux/amd64 windows/amd64

# Default: compilar para todas
all: clean build

# Compilação cross-platform
build:
	@echo ">> Compilando $(APP_NAME) para todas as plataformas..."
	@mkdir -p $(DIST_DIR)
	@for PLATFORM in $(PLATFORMS); do \
		OS=$$(echo $$PLATFORM | cut -d/ -f1); \
		ARCH=$$(echo $$PLATFORM | cut -d/ -f2); \
		EXT=$$(if [ $$OS = "windows" ]; then echo ".exe"; else echo ""; fi); \
		OUTPUT=$(DIST_DIR)/$(APP_NAME)-$$OS-$$ARCH$$EXT; \
		echo "Compilando: $$OS/$$ARCH -> $$OUTPUT"; \
		GOOS=$$OS GOARCH=$$ARCH go build -ldflags="$(LDFLAGS)" -o $$OUTPUT . ; \
	done

# Remover binários
clean:
	@echo ">> Limpando diretório $(DIST_DIR)..."
	@rm -rf $(DIST_DIR)

# Testes (opcional)
test:
	@echo ">> Executando testes..."
	@go test ./...

.PHONY: all build clean test
