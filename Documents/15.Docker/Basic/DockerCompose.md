
---

### Docker compose là gì ?

Docker compose là công cụ dùng để định nghĩa và run multi-container cho Docker application. Với compose bạn sử dụng file YAML để config các services cho application của bạn. Sau đó dùng command để create và run từ những config đó. Sử dụng cũng khá đơn giản chỉ với ba bước:

- Khai báo app’s environment trong Dockerfile.
- Khai báo các services cần thiết để chạy application trong file docker-compose.yml.
- Run docker-compose up để start và run app.
---

Không giống như Dockerfile (build các image). Docker compose dùng để build và run các container. Các thao tác của docker-compose tương tự như lệnh: docker run.

Docker compose cho phép tạo nhiều service(container) giống nhau bằng lệnh:

```none
$ docker-compose scale <tên service> = <số lượng>
```


---
### Demo

Dockerfile

```dockerfile
FROM node:carbon-alpine AS node_builder

WORKDIR /app/webreactjs
COPY /webreactjs/package.json .
RUN npm install
COPY /webreactjs .
LABEL name="webreactjs" version="1.0"
EXPOSE 3000
CMD ["npm", "start"]

FROM golang:1.11 AS go_builder
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .
LABEL name="server" version="1.0"

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=go_builder /main ./
RUN chmod +x ./main
EXPOSE 8080
CMD ./main

```

Tiến hành build Dockerfile

Config services cần start và run trong file **docker-compose.yml**


```yml
version: '2.1'

services:
  webreactjs:
    image: af1205224676
    build: .
    ports:
      - 3000:3000
    restart: always
  servergo:
    image: cef5deda0834
    build: .
    ports:
      - 8080:8080
    restart: always

```

- **version**: chỉ ra phiên bản docker-compose đã sử dụng.
- **services**: thiết lập các services(containers) muốn cài đặt và chạy.
- **image**: chỉ ra image được sử dụng trong lúc tạo ra container.
- **build**: dùng để tạo container.
- **ports**: thiết lập ports chạy tại máy host và trong container.
- **restart**: tự động khởi chạy khi container bị tắt.

Ngoài ra còn có một số lệnh config khác:

**environment**: thiết lập biến môi trường ( thường sử dụng trong lúc config các thông số của db).

**depends_on**: chỉ ra sự phụ thuộc. Tức là services nào phải được cài đặt và chạy trước thì service được config tại đó mới được chạy.

**volumes**: dùng để mount hai thư mục trên host và container với nhau.

Run command như bên dưới:

```none
$ docker-compose up
```

