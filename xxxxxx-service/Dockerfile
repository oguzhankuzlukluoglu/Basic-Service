# Go'nun resmi hafif bir versiyonunu kullanıyoruz
FROM golang:1.20-alpine

# Uygulama dosyalarının konumunu belirliyoruz
WORKDIR /app

# go mod ve sum dosyalarını kopyalıyoruz (bağımlılıkları çekmek için)
COPY go.mod go.sum ./
RUN go mod download

# Uygulama dosyalarını kopyalıyoruz
COPY . .

# Uygulamayı derliyoruz
RUN go build -o main

# Uygulamayı çalıştırıyoruz
CMD ["./main"]

# Portu açıyoruz (servisin portunu belirtiyoruz)
EXPOSE 8083
