# API Auth Go

API de autenticação desenvolvida em Go com Gin, GORM e PostgreSQL.

## 🚀 Executando Localmente

### Pré-requisitos
- Go 1.24+
- PostgreSQL
- Air (para hot-reload)

### Instalação

1. Clone o repositório:
```bash
git clone <repository-url>
cd api-auth-go
```

2. Instale as dependências:
```bash
go mod download
```

3. Instale o Air para hot-reload:
```bash
go install github.com/air-verse/air@latest
```

4. Configure o PostgreSQL:
```bash
# Inicie o PostgreSQL
sudo systemctl start postgresql

# Crie o banco de dados (se não existir)
sudo -u postgres psql -c "CREATE DATABASE auth_api_dev;"
```

### Executando a aplicação

#### Opção 1: Com Air (Recomendado para desenvolvimento)
```bash
air
```

#### Opção 2: Execução direta
```bash
go run ./cmd/api
```

A aplicação estará disponível em `http://localhost:8080`

### Comandos úteis

- **Verificar status do PostgreSQL:**
```bash
sudo systemctl status postgresql@14-main
```

- **Iniciar PostgreSQL:**
```bash
sudo systemctl start postgresql
```

- **Parar PostgreSQL:**
```bash
sudo systemctl stop postgresql
```

- **Testar conexão com o banco:**
```bash
psql -h localhost -U postgres -d auth_api_dev
```

## 🔧 Desenvolvimento

### Hot Reload com Air

A aplicação está configurada com Air para hot reload. Qualquer alteração nos arquivos `.go` irá automaticamente recompilar e reiniciar a aplicação.

O arquivo `.air.toml` já está configurado para monitorar as mudanças.

### Variáveis de Ambiente

As variáveis de ambiente são carregadas automaticamente com valores padrão. Para personalizar, você pode definir as seguintes variáveis:

```bash
# Application
export PORT=8080
export JWT_SECRET=your-secret-key-change-in-production

# Database
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=auth_api_dev
export DB_SSLMODE=disable
```

**Valores padrão:**
- `PORT`: 8080
- `DB_HOST`: localhost
- `DB_PORT`: 5432
- `DB_USER`: postgres
- `DB_PASSWORD`: postgres
- `DB_NAME`: auth_api_dev
- `DB_SSLMODE`: disable
- `JWT_SECRET`: your-secret-key

### Estrutura do Projeto

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
│   │   └── server/
│   └── presentation/
│       ├── handlers/
│       └── routes/
├── .air.toml
└── go.mod
```

## 📝 Endpoints

### Health Check
- `GET /health` - Verificar status da aplicação

### Usuários
- `POST /users` - Criar usuário
- `GET /users/:id` - Buscar usuário por ID
- `PUT /users/:id` - Atualizar usuário
- `DELETE /users/:id` - Deletar usuário

## 🛠️ Tecnologias

- **Go 1.24**
- **Gin** - Framework web
- **GORM** - ORM para Go
- **PostgreSQL** - Banco de dados
- **Air** - Hot reload para desenvolvimento

## 🐛 Troubleshooting

### Problemas comuns

1. **Erro de conexão com PostgreSQL:**
   - Verifique se o PostgreSQL está rodando: `sudo systemctl status postgresql@14-main`
   - Inicie o serviço: `sudo systemctl start postgresql`

2. **Air não encontrado:**
   - Instale o Air: `go install github.com/air-verse/air@latest`
   - O Air será automaticamente adicionado ao PATH

3. **Erro de compilação:**
   - Verifique se todas as dependências estão instaladas: `go mod download`
   - Limpe o cache: `go clean -cache`

4. **Porta já em uso:**
   - Verifique se não há outro processo na porta 8080: `lsof -i :8080`
   - Mude a porta nas variáveis de ambiente: `export PORT=8081` 