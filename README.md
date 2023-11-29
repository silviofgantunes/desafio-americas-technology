# desafio-americas-technology

Projeto de Desenvolvimento de Microserviços para Cadastro de Usuários e para Ordens de Compra e Venda de Criptoativos

O propósito deste teste é avaliar suas habilidades em codificação e arquitetura em um ambiente de desenvolvimento. Você receberá um problema simples para demonstrar suas técnicas de programação, e é recomendado que você exagere um pouco na solução para destacar suas capacidades. A aplicação a ser desenvolvida deve estar pronta para produção, levando em consideração a manutenção por outros desenvolvedores ao longo do tempo.

Requisitos das Aplicações:

Cadastro de Usuários:

Listar, exibir, criar, alterar e excluir usuários.
Tabela de usuários (user): id, name, email, phone_number, created_at, updated_at.

Serviço de Ordens:

Listar, listar por usuário, exibir, criar, excluir ordem (limit).
Tabela de ordens (compra e venda de crypto) (order): id, user_id, pair, amount, direction, type (market, limit), created_at, updated_at.
Comunicação entre os serviços para garantir a consistência de dados.
Padrões REST e documentação contendo endpoints e payloads utilizados.

Critérios de Avaliação Destacados:

Uso de bibliotecas de terceiros e, possivelmente, de um framework, com justificação.
Execução em containers Docker.
Retorno de JSON válido e presença dos recursos citados.
Código testável com testes unitários.
Adesão a melhores práticas de segurança de APIs e diretrizes de estilo de código.

Pontos Considerados Bônus:

Respostas durante o code review e descrição detalhada na "pull request".
Setup da aplicação em apenas um comando ou script facilitador.
Outros tipos de testes, como funcionais e de integração.
Histórico de commits com mensagens descritivas.
