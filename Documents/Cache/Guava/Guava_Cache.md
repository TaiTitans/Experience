
---
Guava Cache là một thư viện mạnh mẽ giúp quản lý bộ nhớ đệm (**in-memory caching**) trong Java. Nó tối ưu hơn so với `ConcurrentHashMap` vì hỗ trợ:  
✅ **Tự động hết hạn (TTL, idle timeout)**  
✅ **Giới hạn dung lượng bộ nhớ cache**  
✅ **Tự động tải lại dữ liệu khi cần thiết**  
✅ **Xử lý dữ liệu bị thiếu (cache miss) với `CacheLoader`**
## 🛠️ 1. Cách Khai Báo Guava Cache

### 📌 Cài đặt dependency

Nếu bạn dùng Maven, thêm vào `pom.xml`:
```java
<dependency>
    <groupId>com.google.guava</groupId>
    <artifactId>guava</artifactId>
    <version>33.0.0-jre</version>
</dependency>
```
## 🔥 2. Tạo Cache Đơn Giản

Guava Cache sử dụng `CacheBuilder` để tạo bộ nhớ cache với các quy tắc tuỳ chỉnh.

#### 📌 Ví dụ 1: Cache cơ bản
```java
import com.google.common.cache.Cache;
import com.google.common.cache.CacheBuilder;
import java.util.concurrent.TimeUnit;

public class GuavaCacheExample {
    public static void main(String[] args) {
        Cache<String, String> cache = CacheBuilder.newBuilder()
                .maximumSize(100) // Giới hạn 100 phần tử
                .expireAfterWrite(10, TimeUnit.MINUTES) // Hết hạn sau 10 phút
                .build();

        cache.put("username", "john_doe");
        System.out.println(cache.getIfPresent("username")); // john_doe

        cache.invalidate("username"); // Xoá phần tử khỏi cache
        System.out.println(cache.getIfPresent("username")); // null
    }
}
```
## ⚡ 3. Các Chính Sách Hết Hạn (Eviction Policies)

Guava Cache có nhiều cách để **quản lý bộ nhớ** và loại bỏ phần tử khỏi cache:

| **Chính sách**                          | **Mô tả**                                              |
| --------------------------------------- | ------------------------------------------------------ |
| `expireAfterWrite(duration, TimeUnit)`  | Xoá phần tử sau **X phút từ khi ghi**                  |
| `expireAfterAccess(duration, TimeUnit)` | Xoá phần tử sau **X phút không truy cập**              |
| `maximumSize(size)`                     | Xoá phần tử **cũ nhất** khi vượt quá kích thước tối đa |
| `maximumWeight(weight)`                 | Quản lý cache dựa trên trọng số của phần tử            |
| `weakKeys()`                            | Tự động xóa key nếu không còn tham chiếu mạnh          |
| `weakValues()`                          | Tự động xóa value nếu không còn tham chiếu mạnh        |
| `softValues()`                          | Xóa value khi bộ nhớ bị áp lực                         |
📌 Ví dụ 2: `expireAfterAccess`
```java
Cache<String, String> cache = CacheBuilder.newBuilder()
    .expireAfterAccess(5, TimeUnit.MINUTES) // Xoá nếu không được truy cập trong 5 phút
    .build();
```
📌 Ví dụ 3: `maximumSize`
```java
Cache<Integer, String> cache = CacheBuilder.newBuilder()
    .maximumSize(3) // Chỉ giữ tối đa 3 phần tử
    .build();

cache.put(1, "One");
cache.put(2, "Two");
cache.put(3, "Three");
cache.put(4, "Four"); // Lúc này "One" sẽ bị xóa

System.out.println(cache.getIfPresent(1)); // null (bị xóa)
```
## 🔄 4. Cache Loader: Tự Động Load Dữ Liệu

Thay vì **tự quản lý cache**, ta có thể dùng `CacheLoader` để **tự động lấy dữ liệu từ nguồn gốc nếu cache miss**.

#### 📌 Ví dụ 4: Dùng `CacheLoader`
```java
import com.google.common.cache.*;

import java.util.concurrent.*;

public class CacheLoaderExample {
    public static void main(String[] args) throws ExecutionException {
        LoadingCache<Integer, String> cache = CacheBuilder.newBuilder()
                .maximumSize(3)
                .expireAfterWrite(5, TimeUnit.MINUTES)
                .build(new CacheLoader<Integer, String>() {
                    @Override
                    public String load(Integer key) {
                        return "Value for " + key;
                    }
                });

        System.out.println(cache.get(1)); // Cache miss → Gọi load(1) → "Value for 1"
        System.out.println(cache.get(2)); // Cache miss → Gọi load(2) → "Value for 2"
    }
}
```
✅ **Nếu key chưa có trong cache, Guava sẽ tự gọi `load(key)` để lấy dữ liệu**.
## 🚀 5. Refresh Cache (Tự Động Tải Lại Dữ Liệu)

Nếu dữ liệu có thể thay đổi thường xuyên, ta có thể dùng **`refreshAfterWrite`** để cache tự động cập nhật mà không cần chờ hết hạn.

#### 📌 Ví dụ 5: Refresh Cache Sau 10 Phút
```java
LoadingCache<String, String> cache = CacheBuilder.newBuilder()
    .refreshAfterWrite(10, TimeUnit.MINUTES) // Tự động tải lại sau 10 phút
    .build(new CacheLoader<String, String>() {
        @Override
        public String load(String key) {
            return "New Data for " + key; // Gọi API / DB để lấy dữ liệu mới
        }
    });
```
✅ **Dữ liệu sẽ được tự động cập nhật sau mỗi 10 phút mà không cần chờ expire**.

## 🎯 6. Ứng Dụng Thực Tế

### 📌 Ví dụ 6: Dùng Cache để Lưu Kết Quả Từ Database
```java
import com.google.common.cache.*;
import java.util.concurrent.*;

public class DatabaseCacheExample {
    private static final LoadingCache<Integer, String> userCache = CacheBuilder.newBuilder()
            .maximumSize(100)
            .expireAfterWrite(10, TimeUnit.MINUTES)
            .build(new CacheLoader<Integer, String>() {
                @Override
                public String load(Integer userId) {
                    return getUserFromDatabase(userId);
                }
            });

    public static void main(String[] args) throws ExecutionException {
        System.out.println(userCache.get(1)); // Cache miss -> Lấy từ DB
        System.out.println(userCache.get(1)); // Cache hit -> Lấy từ cache
    }

    private static String getUserFromDatabase(Integer userId) {
        System.out.println("Fetching from database...");
        return "User" + userId;
    }
}
```
✅ **Chỉ gọi database khi cache miss → Giảm tải DB**
## 📌 7. Khi Nào Nên & Không Nên Dùng Guava Cache?

|✅ **Nên Dùng Khi**|❌ **Không Nên Dùng Khi**|
|---|---|
|Caching **trong bộ nhớ (RAM)**|Dữ liệu **quá lớn**, cần lưu trên disk hoặc Redis|
|Dữ liệu **thay đổi ít**|Cần **chia sẻ cache giữa nhiều server**|
|Dữ liệu cần **truy cập nhanh**|Cần cache **bền vững (Persistent)**|
|Cache chỉ dùng trong **ứng dụng đơn lẻ**|Cần **replication / phân tán**|

Nếu cần cache **đa máy chủ**, nên dùng **Redis, Memcached** thay vì Guava Cache.