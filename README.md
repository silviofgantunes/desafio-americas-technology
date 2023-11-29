# PT-BR (English below) enunciado original - desafio-americas-technology

O objetivo deste teste é avaliar suas habilidades de codificação e arquitetura em um ambiente de desenvolvimento. Você receberá um problema simples para demonstrar suas técnicas de programação. Recomenda-se exagerar um pouco na solução para destacar suas capacidades. A aplicação que você desenvolverá deve ser pronta para produção, com a consideração de manutenção por outros desenvolvedores ao longo do tempo.
Você pode e deve utilizar bibliotecas de terceiros, além de decidir sobre o uso ou não de um framework. Durante a entrevista de revisão de código, esteja preparado para responder perguntas sobre essas escolhas, explicando como e por que foram feitas, e quais alternativas foram consideradas.
Para facilitar o teste, utilizamos o Docker na Americas Technology. Certifique-se de usar Docker para garantir resultados consistentes. Alguns containers estão disponíveis para auxiliar na construção e execução das aplicações, mas você pode modificá-los conforme preferir.
Instruções resumidas:
Clone este repositório.
Crie uma nova branch chamada dev.
Desenvolva duas aplicações básicas (microserviços) que se comuniquem entre si: um cadastro de usuários e um serviço de ordens de compra e venda de crypto.
Crie uma "pull request" da branch dev para a "branch" master, incluindo instruções de execução, tecnologias utilizadas e justificativas para as escolhas feitas.
Requisitos das aplicações:
Cadastro de usuários: Listar, exibir, criar, alterar e excluir usuários.
Tabela de usuários (user): id, name, email, phone_number, created_at, updated_at.
Serviço de ordens: Listar, listar por usuário, exibir, criar, excluir ordem (limit).
Tabela de ordens (compra e venda de crypto) (order): id, user_id, pair, amount, direction, type (market, limit), created_at, updated_at.
Comunicação entre os serviços para garantir a consistência de dados.
Padrões REST e documentação contendo endpoints e payloads utilizados.
Critérios de avaliação destacados:
Uso de bibliotecas de terceiros e possivelmente de um framework, com justificação.
Execução em containers Docker.
Retorno de JSON válido e presença dos recursos citados.
Código testável com testes unitários.
Adesão a melhores práticas de segurança de APIs e diretrizes de estilo de código.
Pontos considerados bônus:
Respostas durante o code review e descrição detalhada na "pull request".
Setup da aplicação em apenas um comando ou script facilitador.
Outros tipos de testes, como funcionais e de integração.
Histórico de commits com mensagens descritivas.
Boa sorte!

# PT-BR (English below) enunciado formatado e reescrito por IA - desafio-americas-technology

Projeto de Desenvolvimento de Microserviços para Cadastro de Usuários e para Ordens de Compra e Venda de Criptoativos

O propósito deste teste é avaliar suas habilidades em codificação e arquitetura em um ambiente de desenvolvimento. Você receberá um problema simples para demonstrar suas técnicas de programação, e é recomendado que você exagere um pouco na solução para destacar suas capacidades. A aplicação a ser desenvolvida deve estar pronta para produção, levando em consideração a manutenção por outros desenvolvedores ao longo do tempo.

Requisitos das Aplicações:

Cadastro de Usuários:

Listar, exibir, criar, alterar e excluir usuários.  Tabela de usuários (user): id, name, email, phone_number, created_at, updated_at.

Serviço de Ordens:

Listar, listar por usuário, exibir, criar, excluir ordem (limit).  Tabela de ordens (compra e venda de crypto) (order): id, user_id, pair, amount, direction, type (market, limit), created_at, updated_at.  Comunicação entre os serviços para garantir a consistência de dados.  Padrões REST e documentação contendo endpoints e payloads utilizados.

Critérios de Avaliação Destacados:

Uso de bibliotecas de terceiros e, possivelmente, de um framework, com justificação.  Execução em containers Docker.  Retorno de JSON válido e presença dos recursos citados.  Código testável com testes unitários.  Adesão a melhores práticas de segurança de APIs e diretrizes de estilo de código.

Pontos Considerados Bônus:

Respostas durante o code review e descrição detalhada na "pull request".  Setup da aplicação em apenas um comando ou script facilitador.  Outros tipos de testes, como funcionais e de integração.  Histórico de commits com mensagens descritivas.

# English - Americas Technology Challenge

The purpose of this test is to assess your coding and architecture skills in a development environment. You will be given a simple problem to demonstrate your programming techniques, and it is recommended that you go a bit beyond in the solution to highlight your capabilities. The application to be developed should be production-ready, taking into account maintenance by other developers over time.

Application Requirements:

User Registration:

List, display, create, update, and delete users. Users table (user): id, name, email, phone_number, created_at, updated_at.

Orders Service:

List, list by user, display, create, delete order (limit). Orders table (buying and selling crypto) (order): id, user_id, pair, amount, direction, type (market, limit), created_at, updated_at. Communication between services to ensure data consistency. REST standards and documentation containing used endpoints and payloads.

Highlighted Evaluation Criteria:

Use of third-party libraries and possibly a framework, with justification. Execution in Docker containers. Valid JSON return and presence of mentioned features. Testable code with unit tests. Adherence to API security best practices and code style guidelines.

Bonus Points Considered:

Responses during code review and detailed description in the "pull request." Application setup in just one command or facilitating script. Other types of tests, such as functional and integration tests. Commit history with descriptive messages.

# Observations

It wasn't really possible to meet all of the bonus points that have been brought up above. The code is still undercommented, and the comments that do exist might be part in Portuguese and part in English. No unitary or integration tests were written for this project so far because of time constraints. In my own words, if I had to sum up what I did here for this project, I would say it consists of two CRUD applications: crud-users and order-service. With CRUD operations, the first one manages user data and the second one does orders. For both of them to be used, one needs to insert a BearerToken, which one can obtain by using a third CRUD application called auth-service that manages admin (an entity very similar to users but used only for the sake of authentication) data. One can do all typical CRUD operations with both crud-users and auth-service, but the latter has an extra endpoint: GenerateToken. With order-service, one can do all operations with the sole exception of updating an order. The deletion of an order is restricted only to limit orders. When creating a order with order-service, the application sends a request to crud-users to check if the user whose ID one inserted into the payload exists. If he or she does, then his or her name is also returned.
Docker was used indeed with four images and four containers being set up: one for the MySQL database, a second one for the auth-service, a third one for crud-users and a fourth and last one for order-service.
I've also uploaded a set of Postman collections (3/one for each service) and an environment. They're in the ./postman folder.
The map of ports is the following: 8082 -> auth-service, 8080 -> crud-users and 8081 -> order-service

HOW TO RUN THIS:

GO TO THE ROOT DIRECTORY (IN THIS CASE desafio-americas-technology) AND EXECUTE THIS: 

```zsh
#!/bin/zsh

chmod +x start_services.sh
```

THEN RUN THIS:

```zsh
#!/bin/zsh

./start_services.sh
```

Some considerations about the packages/libraries most used in this project:

## External Packages Used in the Project

1. **gorm (gorm.io/gorm):**
   - This is a popular Go ORM (Object Relational Mapper) used for database operations.
   - In the project, it's used for working with the MySQL database, defining models, and performing CRUD operations.

2. **gin (github.com/gin-gonic/gin):**
   - Gin is a web framework for Go, widely used for building web applications and APIs.
   - In the project, it's used to define the API routes, handle HTTP requests, and manage middleware.
   - The decision to use Gin was influenced by previous experience, making it a straightforward choice for this project.

3. **uuid (github.com/google/uuid):**
   - The UUID package is used for working with Universally Unique Identifiers (UUIDs).
   - In the project, it's used for generating unique identifiers for the entities.

4. **time (time):**
   - The time package is part of the standard library and is used for handling time-related operations.
   - In the project, it's used for setting timestamps in database records and managing token expiration times.

5. **os (os):**
   - Another standard library package providing a way to interact with the operating system.
   - In the project, it's used for retrieving values from environment variables.

6. **swaggo/gin-swagger (github.com/swaggo/gin-swagger):**
   - This package is an integration of Swagger for Go applications using the Gin framework.
   - In the project, it's used to generate API documentation.
   - Note: There were challenges with properly integrating it into the auth-service's documentation (but crud-users's and order-service's fine though); investigation ongoing.

7. **swaggerFiles (github.com/swaggo/files):**
   - This package is part of the Swagger tools for Go and is used for serving Swagger UI files.
   - In the project, it's used for serving Swagger documentation.


