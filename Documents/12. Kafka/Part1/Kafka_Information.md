
---
# **1. Hiểu Tổng Quan Về Kafka**

## ✅ **Kafka là gì?**

Kafka là một **Message Broker** theo mô hình **Pub/Sub (Event-Driven Architecture)**, được thiết kế để xử lý **dữ liệu streaming** với **tốc độ cao**, **bền vững** và **scalable**.

### 🎯 **Tại sao Kafka được sử dụng rộng rãi?**

- **Độ bền dữ liệu cao**: Kafka lưu trữ dữ liệu theo cơ chế **commit log**, giúp dữ liệu không bị mất nếu consumer chưa xử lý kịp.
- **Hiệu suất cao**: Hỗ trợ **hàng triệu message/giây** với cơ chế **batching, zero-copy**, tối ưu I/O.
- **Scalability**: Dễ dàng scale horizontally với **partitioning**.
- **Event-Driven**: Hỗ trợ kiến trúc **event-driven**, giúp microservices giao tiếp không đồng bộ.
- **Stream Processing**: Kafka + **Kafka Streams, Flink, Spark Streaming** hỗ trợ xử lý dữ liệu real-time.

---

## ✅ **So sánh Kafka với RabbitMQ, ActiveMQ**

|Tính năng|Kafka|RabbitMQ / ActiveMQ|
|---|---|---|
|**Mô hình**|**Pub/Sub (log-based)**|**Message Queue** (Push-based)|
|**Khả năng lưu trữ**|Lưu trữ message theo log (tái sử dụng được)|Message bị xóa sau khi consumer xử lý|
|**Scalability**|**Tốt (Partitioning, Replication)**|Kém hơn do queue phải xử lý theo thứ tự|
|**Performance**|**Cao** (hàng triệu TPS)|Thấp hơn Kafka|
|**Ordering**|Đảm bảo **ordering per partition**|Không đảm bảo nếu có nhiều consumer|
|**Use Case**|Big Data, Streaming, Log Aggregation|Task Queue, Event-driven Microservices|

📌 **Khi nào dùng Kafka?**

- Xử lý **dữ liệu lớn**, cần **tốc độ cao**.
- **Event-driven architecture** (Microservices, Logging, IoT).
- **Real-time streaming** (AI, Fraud Detection, Monitoring).

📌 **Khi nào dùng RabbitMQ / ActiveMQ?**

- Cần đảm bảo **message delivery** (transactional message, email queue).
- **RPC pattern** giữa các microservices.

---

## ✅ **Các thành phần chính trong Kafka**

### 1️⃣ **Producer**

- Gửi message vào Kafka.
- Hỗ trợ **batching, compression** để tối ưu hiệu suất.

### 2️⃣ **Broker**

- Kafka cluster gồm nhiều **broker (server)** lưu trữ message.
- Broker lưu message theo **log-based storage**.

### 3️⃣ **Topic**

- Message được lưu trong **topic**.
- **Mỗi topic có thể có nhiều partition**.

### 4️⃣ **Partition**

- **Chia nhỏ topic** để **scale horizontally**.
- Mỗi partition được **replicate** để đảm bảo **fault tolerance**.

### 5️⃣ **Consumer**

- Đọc dữ liệu từ Kafka topic.
- Kafka hỗ trợ **pull-based model** → Consumer chủ động đọc message.

### 6️⃣ **Consumer Group**

- Tập hợp nhiều consumer để **load balancing**.
- Kafka đảm bảo **mỗi partition chỉ có 1 consumer trong nhóm xử lý**.

### 7️⃣ **ZooKeeper / KRaft**

- **ZooKeeper** (Kafka cũ) quản lý **metadata, leader election**.
- **Kafka mới (3.0+) dùng KRaft** thay thế ZooKeeper → Tăng hiệu suất & loại bỏ dependency.

---

📌 **Tóm tắt:**  
Kafka là một message broker mạnh mẽ, phù hợp cho **big data, event-driven microservices, stream processing**. Nếu hệ thống của bạn cần **high throughput, scalable, real-time event processing**, Kafka là lựa chọn phù hợp. 🚀