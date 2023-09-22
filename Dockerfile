# Sử dụng hình ảnh Golang chính thống làm cơ sở
FROM golang:latest

# Đặt thư mục làm thư mục làm việc trong hình ảnh
WORKDIR /app

# Sao chép mã nguồn của bạn vào hình ảnh
COPY . .

# Biên dịch ứng dụng Golang
RUN go build -o main .

# Mở cổng 8080
EXPOSE 80

# Chạy ứng dụng khi container được khởi chạy
CMD ["./main"]
