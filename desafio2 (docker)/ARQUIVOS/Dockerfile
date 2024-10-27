# Etapa de construção
FROM golang:1.19 AS builder

# Diretório de trabalho dentro do container
WORKDIR /app

# Copia o código Go para o diretório de trabalho
COPY . .

# Compila o código Go
RUN go mod init app && go mod tidy && go build -o app main.go

# Etapa final
FROM golang:1.19

# Diretório de trabalho no container
WORKDIR /app

# Copia o binário gerado na etapa de construção
COPY --from=builder /app/app /app/app
COPY --from=builder /app/random-image.png /app/random-image.png
COPY --from=builder /app/encrypted1.png /app/encrypted1.png

# Define um volume para compartilhar arquivos com o host
VOLUME /app/output

# Define o comando padrão para rodar o app
CMD ["./app"]
