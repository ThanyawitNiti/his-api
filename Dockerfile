# ใช้ Golang เป็น base image
FROM golang:1.20

# กำหนด working directory
WORKDIR /app

# Copy ไฟล์ทั้งหมดเข้า container
COPY . .

# Download dependencies
RUN go mod tidy

# Build Go service
RUN go build -o main .

# Run API
CMD ["/app/main"]
