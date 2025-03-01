
---
## ✅ **Producer là gì?**

Producer là thành phần **gửi message** vào Kafka **topic**. Nó đóng vai trò là nguồn dữ liệu trong hệ thống **event-driven**.

**Ví dụ**:

- Một hệ thống e-commerce có **Order Service**, mỗi khi có đơn hàng mới, service này sẽ **publish** event `"New Order Created"` vào Kafka.
- Các consumer như **Payment Service, Inventory Service** sẽ đọc event này và xử lý tương ứng.

## ✅ **Cách Producer gửi dữ liệu vào Kafka**

### 1️⃣ **Producer gửi message vào một Topic**

Producer chọn **topic** để gửi dữ liệu. Kafka lưu trữ message trong topic dưới dạng **log**.

```java
Properties props = new Properties();
props.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, "localhost:9092");
props.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, "org.apache.kafka.common.serialization.StringSerializer");
props.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, "org.apache.kafka.common.serialization.StringSerializer");

KafkaProducer<String, String> producer = new KafkaProducer<>(props);

ProducerRecord<String, String> record = new ProducerRecord<>("order-topic", "order-123", "New Order Created");
producer.send(record);
producer.close();
```

- `ProducerConfig.BOOTSTRAP_SERVERS_CONFIG`: Địa chỉ Kafka cluster.
- `KEY_SERIALIZER_CLASS_CONFIG`: Serialize key của message (ở đây dùng String).
- `VALUE_SERIALIZER_CLASS_CONFIG`: Serialize value của message.
- `ProducerRecord<>(topic, key, value)`: Tạo message và gửi vào topic.

### 2️⃣ **Partitioning - Chia message vào các partition**

- Mỗi topic có nhiều **partition**, producer quyết định gửi message vào **partition nào**.
- Kafka hỗ trợ 3 cách:
    - **Round-Robin (default)**: Kafka tự động phân phối đều message.
    - **Key-based (Hashing)**: Producer cung cấp key, Kafka hash key để chọn partition.
    - **Custom Partitioning**: Viết logic riêng để phân vùng message.

#### 🔹 **Ví dụ gửi message theo key (hashing-based partitioning)**
```java
ProducerRecord<String, String> record = new ProducerRecord<>("order-topic", "user-1", "Order Created");
```
- Nếu `"user-1"` luôn hash về partition 2 → Các message của `"user-1"` sẽ luôn vào partition 2.
- Giúp **giữ thứ tự message theo key**.

🔹 **Custom Partitioning**
```java
public class CustomPartitioner implements Partitioner {
    @Override
    public int partition(String topic, Object key, byte[] keyBytes, Object value, byte[] valueBytes, Cluster cluster) {
        List<PartitionInfo> partitions = cluster.partitionsForTopic(topic);
        return (key.hashCode() & Integer.MAX_VALUE) % partitions.size();
    }
}
```
- Kafka sẽ gọi **partition()** để quyết định partition nào sẽ nhận message.

### 3️⃣ **Producer Acknowledgment - Đảm bảo gửi thành công**

- **acks=0** → Không cần Kafka xác nhận → Tốc độ nhanh nhưng **có thể mất message**.
- **acks=1** → Kafka xác nhận khi **leader partition** nhận message.
- **acks=all** → Kafka xác nhận khi **tất cả replica** đã lưu message (**an toàn nhất**).

```java
props.put(ProducerConfig.ACKS_CONFIG, "all"); // Đảm bảo không mất dữ liệu
```
### 4️⃣ **Batching & Compression - Tăng hiệu suất**

- **Batching**: Gửi nhiều message cùng lúc thay vì từng cái một.
```java
props.put(ProducerConfig.BATCH_SIZE_CONFIG, 16384);
props.put(ProducerConfig.LINGER_MS_CONFIG, 5); // Delay 5ms để batch nhiều message lại
```

**Compression**: Giảm kích thước message trước khi gửi.
```java
props.put(ProducerConfig.COMPRESSION_TYPE_CONFIG, "gzip"); // Hoặc snappy, lz4
```

## ✅ **Tóm tắt**

- **Producer gửi message vào topic, Kafka lưu trữ dưới dạng log.**
- **Có thể chọn partition bằng key hoặc custom logic.**
- **Sử dụng acks=all để đảm bảo không mất dữ liệu.**
- **Tối ưu hiệu suất bằng batching & compression.**


