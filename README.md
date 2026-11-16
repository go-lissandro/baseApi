** app/ -- teremos arquivos especiais relacionados ao aplicativo na rota, etc.

** configs/ -- ele conterá arquivos como banco de dados e .env que usamos para configurações.

** controllers/ -- teremos um controlador que gerencia as solicitações ao nosso aplicativo.

** migrations/ -- ele conterá arquivos contendo comandos de criação de tabelas que usaremos em nosso.banco de dados

** models/ -- conterá arquivos pertencentes às nossas estruturas de tabela.

** repository/ -- haverá arquivos nos quais executaremos nossas operações de banco de dados.

** services/ -- teremos arquivos de serviço que fornecem a conexão entre o controlador e o repositório.

## Como rodar

```bash
cp .env.example .env   # preencha usuário, senha e JWT_SECRET
docker compose up --build
```

A API sobe em `http://localhost:3000/api/v1`.

| Método | Rota | Acesso |
|--------|------|--------|
| POST | `/create-user` | público (cria usuário com papel `user`) |
| POST | `/login` | público (retorna JWT) |
| POST | `/custommers` | público |
| GET | `/admin` | admin |
| GET | `/admin/custommers` | admin |

Para promover um usuário a admin:

```sql
UPDATE users SET role = 'admin' WHERE email = 'voce@exemplo.com';
```

## Testes

```bash
go test ./...
```
