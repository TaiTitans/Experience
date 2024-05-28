
----


Add library:

```
lettuce-core
spring-boot-starter-data-redis
```

_application.properties_

```
redis.host=localhost
redis.port=6379
```


Sử dụng `RedisTemplate` để tương tác với Redis hoặc sử dụng các annotation như `@Cacheable`, `@CachePut`, và `@CacheEvict` để áp dụng cơ chế caching.

```Java

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Service;

@Service
public class RedisService {

    @Autowired
    private RedisTemplate<String, Object> redisTemplate;

    public void saveData(String key, Object value) {
        redisTemplate.opsForValue().set(key, value);
    }

    public Object getData(String key) {
        return redisTemplate.opsForValue().get(key);
    }
}


```



Ví dụ sử dụng caching với annotation:

```Java
import org.springframework.cache.annotation.Cacheable;
import org.springframework.stereotype.Service;

@Service
public class UserService {

    @Cacheable(value = "users", key = "#userId")
    public User getUserById(String userId) {
        // Giả định lấy dữ liệu từ database
        return userRepository.findById(userId).orElse(null);
    }
}

```

**Kích hoạt caching trong cấu hình Spring Boot:**

```Java
import org.springframework.cache.annotation.EnableCaching;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@EnableCaching
public class Application {
    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }
}

```

**Spring Data** supported **Operations**:

1. `opsForValue()`: Kiểu Key-Value thông thường. Với Value là 1 giá trị String tùy ý.
2. `opsForHash()`: Tương ứng với cấu trúc Hash trong Redis. Value là một Object có cấu trúc
3. `opsForList()`: Tương ứng với cấu trúc List trong Redis. Value là một list.
4. `opsForSet()`: Tương ứng với cấu trúc Set trong Redis.
5. `opsForZSet()`: Tương ứng với cấu trúc ZSet trong Redis.