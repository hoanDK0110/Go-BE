# Sử dụng một hình ảnh golang đã có sẵn làm cơ sở
FROM golang:latest

# Thiết lập thư mục làm việc trong container
WORKDIR /app

# Sao chép mã nguồn ứng dụng vào thư mục /app trong container
COPY . .

# Biên dịch ứng dụng Golang
RUN go build -o main .

# Sử dụng một hình ảnh nhẹ để triển khai ứng dụng
FROM alpine:latest

# Sao chép tệp thực thi từ ứng dụng Golang đã biên dịch vào hình ảnh
COPY --from=0 /app/main .

# Mở cổng mà ứng dụng lắng nghe (thay thế 8080 bằng cổng của ứng dụng của bạn)
EXPOSE 8080

# Khởi chạy ứng dụng khi container được khởi động
CMD ["./main"]
