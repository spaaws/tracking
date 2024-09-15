# Etapa 1: Usando a imagem oficial do Go para construir a aplicação
FROM golang:1.20-alpine

# Definindo o diretório de trabalho dentro do container
WORKDIR /app

# Copiando os arquivos da aplicação para o container
COPY . .

# Baixando dependências
RUN go mod download

# Construindo a aplicação Go
RUN go build -o app main.go

# Expondo a porta 8080
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./app"]
