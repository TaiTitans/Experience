
---
**Spring Factory** là một khái niệm liên quan đến việc tạo và quản lý các bean trong **Spring Framework**. Spring cung cấp khả năng tạo ra các bean thông qua nhiều cách khác nhau, bao gồm sử dụng **Factory Method** hoặc **Factory Bean**.
### 1. **Factory Method:**

Factory Method là một phương pháp để tạo ra các bean thông qua một phương thức tĩnh (static) hoặc không tĩnh trong một class. Spring có thể gọi các phương thức này để tạo và quản lý các bean.
#### a. **Static Factory Method:**

- Một phương thức tĩnh được gọi để tạo ra bean.
```java
public class CarFactory {
    public static Car createCar() {
        return new Car();
    }
}

@Configuration
public class AppConfig {
    @Bean
    public Car car() {
        return CarFactory.createCar();
    }
}
```

#### b. **Instance Factory Method:**

- Một phương thức không tĩnh được gọi trên một instance của class để tạo ra bean.
```java
public class CarFactory {
    public Car createCar() {
        return new Car();
    }
}

@Configuration
public class AppConfig {
    @Bean
    public CarFactory carFactory() {
        return new CarFactory();
    }

    @Bean
    public Car car() {
        return carFactory().createCar();
    }
}

```

### 2. **FactoryBean Interface:**

**`FactoryBean<T>`** là một interface đặc biệt trong Spring, cho phép bạn tùy chỉnh việc tạo ra các bean. Khi một lớp triển khai `FactoryBean`, Spring không trực tiếp tạo ra instance của lớp đó, mà thay vào đó gọi phương thức `getObject()` để tạo ra instance của bean.

#### a. **Cách sử dụng FactoryBean:**

- Triển khai `FactoryBean` để tùy chỉnh cách thức tạo bean.
```java
public class CarFactoryBean implements FactoryBean<Car> {

    @Override
    public Car getObject() throws Exception {
        return new Car();
    }

    @Override
    public Class<?> getObjectType() {
        return Car.class;
    }

    @Override
    public boolean isSingleton() {
        return true; // hoặc false tùy thuộc vào yêu cầu
    }
}

@Configuration
public class AppConfig {
    @Bean
    public CarFactoryBean carFactoryBean() {
        return new CarFactoryBean();
    }
}

```

#### b. **Ưu điểm của FactoryBean:**

- **Tùy chỉnh logic khởi tạo**: Cho phép bạn thực hiện các bước tùy chỉnh khi tạo ra các bean.
- **Quản lý bean phức tạp**: Hữu ích khi cần quản lý việc khởi tạo các đối tượng phức tạp hoặc bên thứ ba.

### 3. **Khi nào sử dụng Spring Factory:**

- **Tạo các đối tượng phức tạp**: Khi việc khởi tạo bean cần nhiều bước hoặc phức tạp.
- **Tích hợp với các thư viện bên ngoài**: Khi cần tạo bean cho các lớp từ thư viện bên ngoài mà bạn không kiểm soát.
- **Tạo các bean tùy chỉnh**: Khi cần tùy chỉnh việc khởi tạo hoặc cấu hình các bean dựa trên các điều kiện runtime.

Spring Factory là một phần quan trọng trong việc quản lý và khởi tạo các bean trong ứng dụng, giúp ứng dụng trở nên linh hoạt và dễ bảo trì hơn.