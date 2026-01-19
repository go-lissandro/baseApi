## Estrutura de Pastas do Projeto

- **app/**  
  Contém arquivos especiais relacionados ao aplicativo, como rotas, inicializações, middlewares, etc.

- **configs/**  
  Contém arquivos de configuração, como banco de dados, variáveis de ambiente (`.env`), e outras configurações globais.

- **controllers/**  
  Contém os controladores responsáveis por gerenciar as requisições feitas ao aplicativo.

- **migrations/**  
  Contém arquivos com comandos de criação e alteração de tabelas que serão usados no banco de dados.

- **models/**  
  Contém arquivos que representam as estruturas das tabelas do banco de dados (models/entidades).

- **repository/**  
  Contém arquivos onde são executadas as operações de acesso ao banco de dados (CRUD, queries, etc).

- **services/**  
  Contém arquivos de serviço que fazem a ligação entre os controllers e os repositories, concentrando as regras de negócio.
