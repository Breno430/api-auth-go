# API Auth Go

Uma API de autenticação desenvolvida em Go com Gin, GORM e PostgreSQL, incluindo sistema RBAC (Role Based Access Control).

## 🚀 Tecnologias

- **Go** - Linguagem de programação
- **Gin** - Framework web
- **GORM** - ORM para Go
- **PostgreSQL** - Banco de dados
- **JWT** - Autenticação
- **Docker** - Containerização
- **RBAC** - Controle de acesso baseado em roles

## 📋 Pré-requisitos

- Docker
- Docker Compose

## 🐳 Executando com Docker

### Ambiente de Desenvolvimento (com Hot Reload)

```bash
# Usando Makefile (recomendado)
make up

# Ou em background
make up-d

# Ou diretamente com docker-compose
docker-compose up --build
docker-compose up -d --build
```

## 🛠️ Comandos Úteis

### Com Makefile (Recomendado)
```bash
# Ver todos os comandos disponíveis
make help

# Iniciar ambiente
make up

# Parar ambiente
make down

# Ver logs da API
make logs

# Ver logs do banco
make logs-db

# Acessar shell da API
make shell

# Acessar shell do banco
make shell-db

# Verificar status
make status

# Limpar ambiente
make clean

# Seed manual (se necessário)
make seed-admin
```

### Com Docker Compose Diretamente
```bash
# Parar os containers
docker-compose down

# Parar e remover volumes
docker-compose down -v

# Ver logs
docker-compose logs -f api

# Acessar container da API
docker-compose exec api sh

# Acessar container do PostgreSQL
docker-compose exec postgres psql -U postgres -d auth_api_dev
```

## 🔧 Variáveis de Ambiente

As variáveis de ambiente são carregadas do arquivo `.env`. O Docker Compose usa as seguintes variáveis:

| Variável | Descrição |
|----------|-----------|
| `PORT` | Porta da API |
| `DB_HOST` | Host do banco de dados (usado como `postgres` no container) |
| `DB_PORT` | Porta do banco de dados |
| `DB_USER` | Usuário do banco |
| `DB_PASSWORD` | Senha do banco |
| `DB_NAME` | Nome do banco |
| `DB_SSLMODE` | Modo SSL do banco |
| `JWT_SECRET` | Chave secreta do JWT |
| `EMAIL_FROM` | Email remetente para envio |
| `EMAIL_PASSWORD` | Senha de app do email |
| `SMTP_HOST` | Servidor SMTP |
| `SMTP_PORT` | Porta do servidor SMTP |

**Nota**: No ambiente Docker, o `DB_HOST` é automaticamente definido como `postgres` (nome do container).

### Configuração Inicial

```bash
# Verificar se o arquivo .env existe
make check-env

# Configurar ambiente (cria .env se não existir)
make setup

# Iniciar projeto (seed automática)
make up
```

**Nota**: A seed é executada automaticamente na primeira inicialização, criando o usuário admin padrão.

### Exemplo de Arquivo .env

```env
# Server Configuration
PORT=8080

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=auth_api_dev
DB_SSLMODE=disable

# JWT Configuration
JWT_SECRET=your-secret-key-change-in-production

# Email Configuration
EMAIL_FROM=seu-email@gmail.com
EMAIL_PASSWORD=sua-senha-de-app-gerada
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
```

**⚠️ Segurança**: Em produção, sempre altere as senhas padrão e chaves secretas!

## 🔐 Sistema RBAC (Role Based Access Control)

A API implementa um sistema completo de controle de acesso baseado em roles com dois perfis:

### 👥 Perfis de Usuário

#### **Admin**
- ✅ Acesso completo ao sistema
- ✅ Pode criar novos usuários
- ✅ Pode listar todos os usuários com filtros
- ✅ Pode visualizar, atualizar e deletar qualquer usuário
- ✅ Pode acessar todas as rotas

#### **User**
- ✅ Acesso limitado aos próprios dados
- ✅ Pode visualizar e atualizar apenas seus próprios dados
- ✅ Pode listar apenas seus próprios dados
- ❌ **NÃO pode se deletar**
- ❌ **NÃO pode criar usuários**
- ❌ Não pode acessar dados de outros usuários

### 🌱 Seed Automática

A API executa automaticamente uma seed na inicialização que cria o usuário admin padrão:

- **Email**: admin@example.com
- **Password**: admin123
- **Role**: admin

A seed só executa se o usuário admin ainda não existir, garantindo que não seja criado duplicado.

## 📊 Endpoints

A API estará disponível em `http://localhost:8080`

### 🔓 Rotas Públicas
```
POST /api/v1/users/login      # Login
POST /api/v1/password-reset/request  # Solicitar reset de senha
POST /api/v1/password-reset/reset    # Resetar senha
```

### 🔒 Rotas Protegidas (Todos os usuários autenticados)
```
GET /api/v1/profile           # Ver perfil próprio
GET /api/v1/users            # Listar usuários (admin: todos, user: apenas próprio)
GET /api/v1/users/:id        # Ver usuário específico (admin: qualquer, user: apenas próprio)
PUT /api/v1/users/:id        # Atualizar usuário (admin: qualquer, user: apenas próprio)
DELETE /api/v1/users/:id     # Deletar usuário (apenas admin)
```

### 👑 Rotas de Administração (Apenas Admin)
```
POST /api/v1/admin/users     # Criar usuário (apenas admin)
```

## 🔍 Filtros de Listagem

### Query Parameters Disponíveis

| Parâmetro | Tipo | Descrição | Exemplo |
|-----------|------|-----------|---------|
| `name` | string | Filtrar por nome (busca parcial) | `?name=joão` |
| `email` | string | Filtrar por email (busca parcial) | `?email=gmail` |
| `role` | string | Filtrar por role | `?role=admin` |
| `page` | int | Número da página | `?page=2` |
| `limit` | int | Itens por página (max: 100) | `?limit=20` |
| `sort_by` | string | Campo para ordenar | `?sort_by=name` |
| `sort_order` | string | Ordem (asc/desc) | `?sort_order=asc` |

### Campos de Ordenação Válidos
- `name` - Nome do usuário
- `email` - Email do usuário
- `role` - Role do usuário
- `created_at` - Data de criação
- `updated_at` - Data de atualização

## 📝 Exemplos de Uso

### 1. Iniciar o Projeto (Seed Automática)
```bash
# A seed é executada automaticamente na inicialização
make up

# Ou em background
make up-d
```

**Nota**: O usuário admin será criado automaticamente na primeira execução:
- **Email**: admin@example.com
- **Password**: admin123
- **Role**: admin

### 2. Login como Admin
```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "admin123"
  }'
```

### 3. Criar Usuário (Apenas Admin)
```bash
curl -X POST http://localhost:8080/api/v1/admin/users \
  -H "Authorization: Bearer <token_do_admin>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Novo Usuário",
    "email": "novo@email.com",
    "password": "senha123"
  }'
```

### 4. Listar Usuários com Filtros (Admin)
```bash
# Listar todos
curl -X GET "http://localhost:8080/api/v1/users" \
  -H "Authorization: Bearer <token_do_admin>"

# Filtrar por nome
curl -X GET "http://localhost:8080/api/v1/users?name=joão" \
  -H "Authorization: Bearer <token_do_admin>"

# Paginação
curl -X GET "http://localhost:8080/api/v1/users?page=1&limit=5" \
  -H "Authorization: Bearer <token_do_admin>"

# Ordenação
curl -X GET "http://localhost:8080/api/v1/users?sort_by=name&sort_order=asc" \
  -H "Authorization: Bearer <token_do_admin>"
```

### 5. Atualizar Usuário
```bash
# Admin pode atualizar qualquer usuário
curl -X PUT http://localhost:8080/api/v1/users/<user_id> \
  -H "Authorization: Bearer <token_do_admin>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Novo Nome",
    "email": "novo@email.com",
    "role": "user"
  }'

# User só pode atualizar seus próprios dados
curl -X PUT http://localhost:8080/api/v1/users/<seu_user_id> \
  -H "Authorization: Bearer <token_do_user>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Meu Novo Nome",
    "email": "meu@email.com",
    "role": "user"
  }'
```

### 6. Deletar Usuário (Apenas Admin)
```bash
curl -X DELETE http://localhost:8080/api/v1/users/<user_id> \
  -H "Authorization: Bearer <token_do_admin>"
```

## 🗄️ Banco de Dados

O PostgreSQL será executado com as credenciais definidas no arquivo `.env`:
- **Host**: localhost
- **Porta**: Definida em `DB_PORT` (padrão: 5432)
- **Database**: Definido em `DB_NAME` (padrão: auth_api_dev)
- **Usuário**: Definido em `DB_USER` (padrão: postgres)
- **Senha**: Definida em `DB_PASSWORD` (padrão: postgres)

## 🔄 Hot Reload (Desenvolvimento)

No ambiente de desenvolvimento, a API usa o [Air](https://github.com/cosmtrek/air) para hot reload automático. Qualquer alteração no código será automaticamente recompilada e reiniciada.

## 🛠️ Makefile

O projeto inclui um Makefile completo com comandos úteis para desenvolvimento:

### Comandos Disponíveis
```bash
# Ver todos os comandos disponíveis
make help

# Iniciar ambiente de desenvolvimento
make up

# Iniciar em background
make up-d

# Parar ambiente
make down

# Ver logs da API
make logs

# Acessar shell da API
make shell

# Limpar ambiente
make clean

# Verificar arquivo .env
make check-env

# Configurar ambiente
make setup

# Seed manual (se necessário)
make seed-admin
```

## 📁 Estrutura do Projeto

```
api-auth-go/
├── cmd/
│   ├── api/
│   │   └── main.go
├── internal/
│   ├── domain/
│   │   ├── entities/
│   │   ├── repositories/
│   │   └── usecases/
│   ├── infrastructure/
│   │   ├── config/
│   │   ├── database/
│   │   ├── repositories/
│   │   ├── server/
│   │   └── services/
│   └── presentation/
│       ├── handlers/
│       ├── middleware/
│       └── routes/
├── Dockerfile
├── docker-compose.yml
├── Makefile
├── .air.toml
└── .dockerignore
```

## 📧 Configuração do Serviço de Email

Para configurar o envio de emails, você precisa obter os valores corretos do seu provedor de email:

### 🔧 Como Obter os Valores para Gmail

#### **1. Ativar Autenticação de 2 Fatores**
1. Acesse: https://myaccount.google.com/security
2. Ative "Verificação em duas etapas"

#### **2. Gerar Senha de App**
1. Acesse: https://myaccount.google.com/apppasswords
2. Selecione "Email" e "Outro (nome personalizado)"
3. Digite "API Auth Go" como nome
4. Clique em "Gerar"
5. **Copie a senha gerada (16 caracteres)**

#### **3. Configurar no .env**
```env
EMAIL_FROM=seu-email@gmail.com
EMAIL_PASSWORD=sua-senha-de-app-gerada
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
```

### 📧 Outros Provedores

#### **Outlook/Hotmail**
```env
EMAIL_FROM=seu-email@outlook.com
EMAIL_PASSWORD=sua-senha-de-app
SMTP_HOST=smtp-mail.outlook.com
SMTP_PORT=587
```

#### **Yahoo**
```env
EMAIL_FROM=seu-email@yahoo.com
EMAIL_PASSWORD=sua-senha-de-app
SMTP_HOST=smtp.mail.yahoo.com
SMTP_PORT=587
```

## 🚨 Segurança

⚠️ **Importante**: Em produção, sempre altere as senhas padrão e chaves secretas configuradas no Docker Compose.

### 🔐 Validações de Segurança

- ✅ **UUID Validation**: Todos os endpoints que recebem ID validam UUID
- ✅ **Input Validation**: Dados de entrada são validados na entidade
- ✅ **Filter Validation**: Filtros são validados antes da consulta
- ✅ **Role-based Access**: Controle de acesso baseado no role
- ✅ **SQL Injection Protection**: Filtros são aplicados com prepared statements
- ✅ **User Self-Delete Prevention**: Usuários não podem se deletar
- ✅ **Admin-Only User Creation**: Apenas admins podem criar usuários

## 🔄 Hot Reload

O ambiente usa o [Air](https://github.com/cosmtrek/air) para hot reload automático. Qualquer alteração no código será automaticamente recompilada e reiniciada.

## 📝 Migrations

As migrations do GORM serão executadas automaticamente quando a aplicação iniciar. Certifique-se de que suas migrations estão configuradas corretamente no código.

## 📚 Documentação Adicional

- [ENV_VARIABLES.md](ENV_VARIABLES.md) - Documentação completa das variáveis de ambiente
- [Makefile](Makefile) - Comandos disponíveis para desenvolvimento 