# API Auth Go

Uma API de autenticação desenvolvida em Go com Gin, GORM e PostgreSQL.

## 🚀 Tecnologias

- **Go** - Linguagem de programação
- **Gin** - Framework web
- **GORM** - ORM para Go
- **PostgreSQL** - Banco de dados
- **JWT** - Autenticação
- **Docker** - Containerização

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

**Nota**: No ambiente Docker, o `DB_HOST` é automaticamente definido como `postgres` (nome do container).

### Configuração Inicial

```bash
# Verificar se o arquivo .env existe
make check-env

# Configurar ambiente (cria .env se não existir)
make setup
```

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
```

**⚠️ Segurança**: Em produção, sempre altere as senhas padrão e chaves secretas!

## 📊 Endpoints

A API estará disponível em `http://localhost:8080`

### Endpoints de Saúde
- `GET /health` - Verificar status da API

### Endpoints de Usuário
- `POST /api/v1/users/signup` - Registrar novo usuário
- `POST /api/v1/users/login` - Fazer login
- `GET /api/v1/profile` - Obter perfil do usuário (requer autenticação)

### Endpoints de Recuperação de Senha
- `POST /api/v1/password-reset/request` - Solicitar recuperação de senha
- `POST /api/v1/password-reset/reset` - Redefinir senha com token

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
```

## 📁 Estrutura do Projeto

```
api-auth-go/
├── cmd/
│   └── api/
│       └── main.go
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

Para que a funcionalidade de recuperação de senha funcione corretamente, você precisa configurar o serviço de email. Aqui estão as instruções para diferentes provedores:

### 🔧 Configuração para Gmail (Recomendado)

#### **1. Ativar Autenticação de 2 Fatores**
1. Acesse: https://myaccount.google.com/security
2. Ative "Verificação em duas etapas"

#### **2. Gerar Senha de App**
1. Acesse: https://myaccount.google.com/apppasswords
2. Selecione "Email" e "Outro (nome personalizado)"
3. Digite "API Auth Go" como nome
4. Clique em "Gerar"
5. **Copie a senha gerada (16 caracteres)**

#### **3. Configurar Variáveis de Ambiente**
Adicione ao seu arquivo `.env`:

```env
# Email Configuration (Gmail)
EMAIL_FROM=seu-email@gmail.com
EMAIL_PASSWORD=sua-senha-de-app-gerada
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
```

**Exemplo:**
```env
EMAIL_FROM=joao.silva@gmail.com
EMAIL_PASSWORD=abcd efgh ijkl mnop
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
```

### 📧 Configuração para Outlook/Hotmail

```env
EMAIL_FROM=seu-email@outlook.com
EMAIL_PASSWORD=sua-senha-de-app
SMTP_HOST=smtp-mail.outlook.com
SMTP_PORT=587
```

### 📧 Configuração para Yahoo

```env
EMAIL_FROM=seu-email@yahoo.com
EMAIL_PASSWORD=sua-senha-de-app
SMTP_HOST=smtp.mail.yahoo.com
SMTP_PORT=587
```

### 📧 Configuração para ProtonMail

```env
EMAIL_FROM=seu-email@protonmail.com
EMAIL_PASSWORD=sua-senha-de-app
SMTP_HOST=127.0.0.1
SMTP_PORT=1025
```

### 🧪 Testando o Envio de Email

#### **1. Iniciar a Aplicação**
```bash
make up
```

#### **2. Cadastrar um Usuário**
```bash
curl -X POST http://localhost:8080/api/v1/users/signup \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Teste Usuario",
    "email": "seu-email@gmail.com",
    "password": "123456"
  }'
```

#### **3. Solicitar Recuperação de Senha**
```bash
curl -X POST http://localhost:8080/api/v1/password-reset/request \
  -H "Content-Type: application/json" \
  -d '{
    "email": "seu-email@gmail.com"
  }'
```

#### **4. Verificar os Logs**
```bash
# Ver logs da aplicação
make logs

# Logs esperados:
# 2024/01/15 10:30:00 Email de recuperação enviado para: seu-email@gmail.com
```

#### **5. Verificar o Email**
1. Abra seu email (Gmail)
2. Procure por um email com assunto "Recuperação de Senha"
3. O email conterá um PIN de 6 dígitos

### 🚨 Troubleshooting

#### **Erro: "authentication failed"**
- Verifique se a senha de app está correta
- Certifique-se de que a autenticação de 2 fatores está ativada
- Gere uma nova senha de app se necessário

#### **Erro: "connection refused"**
- Verifique se o SMTP_HOST e SMTP_PORT estão corretos
- Teste a conectividade: `telnet smtp.gmail.com 587`

#### **Erro: "invalid credentials"**
- Use a senha de app, não a senha normal da conta
- Verifique se o email está correto

### 🔧 Configuração Alternativa (Sem Email Real)

Se você não quiser configurar um email real, pode modificar o `EmailService` para apenas logar:

```go
// Em internal/infrastructure/services/email_service.go
func (es *EmailService) SendPasswordResetEmail(to, name, token string) error {
    log.Printf("SIMULAÇÃO: Email enviado para %s com token %s", to, token)
    return nil
}
```

### 📋 Checklist de Configuração

- [ ] Ativar autenticação de 2 fatores no Gmail
- [ ] Gerar senha de app
- [ ] Configurar arquivo `.env` com as variáveis
- [ ] Reiniciar a aplicação
- [ ] Testar cadastro de usuário
- [ ] Testar solicitação de recuperação
- [ ] Verificar logs da aplicação
- [ ] Verificar recebimento do email

## 🚨 Segurança

⚠️ **Importante**: Em produção, sempre altere as senhas padrão e chaves secretas configuradas no Docker Compose.

## 🔄 Hot Reload

O ambiente usa o [Air](https://github.com/cosmtrek/air) para hot reload automático. Qualquer alteração no código será automaticamente recompilada e reiniciada.

## 📝 Migrations

As migrations do GORM serão executadas automaticamente quando a aplicação iniciar. Certifique-se de que suas migrations estão configuradas corretamente no código.

## 📚 Documentação Adicional

- [ENV_VARIABLES.md](ENV_VARIABLES.md) - Documentação completa das variáveis de ambiente
- [Makefile](Makefile) - Comandos disponíveis para desenvolvimento 