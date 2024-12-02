

---


**Quarkus** là một framework Java hiện đại được thiết kế đặc biệt cho Kubernetes và môi trường đám mây. Nó tối ưu hóa cho thời gian khởi động nhanh, kích thước bộ nhớ nhỏ và tích hợp tốt với các công nghệ Java phổ biến như Jakarta EE, Hibernate, và MicroProfile.

## **1. Giới thiệu về Quarkus**

### **Tại sao chọn Quarkus?**

- **Thời gian khởi động nhanh**: Tích hợp GraalVM để tạo native image.
- **Bộ nhớ nhỏ gọn**: Giảm footprint của ứng dụng.
- **Lập trình phản ứng (Reactive)**: Hỗ trợ tốt cho các ứng dụng hiệu năng cao.
- **Dành cho đám mây**: Tích hợp trực tiếp với Kubernetes và Docker.
- **Developer Experience (DX)**: Hỗ trợ chế độ Dev Mode cho phép tự động reload code.

### **Yêu cầu cài đặt:**

1. **JDK 11 trở lên.**
2. **Maven hoặc Gradle.**
3. **Quarkus CLI** _(tùy chọn)_:
```bash
brew install quarkusio/tap/quarkus
```

## **2. Bắt đầu với Quarkus**

### **a. Tạo ứng dụng Quarkus**

1. **Tạo project với Maven:**
```bash
mvn io.quarkus:quarkus-maven-plugin:create \
   -DprojectGroupId=com.example \
   -DprojectArtifactId=quarkus-demo \
   -DclassName="com.example.GreetingResource" \
   -Dpath="/hello"
```

- `-Dpath="/hello"`: Cấu hình endpoint ban đầu.


**Cấu trúc dự án:**

```bash
quarkus-demo
├── src
│   ├── main
│   │   ├── java/com/example/GreetingResource.java
│   │   ├── resources/application.properties
│   │   └── resources/META-INF/resources

```

### **b. Chạy ứng dụng trong chế độ Dev Mode**

```bash
cd quarkus-demo
./mvnw compile quarkus:dev
```

- Truy cập: `http://localhost:8080/hello`.

## **4. Các tính năng nổi bật của Quarkus**


### **a. Native Image với GraalVM**

Quarkus hỗ trợ native image để cải thiện hiệu năng và giảm bộ nhớ sử dụng:

```
./mvnw package -Pnative
./target/quarkus-demo-1.0.0-runner

```

### **b. Reactive Programming**

Tạo ứng dụng phi đồng bộ (non-blocking):

```java
@GET
@Produces(MediaType.TEXT_PLAIN)
public Uni<String> hello() {
    return Uni.createFrom().item("Hello, Reactive Quarkus!");
}

```

### **c. Tích hợp Kubernetes**

1.**Thêm extension Kubernetes**:
```
./mvnw quarkus:add-extension -Dextensions="kubernetes"
```

**Cấu hình Kubernetes**:
```
quarkus.kubernetes.namespace=default
quarkus.kubernetes.name=quarkus-demo

```

**Triển khai lên Kubernetes**:
```
./mvnw clean package -Dquarkus.kubernetes.deploy=true

```
### **1. Tổng quan**

|**Tiêu chí**|**Quarkus**|**Spring Boot**|
|---|---|---|
|**Mục tiêu**|Tối ưu hóa cho Cloud và Kubernetes, đặc biệt trong môi trường serverless.|Framework toàn diện, linh hoạt cho mọi loại ứng dụng.|
|**Triết lý thiết kế**|Native-first (hỗ trợ tốt với GraalVM và native image).|JVM-first (tối ưu cho JVM, nhưng vẫn hỗ trợ native).|
|**Hệ sinh thái**|Nhẹ, tập trung vào microservices và đám mây.|Lớn và toàn diện, hỗ trợ cả monolith và microservices.|

### **2. Hiệu năng**

| **Tiêu chí**               | **Quarkus**                                    | **Spring Boot**                                    |
| -------------------------- | ---------------------------------------------- | -------------------------------------------------- |
| **Thời gian khởi động**    | Nhanh hơn đáng kể (đặc biệt với native image). | Chậm hơn do tập trung vào JVM runtime.             |
| **Sử dụng bộ nhớ**         | Ít hơn, tối ưu cho container-based workloads.  | Tốn bộ nhớ hơn, phù hợp với JVM runtime.           |
| **Native Image (GraalVM)** | Tích hợp chặt chẽ, dễ dàng tạo native image.   | Hỗ trợ nhưng cần cấu hình bổ sung (Spring Native). |
| **Non-blocking I/O**       | Hỗ trợ lập trình Reactive mạnh mẽ (Vert.x).    | Hỗ trợ Reactive tốt (Spring WebFlux).              |
### **3. Developer Experience**

| **Tiêu chí**               | **Quarkus**                                          | **Spring Boot**                                  |
| -------------------------- | ---------------------------------------------------- | ------------------------------------------------ |
| **Chế độ Dev Mode**        | Hỗ trợ tự động reload code nhanh và mượt.            | Spring DevTools hỗ trợ tương tự, nhưng chậm hơn. |
| **Hệ sinh thái extension** | Tích hợp nhiều extension nhưng không lớn như Spring. | Ecosystem khổng lồ, hỗ trợ đa dạng use case.     |
| **Cộng đồng và tài liệu**  | Cộng đồng đang phát triển, nhiều hướng dẫn.          | Cộng đồng lớn, tài liệu phong phú.               |
### **4. Triển khai**

| **Tiêu chí**                   | **Quarkus**                                                                       | **Spring Boot**                                       |
| ------------------------------ | --------------------------------------------------------------------------------- | ----------------------------------------------------- |
| **Triển khai trên Kubernetes** | Tích hợp gốc (native) với Kubernetes và OpenShift (quarkus-kubernetes extension). | Cần thêm cấu hình (e.g., Spring Cloud Kubernetes).    |
| **Tối ưu hóa Cloud Native**    | Rất phù hợp, nhờ khởi động nhanh và nhẹ.                                          | Hỗ trợ tốt nhưng yêu cầu thêm cấu hình và tài nguyên. |
| **Hỗ trợ Serverless**          | Phù hợp với môi trường serverless như AWS Lambda, Azure Functions.                | Hỗ trợ nhưng không tối ưu như Quarkus.                |
### **5. Công nghệ và API hỗ trợ**

| **Tiêu chí**             | **Quarkus**                                            | **Spring Boot**                        |
| ------------------------ | ------------------------------------------------------ | -------------------------------------- |
| **REST API**             | Hỗ trợ JAX-RS natively (Jakarta RESTful Web Services). | Hỗ trợ Spring MVC hoặc Spring WebFlux. |
| **Dependency Injection** | CDI (Context and Dependency Injection).                | Spring Framework (Spring IoC).         |
| **ORM**                  | Hibernate ORM (JPA).                                   | Hibernate ORM (JPA).                   |
| **Reactive Programming** | Vert.x-based Reactive Stack.                           | Project Reactor-based stack (WebFlux). |
| **Security**             | Quarkus Security hoặc Keycloak integration.            | Spring Security (toàn diện hơn).       |
### **6. Hệ sinh thái và Tích hợp**

| **Tiêu chí**              | **Quarkus**                                                                                 | **Spring Boot**                                         |
| ------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------- |
| **Hệ sinh thái lớn**      | Tập trung vào các công nghệ cloud-native và container-first như Kubernetes, Kafka, GraalVM. | Hỗ trợ cả monolith lẫn microservices.                   |
| **Hỗ trợ thư viện**       | Nhiều extension tích hợp sẵn: Hibernate, Kafka, RESTEasy, Kubernetes.                       | Spring Starter Projects với hàng trăm dependency.       |
| **Tích hợp MicroProfile** | Tích hợp Jakarta EE và MicroProfile chặt chẽ.                                               | Không hỗ trợ MicroProfile (dùng Spring Cloud thay thế). |
### **7. Khi nào nên dùng?**

| **Trường hợp sử dụng**                 | **Quarkus**                                                                 | **Spring Boot**                                     |
| -------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------- |
| **Ứng dụng Cloud Native**              | Rất phù hợp nhờ thời gian khởi động nhanh và footprint nhỏ.                 | Có thể sử dụng, nhưng tốn tài nguyên hơn.           |
| **Ứng dụng monolithic**                | Có thể dùng nhưng không phải ưu tiên chính.                                 | Rất phù hợp với các ứng dụng lớn, phức tạp.         |
| **Serverless**                         | Lý tưởng, đặc biệt khi tích hợp với AWS Lambda hoặc Google Cloud Functions. | Có thể sử dụng, nhưng không tối ưu bằng Quarkus.    |
| **Hệ thống doanh nghiệp (Enterprise)** | Hỗ trợ tốt, đặc biệt khi dùng với MicroProfile.                             | Lựa chọn hàng đầu, nhờ hệ sinh thái lớn và ổn định. |
| **Dự án lâu dài, nhiều thành viên**    | Tốt, nhưng cần nhiều thời gian học tập.                                     | Tốt hơn vì có nhiều tài liệu và cộng đồng lớn.      |
### **8. Tóm tắt**

| **Tiêu chí**   | **Quarkus**                                  | **Spring Boot**                                          |
| -------------- | -------------------------------------------- | -------------------------------------------------------- |
| **Ưu điểm**    | - Nhanh và nhẹ.                              | - Hệ sinh thái rộng lớn, phù hợp với mọi nhu cầu.        |
|                | - Tích hợp tốt với Kubernetes, Cloud Native. | - Dễ học và phổ biến trong cộng đồng.                    |
|                | - Dev Mode mạnh mẽ, GraalVM Native tốt.      | - Hỗ trợ đầy đủ tính năng từ monolith đến microservices. |
| **Nhược điểm** | - Cộng đồng chưa lớn như Spring Boot.        | - Khởi động chậm hơn, tiêu tốn tài nguyên hơn.           |