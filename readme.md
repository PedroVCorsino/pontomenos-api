# Sistema de Registro de Ponto

O PontoMenos-API é uma API robusta desenvolvida para facilitar o gerenciamento de horários de trabalho, permitindo a autenticação de usuários, registro de ponto (entrada, intervalos e saída) e geração de relatórios mensais dos pontos registrados. Este projeto foi construído utilizando Go, Gin Gonic, GORM, RabbitMQ e PostgreSQL.

## Funcionalidades

- **Autenticação de Usuário:** Os usuários podem se autenticar usando um nome de usuário ou matrícula e senha.
- **Registro de Ponto:** Permite o registro do horário de entrada, intervalos e saída, capturando automaticamente a data e hora do registro.
- **Visualização de Registros:** Os usuários podem visualizar seus registros de ponto, incluindo detalhes como data, horários de entrada, intervalos, saída e total de horas trabalhadas.
- **Relatórios:** Geração de relatórios mensais (espelho de ponto), enviando-os por e-mail aos solicitantes.
- **Segurança:** Garantia de que os dados dos usuários sejam armazenados de forma segura e que sua privacidade seja mantida.
- **Disponibilidade:** Sistema disponível 24/7, com um tempo de resposta para serviços de marcação de ponto de até 5 segundos.

## Tecnologias Utilizadas

- **Go:** Linguagem de programação usada para desenvolver a lógica do sistema.
- **Gin Gonic:** Framework web usado para facilitar a criação de rotas e manipulação de requisições HTTP.
- **GORM:** ORM para Go, utilizado para interagir com o banco de dados PostgreSQL.
- **RabbitMQ:** Sistema de mensagens que controla as requisições, assegurando que sejam processadas eficientemente.
- **PostgreSQL:** Sistema de gerenciamento de banco de dados relacional para armazenamento dos dados.

## Pré-requisitos

Antes de iniciar, certifique-se de ter instalado em sua máquina:
- Go (versão 1.22.1 ou superior)
- Docker e Docker Compose (para RabbitMQ e PostgreSQL)

## Configuração e Instalação

1. Clone o repositório:
```bash
git clone <url-do-repositorio>

go run main.go
```


