
---
`ApplicationContext` là **Container cốt lõi** trong Spring Framework, quản lý **vòng đời của Bean**, **Dependency Injection (DI)** và các tính năng nâng cao như **Event Handling, Internationalization, và AOP**.

## **1. Vai Trò của ApplicationContext**

- **Quản lý Bean**: Tạo, cấu hình, và quản lý vòng đời Bean.
- **Hỗ trợ Dependency Injection (DI)**: Inject dependencies vào Bean.
- **Hỗ trợ AOP**: Áp dụng Aspect-Oriented Programming (AOP).
- **Hỗ trợ sự kiện (Event Handling)**: Publish và Listen các sự kiện trong Spring.
- **Hỗ trợ đa ngôn ngữ (Internationalization - i18n)**: Quản lý Resource Bundles.
- **Tích hợp với môi trường (Environment Abstraction)**: Cung cấp thông tin cấu hình.

## **2. Các Loại ApplicationContext**

Spring cung cấp nhiều loại `ApplicationContext`, mỗi loại phù hợp với một trường hợp sử dụng cụ thể.

### **2.1. ClassPathXmlApplicationContext**

- Đọc cấu hình từ file XML trong `classpath`.
- Phù hợp với ứng dụng legacy dùng XML.

```java
ApplicationContext context = new ClassPathXmlApplicationContext("applicationContext.xml");
MyBean myBean = context.getBean(MyBean.class);
```

### **2.2. FileSystemXmlApplicationContext**

- Tương tự `ClassPathXmlApplicationContext` nhưng lấy file XML từ đường dẫn hệ thống.
```java
ApplicationContext context = new FileSystemXmlApplicationContext("/path/to/applicationContext.xml");
```

### **2.3. AnnotationConfigApplicationContext**

- Dùng trong ứng dụng hiện đại với cấu hình bằng Annotation (`@Configuration`, `@ComponentScan`).
```java
ApplicationContext context = new AnnotationConfigApplicationContext(AppConfig.class);
MyBean myBean = context.getBean(MyBean.class);
```

```java
@Configuration
@ComponentScan("com.example")
public class AppConfig {
}
```


### **2.4. WebApplicationContext**

- Dùng trong ứng dụng Spring Web (Spring MVC, Spring Boot).
- Có thể được khởi tạo bởi `DispatcherServlet`.

```java
@WebServlet(name = "dispatcher", urlPatterns = "/")
public class MyServlet extends DispatcherServlet {
    public MyServlet() {
        super(new AnnotationConfigWebApplicationContext());
    }
}
```

## **3. Vòng Đời của ApplicationContext**

### **3.1. Quá trình khởi tạo ApplicationContext**

1. **Tạo đối tượng ApplicationContext** (XML hoặc Java-based).
2. **Quét và khởi tạo Bean** từ cấu hình (`@ComponentScan`, XML, v.v.).
3. **Inject Dependencies** vào Bean (`@Autowired`, Constructor Injection, Setter Injection).
4. **Kích hoạt các BeanPostProcessor** (`@PostConstruct`, `BeanFactoryPostProcessor`).
5. **Sẵn sàng phục vụ request**.


## **5. ApplicationContext vs BeanFactory**

|**Đặc điểm**|**ApplicationContext**|**BeanFactory**|
|---|---|---|
|**Tạo Bean**|Tạo **tất cả Bean** ngay từ đầu (Eager Initialization)|Tạo Bean khi cần (Lazy Initialization)|
|**Hỗ trợ AOP**|✅ Có|❌ Không|
|**Hỗ trợ sự kiện**|✅ Có|❌ Không|
|**Hỗ trợ i18n**|✅ Có|❌ Không|
|**Khi nào dùng?**|Dùng trong ứng dụng Spring chuẩn|Dùng khi cần **khởi tạo nhẹ** (Embedded Device, IoT, Mobile)|

💡 **Trong Spring Boot, ApplicationContext là mặc định** và hiếm khi dùng `BeanFactory` trực tiếp.


## **Event Handling trong ApplicationContext**

ApplicationContext hỗ trợ hệ thống **event-driven**, giúp Bean lắng nghe các sự kiện trong ứng dụng.

## **Inject ApplicationContext vào Bean**

Spring cho phép inject `ApplicationContext` vào một Bean để truy cập các Bean khác.


## **Kết Luận**

🔹 **ApplicationContext là Container quan trọng trong Spring, quản lý Bean và Dependency Injection.**  
🔹 **Có nhiều loại ApplicationContext: XML-based, Annotation-based, Web-based.**  
🔹 **Nó hỗ trợ các tính năng như Event Handling, AOP, Transaction, Lazy Initialization.**  
🔹 **ApplicationContext nên được dùng thay vì BeanFactory vì có nhiều tính năng mở rộng.**