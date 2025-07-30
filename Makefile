# Makefile para API Auth Go
# Uso: make [comando]

# Variáveis
DOCKER_COMPOSE = docker-compose
GO = go

# Cores para output
GREEN = \033[0;32m
YELLOW = \033[1;33m
RED = \033[0;31m
BLUE = \033[0;34m
NC = \033[0m # No Color

# Função para imprimir mensagens coloridas
define print_info
	@echo "$(GREEN)[INFO]$(NC) $1"
endef

define print_warning
	@echo "$(YELLOW)[WARNING]$(NC) $1"
endef

define print_header
	@echo "$(BLUE)================================$(NC)"
	@echo "$(BLUE)  API Auth Go - Makefile$(NC)"
	@echo "$(BLUE)================================$(NC)"
endef

# Comandos principais
.PHONY: help up up-d down logs shell clean check-env setup

# Comando padrão
help: ## Mostrar esta ajuda
	$(call print_header)
	@echo ""
	@echo "Comandos disponíveis:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-15s$(NC) %s\n", $$1, $$2}'
	@echo ""

# Docker commands
up: ## Iniciar ambiente de desenvolvimento
	$(call print_info,"Iniciando ambiente de desenvolvimento...")
	$(DOCKER_COMPOSE) up --build

up-d: ## Iniciar ambiente de desenvolvimento em background
	$(call print_info,"Iniciando ambiente de desenvolvimento em background...")
	$(DOCKER_COMPOSE) up -d --build

down: ## Parar ambiente de desenvolvimento
	$(call print_info,"Parando ambiente de desenvolvimento...")
	$(DOCKER_COMPOSE) down

logs: ## Mostrar logs da API
	$(call print_info,"Mostrando logs da API...")
	$(DOCKER_COMPOSE) logs -f api

shell: ## Acessar shell do container da API
	$(call print_info,"Acessando shell do container da API...")
	$(DOCKER_COMPOSE) exec api sh

# Clean commands
clean: ## Parar e remover containers e volumes
	$(call print_warning,"Isso irá remover todos os containers e volumes. Tem certeza? (y/N)")
	@read -p "Confirma? [y/N]: " confirm && [ "$$confirm" = "y" ] || [ "$$confirm" = "Y" ] || (echo "Operação cancelada." && exit 1)
	$(call print_info,"Limpando ambiente...")
	$(DOCKER_COMPOSE) down -v
	docker system prune -f

# Utility commands
check-env: ## Verificar se arquivo .env existe
	$(call print_info,"Verificando arquivo .env...")
	@if [ ! -f .env ]; then \
		echo "$(RED)[ERROR]$(NC) Arquivo .env não encontrado!"; \
		echo "Copie o arquivo .env.example para .env e configure as variáveis."; \
		exit 1; \
	fi
	$(call print_info,"Arquivo .env encontrado!")

setup: ## Configurar ambiente de desenvolvimento
	$(call print_info,"Configurando ambiente de desenvolvimento...")
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo "$(YELLOW)[WARNING]$(NC) Arquivo .env criado a partir do .env.example"; \
		echo "Configure as variáveis no arquivo .env antes de continuar."; \
	fi
	$(GO) mod download
	$(call print_info,"Ambiente configurado!")

# Default target
.DEFAULT_GOAL := help 