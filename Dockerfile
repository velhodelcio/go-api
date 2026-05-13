FROM golang:1.26.3

# Set working directory
WORKDIR /go/src/app

# Instala o Air para Hot Reload
RUN go install github.com/air-verse/air@latest

# Copia apenas os arquivos de dependencia primeiro (otimiza o cache do Docker)
COPY go.mod go.sum ./
RUN go mod download

# Copia o resto do codigo
COPY . .

EXPOSE 8000

# O comando padrao agora eh rodar o Air, que observa mudancas e faz o build
CMD ["air", "-c", ".air.toml"]