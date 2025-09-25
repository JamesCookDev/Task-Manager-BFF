# BFF (Backend for Frontend) - Gerenciador de Tarefas

Este repositório contém o serviço de **Backend for Frontend (BFF)** para o projeto Gerenciador de Tarefas. Construído em Go (Golang), este serviço atua como uma camada intermediária otimizada entre a interface do usuário (front-end) e a API principal (back-end).

---

## 🏗️ Papel na Arquitetura

O BFF é uma peça central na nossa arquitetura de microsserviços, projetado para:

1.  **Simplificar a Lógica do Cliente:** O front-end (React) faz chamadas simples e diretas para este BFF, sem precisar conhecer a complexidade da API principal do Django.
2.  **Atuar como um Gateway:** É o único ponto de entrada para o front-end, o que melhora a segurança e o controle.
3.  **Otimizar a Comunicação:** O BFF é responsável por se comunicar com a API Core (Django) de forma eficiente. No futuro, ele também pode ser responsável por agregar dados de múltiplas fontes ou implementar uma camada de cache para aumentar a performance.

O fluxo de comunicação é: `Frontend (React) <--> BFF (Go) <--> Backend (Django)`.

---

## ✅ Principais Funcionalidades

* **Servidor HTTP Leve e Performático:** Construído com a biblioteca padrão do Go e o roteador `chi`.
* **Proxy para a API Core:** Busca dados da API principal do Django de forma autenticada e segura.
* **Estrutura Organizada:** O código é dividido em pacotes (`handlers`, `clients`) para separação de responsabilidades e fácil manutenção.
* **Configuração Flexível:** Lê configurações de variáveis de ambiente, facilitando a execução tanto localmente quanto em contêineres Docker.

---

## 🛠️ Tecnologias Utilizadas

| Tecnologia | Propósito |
| :--- | :--- |
| **Go (Golang)** | Linguagem de programação principal, escolhida por sua performance e concorrência. |
| **Chi** | Roteador HTTP leve e idiomático para a definição dos endpoints. |
| **Godotenv** | Gerenciamento de variáveis de ambiente para desenvolvimento local. |

---

## 🚀 Como Executar o Projeto

Existem duas maneiras de executar o BFF: localmente para desenvolvimento rápido ou como parte do ecossistema Docker completo.

### Pré-requisitos

* [Go (versão 1.21+ an
terior)](https://go.dev/dl/)
* O serviço de **Back-end Django** deve estar rodando para que o BFF possa se comunicar com ele.

### 1. Rodando Localmente (com `go run`)

Este modo é ideal para desenvolver e testar o BFF rapidamente.

1.  **Clone o repositório:**
    ```bash
    git clone [https://github.com/JamesCookDev/gerenciador-tarefas-bff.git](https://github.com/JamesCookDev/gerenciador-tarefas-bff.git)
    cd gerenciador-tarefas-bff
    ```
2.  **Crie e configure o arquivo `.env`:**
    Copie o template de exemplo e preencha as variáveis.
    ```bash
    cp .env.example .env
    ```
    * Abra o arquivo `.env` e adicione um **token JWT de longa duração** válido, gerado pela API do Django. A variável `DJANGO_API_URL` já deve estar configurada para `localhost`.

3.  **Execute o servidor:**
    ```bash
    go run main.go
    ```
    O servidor BFF estará rodando em `http://localhost:8080`.

### 2. Rodando com Docker Compose (Recomendado)

Este serviço foi projetado para ser orquestrado pelo `docker-compose.yml` principal do projeto.

1.  **Siga as instruções** no `README.md` do repositório do **back-end** (`gerenciador-tarefas-backend`).
2.  O comando `docker-compose up` na raiz do projeto irá construir e iniciar automaticamente o contêiner do BFF junto com os outros serviços.
3.  Neste modo, o BFF se comunicará com o back-end usando o nome do serviço (`http://backend:8000/...`) através da rede interna do Docker.

---

## 🗺️ Endpoints Expostos

| Endpoint | Método HTTP | Descrição |
| :--- | :--- | :--- |
| `/api/projetos` | `GET` | Retorna a lista de projetos do usuário, buscando os dados da API Core do Django. |

---

