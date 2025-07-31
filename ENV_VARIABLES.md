# Variáveis de Ambiente

Este documento descreve as variáveis de ambiente utilizadas pela API Auth Go.

## 📋 Variáveis Disponíveis

### Server Configuration
| Variável | Padrão | Descrição |
|----------|--------|-----------|
| `PORT` | `8080` | Porta onde a API será executada |

### Database Configuration
| Variável | Padrão | Descrição |
|----------|--------|-----------|
| `DB_HOST` | `localhost` | Host do banco de dados |
| `DB_PORT` | `5432` | Porta do banco de dados |
| `DB_USER` | `postgres` | Usuário do banco de dados |
| `DB_PASSWORD` | `postgres` | Senha do banco de dados |
| `DB_NAME` | `auth_api_dev` | Nome do banco de dados |
| `DB_SSLMODE` | `disable` | Modo SSL do banco de dados |

### JWT Configuration
| Variável | Padrão | Descrição |
|----------|--------|-----------|
| `JWT_SECRET` | `your-secret-key` | Chave secreta para assinatura dos tokens JWT |

### Email Configuration
| Variável | Padrão | Descrição |
|----------|--------|-----------|
| `EMAIL_FROM` | - | Email de origem para envio de emails |
| `EMAIL_PASSWORD` | - | Senha do email de origem |
| `SMTP_HOST` | - | Host do servidor SMTP |
| `SMTP_PORT` | - | Porta do servidor SMTP |

### SMS Configuration
**Nota:** O envio de SMS foi temporariamente desabilitado. A funcionalidade está focada apenas no envio de email.

## 🔧 Configuração

### Para Desenvolvimento Local
1. Copie o arquivo `.env.example` para `.env`
2. Configure as variáveis conforme necessário
3. Execute `make setup` para verificar a configuração

### Para Docker
- As variáveis são carregadas automaticamente do arquivo `.env`
- O `DB_HOST` é automaticamente definido como `postgres` no ambiente Docker

## ⚠️ Segurança

**IMPORTANTE**: Em produção, sempre altere:
- `JWT_SECRET` para uma chave forte e única
- `DB_PASSWORD` para uma senha segura
- `DB_USER` para um usuário específico da aplicação

## 📝 Exemplo de Arquivo .env

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
EMAIL_FROM=your-email@gmail.com
EMAIL_PASSWORD=your-app-password
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587

# SMS Configuration
# Nota: SMS temporariamente desabilitado
``` 