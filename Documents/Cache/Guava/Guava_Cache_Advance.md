
---
Guava Cache là một bộ nhớ đệm trong RAM với khả năng tùy chỉnh cao. Để hiểu sâu về nó, ta cần đào sâu vào **core** của Guava Cache, bao gồm:

1. **Cấu trúc dữ liệu bên trong (`ConcurrentMap`, `Segment`)**
2. **Flow hoạt động của `get(key)`, `put(key, value)`, `invalidate(key)`**
3. **Cơ chế loại bỏ phần tử (Eviction) & Cập nhật tự động (`refreshAfterWrite`)**
4. **Cách Guava Cache xử lý `concurrency`**

Hãy cùng đi sâu vào từng phần! 🚀

---
## **1️⃣ Kiến Trúc Bên Trong Guava Cache**

Guava Cache không chỉ là một `ConcurrentHashMap` đơn thuần mà nó sử dụng **mô hình phân đoạn (`Segment`)** để tối ưu hiệu suất.

### 🏗 **Cấu trúc dữ liệu bên trong**

📌 **Guava Cache sử dụng `LocalCache<K, V>` làm lớp chính**

- **`LocalCache<K, V>`** → Quản lý toàn bộ cache
- **`Segment<K, V>`** → Chia cache thành nhiều phần nhỏ (giống `ConcurrentHashMap`)
- **`ReferenceEntry<K, V>`** → Đại diện cho một phần tử trong cache
- **`ValueReference<K, V>`** → Lưu trữ dữ liệu thực tế

💡 **Mục tiêu:** Giảm tranh chấp giữa các thread bằng cách chia nhỏ thành nhiều `Segment`.

```java
┌─────────────────── LocalCache<K, V> ───────────────────┐
│                                                       │
│ ┌────── Segment<K, V> ────────┐  ┌────── Segment<K, V> ────────┐
│ │                              │  │                              │
│ │ ┌── ReferenceEntry<K, V> ─┐  │  │ ┌── ReferenceEntry<K, V> ─┐  │
│ │ │ key1 → ValueReference<V> │  │  │ key2 → ValueReference<V> │  │
│ │ └──────────────────────────┘  │  │ └──────────────────────────┘  │
│ │ ...                          │  │ ...                          │
│ └──────────────────────────────┘  └──────────────────────────────┘
└──────────────────────────────────────────────────────────┘
```
## **2️⃣ Flow Hoạt Động của `get(key)`**

📌 **Khi bạn gọi `cache.get(key)`, Guava Cache thực hiện các bước sau:**

### 🚀 **Luồng hoạt động của `get(key)`**

1. **Tính toán vị trí `Segment` dựa trên `key.hashCode()`**
2. **Tìm `ReferenceEntry` trong `Segment`**
    - Nếu **tìm thấy** & chưa **hết hạn** → **Trả về value**
    - Nếu **không tìm thấy hoặc đã hết hạn** → **Gọi `CacheLoader.load(key)` để lấy dữ liệu mới**
3. **Cập nhật lại cache** nếu có dữ liệu mới

```java
V get(K key) {
    int hash = hash(key);
    Segment<K, V> segment = segmentFor(hash);
    return segment.get(key, hash);
}
```
👉 **Tương tự như `ConcurrentHashMap`, nhưng có thêm cơ chế hết hạn (`expireAfterWrite`, `expireAfterAccess`)**

## **3️⃣ Flow Hoạt Động của `put(key, value)`**

📌 **Khi gọi `cache.put(key, value)`, Guava Cache thực hiện:**

1. **Xác định `Segment` dựa trên hash của key**
2. **Thêm hoặc cập nhật `ReferenceEntry<K, V>` trong `Segment`**
3. **Kiểm tra kích thước cache (`maximumSize`, `maximumWeight`)**
    - Nếu quá giới hạn → **Loại bỏ phần tử cũ nhất (LRU - Least Recently Used)**
4. **Nếu dùng `expireAfterWrite`, đặt timestamp hết hạn**

```java
void put(K key, V value) {
    int hash = hash(key);
    Segment<K, V> segment = segmentFor(hash);
    segment.put(key, hash, value);
}
```
💡 **Guava Cache hỗ trợ LRU bằng cách duy trì danh sách liên kết (Linked List) trong `Segment`**

## **4️⃣ Cơ chế loại bỏ phần tử (`Eviction`)**

📌 **Guava Cache sử dụng nhiều chính sách để xóa phần tử:**

|**Chính sách**|**Hoạt động**|
|---|---|
|`maximumSize(n)`|Lưu tối đa `n` phần tử, xóa phần tử cũ khi đạt giới hạn|
|`expireAfterWrite(t, unit)`|Xóa sau `t` thời gian từ khi ghi vào cache|
|`expireAfterAccess(t, unit)`|Xóa nếu không truy cập trong `t` thời gian|
|`weakKeys(), weakValues()`|Xóa key/value nếu không còn tham chiếu mạnh|
|`softValues()`|Xóa value khi JVM thiếu bộ nhớ|

👉 **Khi quá giới hạn, Guava Cache dùng thuật toán `LRU` trong mỗi `Segment` để xóa phần tử ít được sử dụng nhất.**

## **5️⃣ `refreshAfterWrite` Hoạt Động Như Thế Nào?**

📌 **`refreshAfterWrite` cho phép cập nhật dữ liệu mà không cần chờ `expireAfterWrite`**

### 🔥 **Flow hoạt động:**

1. **Người dùng truy cập key**
2. **Nếu key đã tồn tại & quá hạn `refreshAfterWrite`, một thread nền (`background refresh thread`) sẽ tải lại dữ liệu**
3. **Trong khi tải lại, cache vẫn trả về dữ liệu cũ**

👉 **Cơ chế này giúp cache luôn cập nhật mà không chặn request**
```java
LoadingCache<String, String> cache = CacheBuilder.newBuilder()
    .refreshAfterWrite(10, TimeUnit.MINUTES)
    .build(new CacheLoader<String, String>() {
        @Override
        public String load(String key) {
            return fetchDataFromDB(key);
        }
    });
```
## **6️⃣ Guava Cache Xử Lý `Concurrency` Như Thế Nào?**

📌 **Guava Cache sử dụng `Segment<K, V>` để chia nhỏ cache thành nhiều phần, giúp tăng hiệu suất đọc/ghi đồng thời**

### 🚀 **Các cơ chế đảm bảo `thread-safe` trong Guava Cache:**

1. **Dùng `ReentrantLock` trong `Segment` để kiểm soát ghi (`put`)**
2. **Dùng `volatile` để tránh vấn đề `visibility` giữa các thread**
3. **Dùng `AtomicReference` để đảm bảo cập nhật dữ liệu an toàn**

