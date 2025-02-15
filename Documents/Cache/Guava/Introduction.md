
---
Guava là một thư viện mã nguồn mở của Google giúp mở rộng và cải thiện các API Java chuẩn, cung cấp nhiều tiện ích mạnh mẽ để làm việc với collections, caching, strings, concurrency, I/O, v.v.

Tôi sẽ hướng dẫn bạn theo cách của một **Senior Java Developer**, tập trung vào:

1. **Kiến thức nền tảng**: Tại sao cần dùng Guava? Guava giải quyết vấn đề gì trong Java?
2. **Các module quan trọng**: Collections, Caching, Functional Programming, Concurrency, Strings, Files, I/O...
3. **Ứng dụng thực tế**: Làm thế nào để sử dụng Guava trong dự án thực tế?


----
## 🚀 1. Vì sao cần dùng Guava?

Java có sẵn Collections Framework và các API khác, vậy tại sao cần Guava?

- **Code ngắn gọn hơn**: Viết code Java "thuần" thường dài dòng, Guava giúp tinh gọn lại.
- **Tối ưu hiệu suất**: Guava có các cấu trúc dữ liệu tối ưu hơn `ArrayList`, `HashMap`,...
- **Hỗ trợ Functional Programming**: Java 8+ có Stream API, nhưng Guava có các cách tiếp cận khác như `FluentIterable`, `Function`,...
- **Cơ chế Caching mạnh mẽ**: `CacheBuilder` giúp quản lý cache tốt hơn so với `ConcurrentHashMap`.
- **Xử lý bất đồng bộ dễ dàng**: `ListenableFuture`, `Futures`, `RateLimiter` giúp cải thiện concurrency.
---
## 🔥 2. Các Module Quan Trọng

### 🛠️ 2.1. Guava Collections

Guava bổ sung nhiều cấu trúc dữ liệu mạnh hơn `java.util.Collection`:

- **Immutable Collections**: `ImmutableList`, `ImmutableMap`, `ImmutableSet`
- **Multimap**: Cho phép một key ánh xạ tới nhiều giá trị
- **BiMap**: Map 2 chiều (key -> value và value -> key)
- **Table**: Dữ liệu dạng bảng (giống `Map<Row, Column, Value>`)
- **ClassToInstanceMap**: Map để lưu các instance theo kiểu dữ liệu

🔹 Ví dụ
```java
ImmutableList<String> list = ImmutableList.of("A", "B", "C");
ImmutableMap<String, Integer> map = ImmutableMap.of("A", 1, "B", 2);

Multimap<String, String> multimap = ArrayListMultimap.create();
multimap.put("user1", "email1@gmail.com");
multimap.put("user1", "email2@gmail.com");
System.out.println(multimap); // {user1=[email1@gmail.com, email2@gmail.com]}
```

### 🚀 2.2. Guava Cache

Guava Cache mạnh hơn `ConcurrentHashMap` vì hỗ trợ:

- **Tự động loại bỏ dữ liệu theo thời gian** (`expireAfterWrite`, `expireAfterAccess`)
- **Giới hạn dung lượng** (`maximumSize`)
- **Tải dữ liệu khi cần thiết** (`CacheLoader`)

#### 🔹 Ví dụ
```java
Cache<String, String> cache = CacheBuilder.newBuilder()
    .maximumSize(100) // Giới hạn 100 phần tử
    .expireAfterWrite(10, TimeUnit.MINUTES) // Hết hạn sau 10 phút
    .build();

cache.put("username", "john_doe");
System.out.println(cache.getIfPresent("username")); // john_doe
```
### ⚡ 2.3. Guava Functional Programming

Guava giúp viết code functional hơn trước khi Java 8 ra đời.

#### 🔹 Ví dụ sử dụng `Function`
```java
Function<String, Integer> lengthFunction = new Function<String, Integer>() {
    @Override
    public Integer apply(String input) {
        return input.length();
    }
};
System.out.println(lengthFunction.apply("Guava")); // 5
```
Từ Java 8, bạn có thể dùng Lambda thay thế:
```java
Function<String, Integer> lengthFunction = String::length;
System.out.println(lengthFunction.apply("Guava")); // 5
```
### 🔄 2.4. Guava Concurrency

Guava bổ sung các công cụ mạnh để xử lý concurrency:

- `ListenableFuture`: Cho phép lắng nghe kết quả của một Future
- `Futures`: Hỗ trợ chaining các Future
- `RateLimiter`: Giới hạn tốc độ thực thi

#### 🔹 Ví dụ sử dụng `RateLimiter`
```java
RateLimiter limiter = RateLimiter.create(2.0); // 2 requests/second
for (int i = 0; i < 5; i++) {
    limiter.acquire(); // Chờ đến lượt
    System.out.println("Request " + i + " at " + System.currentTimeMillis());
}
```
### 📝 2.5. Guava Strings & Files

Guava có nhiều tiện ích mạnh để xử lý chuỗi và file.

#### 🔹 Xử lý chuỗi
```java
String result = Joiner.on(", ").skipNulls().join("Java", null, "Guava");
System.out.println(result); // Java, Guava
```
🔹 Đọc file nhanh chóng
```java
List<String> lines = Files.readLines(new File("data.txt"), Charsets.UTF_8);
```
## 🏗️ 3. Ứng Dụng Thực Tế

Guava thường được sử dụng trong:  
✅ Xử lý dữ liệu lớn với `ImmutableCollection`, `Multimap`, `Table`  
✅ Quản lý cache hiệu quả với `CacheBuilder`  
✅ Giới hạn tốc độ API với `RateLimiter`  
✅ Xử lý concurrency nâng cao với `ListenableFuture`

Ví dụ: Xây dựng API rate-limiting với `RateLimiter`
```java
@RestController
public class ApiController {
    private final RateLimiter limiter = RateLimiter.create(5.0); // 5 requests/second

    @GetMapping("/limited")
    public ResponseEntity<String> getLimitedResource() {
        if (!limiter.tryAcquire()) {
            return ResponseEntity.status(HttpStatus.TOO_MANY_REQUESTS).body("Rate limit exceeded");
        }
        return ResponseEntity.ok("Success");
    }
}
```
## 🎯 Tổng Kết

- Guava giúp code Java ngắn gọn, hiệu quả hơn
- Cung cấp các cấu trúc dữ liệu mạnh mẽ hơn Java Collections
- Quản lý cache linh hoạt với `CacheBuilder`
- Hỗ trợ functional programming & concurrency mạnh mẽ