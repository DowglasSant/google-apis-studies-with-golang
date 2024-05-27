# 🌐 Projeto de Tradução e Análise de Sentimento

![Go](https://img.shields.io/badge/Go-1.22.3-blue)
![Google Cloud](https://img.shields.io/badge/Google%20Cloud-Natural%20Language%20API-yellow)
![Google Cloud](https://img.shields.io/badge/Google%20Cloud-Translation%20API-yellow)

Este é um projeto em Go (Golang) que utiliza a API de Linguagem Natural e a API de Tradução da Google Cloud para realizar tradução e análise de sentimentos em textos.

## 🛠️ Funcionalidades

- **Leitura de Arquivo**: O sistema lê um arquivo de texto do diretório especificado.
- **Tradução (Opcional)**: Ele identifica automaticamente o idioma do texto. Se o idioma não for português, o texto é traduzido para o português brasileiro e adicionado a um novo arquivo de saída.
- **Análise de Sentimento**: Em seguida, o sistema lê o texto traduzido (ou original, se não foi traduzido) e realiza uma análise de sentimentos usando a API de Linguagem Natural da Google Cloud.
- **Exibição de Resultados**: O resultado da análise de sentimento é exibido na tela, indicando o sentimento principal do texto e sua magnitude.

## 🔧 Tecnologias Utilizadas

- **Go (Golang)**: Linguagem de programação utilizada para desenvolver a aplicação.
- **Google Cloud Natural Language API**: API usada para análise de sentimento.
- **Google Cloud Translation API**: API usada para tradução de texto.

## 🚀 Instalação e Execução

1. Certifique-se de ter o Go instalado em sua máquina.
2. Clone este repositório: `git clone https://github.com/seu-usuario/seu-projeto.git`
3. Configure as credenciais da Google Cloud seguindo as instruções do [Google Cloud Console](https://console.cloud.google.com/)
4. Execute `go mod tidy` para instalar as dependências do projeto.
5. Execute `go run cmd/main.go` para iniciar a aplicação.

## 🤝 Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir um issue ou enviar um pull request.

## 📝 Licença

Este projeto está licenciado sob a licença [MIT](https://opensource.org/licenses/MIT).
