
---
Dockerfile là file config cho Docker để build ra image. Nó dùng một image cơ bản để xây dựng lớp image ban đầu. Một số image cơ bản: python, unbutu and alpine. Sau đó nếu có các lớp bổ sung thì nó được xếp chồng lên lớp cơ bản. Cuối cùng một lớp mỏng có thể được xếp chồng lên nhau trên các lớp khác trước đó.

Các config :

- **FROM** — chỉ định image gốc: python, unbutu, alpine…
- **LABEL** — cung cấp metadata cho image. Có thể sử dụng để add thông tin maintainer. Để xem các label của images, dùng lệnh docker inspect.
- **ENV** — thiết lập một biến môi trường.
- **RUN** — Có thể tạo một lệnh khi build image. Được sử dụng để cài đặt các package vào container.
- **COPY** — Sao chép các file và thư mục vào container.
- **ADD** — Sao chép các file và thư mục vào container.
- **CMD** — Cung cấp một lệnh và đối số cho container thực thi. Các tham số có thể được ghi đè và chỉ có một CMD.
- **WORKDIR** — Thiết lập thư mục đang làm việc cho các chỉ thị khác như: RUN, CMD, ENTRYPOINT, COPY, ADD,…
- **ARG** — Định nghĩa giá trị biến được dùng trong lúc build image.
- **ENTRYPOINT** — cung cấp lệnh và đối số cho một container thực thi.
- **EXPOSE** — khai báo port lắng nghe của image.
- **VOLUME** — tạo một điểm gắn thư mục để truy cập và lưu trữ data.
---
Demo:

```dockerfile
# Use an official OpenJDK runtime as a parent image
FROM openjdk:17-jdk-slim

# Set the working directory in the container
WORKDIR /app

# Copy the application JAR file to the container
COPY target/your-spring-boot-app.jar app.jar

# Make port 8080 available to the world outside this container
EXPOSE 8080

# Run the Spring Boot application
ENTRYPOINT ["java", "-jar", "app.jar"]

```

```sh
docker build -t your-spring-boot-app .
```

```sh
docker run -p 8080:8080 your-spring-boot-app
```

