# PontoMenos-API
O PontoMenos-API é uma API robusta desenvolvida para facilitar o gerenciamento de horários de trabalho, permitindo a autenticação de usuários, registro de ponto (entrada, intervalos e saída) e geração de relatórios mensais dos pontos registrados. Este projeto foi construído utilizando Go, Gin Gonic, GORM, RabbitMQ e PostgreSQL.

## Desenho da solução "Fase01" (implementado)
  ![RegistroDePonto](https://github.com/PedroVCorsino/pontomenos-api/assets/61948860/97393de8-2f9c-4799-b8cc-5ddc3f30bfb1)
 - A autenticação do usuario é feita pelo proprio sistema que utiliza a biblioteca [jwt-go](https://pkg.go.dev/github.com/golang-jwt/jwt/v5@v5.2.1#section-readme) para validar o dados do user e devolver um token JWT que será usado para autorizar as requisiçoes.
 - A API do PontoMenos recebe requisições para o registro de ponto e as emcaminha para uma fila (RabbitMQ) responsavel por lidar com altas cargas de acesso, especialmente durante os picos. 
 - O Sistema PontoMenos consome a fila e processa, as informaçoes registra em um banco de dados (PostgreSQL)

### Funcionalidades
- **Autenticação de Usuário:** Os usuários podem se autenticar usando um nome de usuário ou matrícula e senha.
- **Registro de Ponto:** Permite o registro do horário de entrada, intervalos e saída, capturando automaticamente a data e hora do registro.
- **Visualização de Registros:** Os usuários podem visualizar seus registros de ponto, incluindo detalhes como data, horários de entrada, intervalos, saída e total de horas trabalhadas.
- **Relatórios:** Geração de relatórios mensais (espelho de ponto), enviando-os por e-mail aos solicitantes.

## Desenho da solução "Fase02" (Não implementado)
  ![EvolucaoRegistroDePonto png](https://github.com/PedroVCorsino/pontomenos-api/assets/61948860/06d80711-2146-40e5-baca-8a4107666b72)
- Nessa fase separamos as responsabilidades em 2 microsserviços
- Uma API responsavel por receber e devolver informaçoes ao font-end, a mesma pode receber requisiçoes de diversos microsserviços, seja um aplicativo mobile ou um sistema web.
- Teremos um serviço consumidor exclusivo para o registro de pontos, responsavel apenas por consumir e processar a fila de registros.

### Funcionalidades propostas
  - **Edição de Registros:** Em caso de erros, o sistema deve permitir que usuários editem seus registros de ponto. No entanto, essa funcionalidade deve ser restrita e possivelmente necessitar de aprovação de um supervisor.
- **Notificações:** O sistema deve ser capaz de enviar notificações para lembrar os usuários de registrar seu ponto.
- **Administração:** Deve haver uma interface de administração para gerenciar usuários, aprovar edições de registros e visualizar relatórios.
- **Relatórios:** O sistema deve ser capaz de gerar os relatórios com visão administrativa, para uso dos gestores

## Tecnologias Utilizadas

- **Go:** Linguagem de programação usada para desenvolver a lógica do sistema.
  - Go foi escolhida pela sua eficiência, simplicidade e suporte robusto a concorrência, facilitando o desenvolvimento de um sistema de registro de ponto escalável e de alta performance.
- **Gin Gonic:** Framework web usado para facilitar a criação de rotas e manipulação de requisições HTTP.
  - O [Gin](https://gin-gonic.com/) foi escolhido por ser um framework web de alta performance para Go, oferecendo uma abordagem simples e eficaz para a construção de APIs REST, crucial para a rápida resposta e manipulação de registros de ponto. 
- **GORM:** ORM para Go, utilizado para interagir com o banco de dados PostgreSQL.
  - [Gorm](https://gorm.io/index.html) é utilizado por sua abstração poderosa e desenvolvimento ágil de operações de banco de dados, compatível com Go, facilitando a interação com PostgreSQL de maneira eficiente e produtiva. 
- **RabbitMQ:** Sistema de mensagens que controla as requisições, assegurando que sejam processadas eficientemente.
  - RabbitMQ foi escolhido para gerenciar filas de mensagens de forma confiável, permitindo o processamento assíncrono e a escalabilidade do sistema de registros de ponto, garantindo entrega eficiente e ordenada das mensagens. 
- **PostgreSQL:** Sistema de gerenciamento de banco de dados relacional para armazenamento dos dados.
  - PostgreSQL foi selecionado por sua robustez, escalabilidade e suporte a transações complexas, oferecendo um armazenamento seguro e eficiente para os dados de registro de ponto.   

## Pré-requisitos

Antes de iniciar, certifique-se de ter instalado em sua máquina:
- Go (versão 1.22.1 ou superior)
- Docker e Docker Compose (para RabbitMQ e PostgreSQL)
