
---

Jakarta EE (trước đây là Java EE) là một nền tảng mạnh mẽ để xây dựng các ứng dụng Java doanh nghiệp. Dưới đây là một chương trình chi tiết để giúp bạn nắm vững Jakarta EE từ cơ bản đến nâng cao.


## **1. Giới thiệu về Jakarta EE**

Jakarta EE là một tập hợp các đặc tả (specifications) để phát triển và triển khai các ứng dụng Java trên quy mô doanh nghiệp.

### **Tính năng chính**:

- **Hỗ trợ ứng dụng web** (JSP, Servlets, JSF).
- **Xử lý giao dịch và bảo mật**.
- **Kết nối cơ sở dữ liệu** (JPA, JDBC).
- **API cho RESTful và SOAP web services** (JAX-RS, JAX-WS).
- **Tích hợp với CDI** (Context and Dependency Injection).

### **Công cụ cần thiết**:

1. **JDK**: Từ JDK 11 trở lên.
2. **Server Jakarta EE**: WildFly, Payara, TomEE, Open Liberty.
3. **IDE**: IntelliJ IDEA, Eclipse, hoặc NetBeans.


## **2. Các thành phần cốt lõi trong Jakarta EE**

### **a. Servlet (Web Tier)**

Servlet là thành phần cơ bản để xử lý các yêu cầu HTTP.

#### **Ví dụ: Tạo một Servlet đơn giản**

1. **Tạo Servlet:**
```java
@WebServlet("/hello")
public class HelloServlet extends HttpServlet {
    @Override
    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws IOException {
        response.setContentType("text/html");
        response.getWriter().println("<h1>Hello, Jakarta EE!</h1>");
    }
}

```

2. **Triển khai trên server (e.g., WildFly).**

- Deploy WAR file vào thư mục `deployments`.

3. **Truy cập ứng dụng**: Mở trình duyệt và truy cập `http://localhost:8080/app-name/hello`.

### **b. JSP (JavaServer Pages)**

JSP là công nghệ hiển thị giao diện web trong Jakarta EE.

#### **Ví dụ: Trang JSP đơn giản**

```jsp
<%@ page language="java" contentType="text/html; charset=UTF-8" %>
<!DOCTYPE html>
<html>
<head>
    <title>Jakarta EE</title>
</head>
<body>
    <h1>Xin chào từ JSP!</h1>
</body>
</html>
```

### **c. JPA (Java Persistence API)**

JPA cung cấp cơ chế quản lý cơ sở dữ liệu một cách ORM (Object-Relational Mapping).

#### **Cấu hình JPA với persistence.xml**

```xml
<persistence xmlns="http://xmlns.jcp.org/xml/ns/persistence" version="2.1">
    <persistence-unit name="examplePU">
        <class>com.example.entity.User</class>
        <properties>
            <property name="jakarta.persistence.jdbc.url" value="jdbc:mysql://localhost:3306/demo"/>
            <property name="jakarta.persistence.jdbc.user" value="root"/>
            <property name="jakarta.persistence.jdbc.password" value="password"/>
            <property name="jakarta.persistence.jdbc.driver" value="com.mysql.cj.jdbc.Driver"/>
        </properties>
    </persistence-unit>
</persistence>

```

### **d. JAX-RS (RESTful Web Services)**

JAX-RS cho phép xây dựng các API REST dễ dàng.

#### **Ví dụ: API REST đơn giản**

1. **Tạo Resource:**
```java
@Path("/users")
@Produces(MediaType.APPLICATION_JSON)
@Consumes(MediaType.APPLICATION_JSON)
public class UserResource {

    @GET
    public String getUsers() {
        return "[{\"name\":\"John Doe\"}]";
    }

    @POST
    public Response createUser(String userJson) {
        return Response.status(Response.Status.CREATED).entity(userJson).build();
    }
}

```

2. **Cấu hình JAX-RS trong `web.xml`:**
```xml
<servlet>
    <servlet-name>JAX-RS</servlet-name>
    <servlet-class>jakarta.ws.rs.core.Application</servlet-class>
</servlet>
<servlet-mapping>
    <servlet-name>JAX-RS</servlet-name>
    <url-pattern>/api/*</url-pattern>
</servlet-mapping>

```


3. **Triển khai và truy cập API:**

- `GET http://localhost:8080/app-name/api/users`
- `POST http://localhost:8080/app-name/api/users`


### **e. CDI (Context and Dependency Injection)**

CDI giúp tiêm phụ thuộc và quản lý vòng đời bean.

#### **Ví dụ: Tiêm phụ thuộc với CDI**

1. **Bean đơn giản:**

```java
@ApplicationScoped
public class GreetingService {
    public String greet(String name) {
        return "Hello, " + name;
    }
}

```


2. **Sử dụng Bean:**

```java
@Path("/greet")
public class GreetingResource {

    @Inject
    private GreetingService greetingService;

    @GET
    @Path("/{name}")
    public String greet(@PathParam("name") String name) {
        return greetingService.greet(name);
    }
}
```

