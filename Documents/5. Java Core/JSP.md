
---

### **1. Tổng quan về JSP**

JSP là một công nghệ Java được sử dụng để tạo các trang web động trên server. Nó cho phép nhúng mã Java vào HTML bằng cách sử dụng các tag đặc biệt.

#### **Ưu điểm của JSP**:

- Dễ dàng tích hợp với Java code.
- Được biên dịch thành servlet, tối ưu hóa hiệu suất.
- Hỗ trợ phân tách logic ứng dụng (business logic) và giao diện người dùng (presentation).

#### **Kiến thức nền tảng cần có**:

- Hiểu biết về Java cơ bản.
- Biết về Servlet (JSP là phần mở rộng của Servlet).
- HTML, CSS, và cơ bản về web.

### **2. Cài đặt môi trường**

#### **Công cụ cần thiết**:

1. **JDK**: Đảm bảo bạn đã cài đặt JDK (Java Development Kit).
2. **IDE**: Sử dụng IntelliJ IDEA, Eclipse hoặc NetBeans.
3. **Server**: Cài đặt Apache Tomcat (một server hỗ trợ JSP).

#### **Thiết lập dự án JSP**:

- Tạo một dự án **Dynamic Web Project** trên Eclipse hoặc sử dụng Maven để cấu hình.
- Đảm bảo file `web.xml` (deployment descriptor) được cấu hình đúng.


### **3. Cấu trúc cơ bản của một file JSP**

Một file JSP có đuôi **`.jsp`** và bao gồm:

- HTML code.
- JSP tag để nhúng Java code.

#### **Ví dụ 1: Trang JSP đơn giản**

```jsp
<%@ page language="java" contentType="text/html; charset=UTF-8" pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html>
<head>
    <title>Hello JSP</title>
</head>
<body>
    <h1>Chào mừng bạn đến với JSP</h1>
    <% 
        String name = "Java Developer";
        out.println("Hello, " + name + "!");
    %>
</body>
</html>

```

#### **Giải thích**:

- `<%@ page ... %>`: Chỉ định các thuộc tính trang (ví dụ: encoding, content type).
- `<% ... %>`: Nhúng mã Java.


### **4. Các thành phần quan trọng trong JSP**

#### **a. Scriptlet**:

Dùng để chèn mã Java:

```jsp
<% 
    int a = 5;
    int b = 10;
    out.println("Tổng: " + (a + b));
%>

```

#### **b. Expression**:

In kết quả trực tiếp vào trang web:
```jsp
Tổng: <%= a + b %>
```

#### **c. Declaration**:

Khai báo biến hoặc phương thức:

```jsp
<%! 
    int multiply(int a, int b) {
        return a * b;
    }
%>
```

#### **d. Directives**:

Cung cấp meta-instructions cho container JSP.

**Page Directive**:
```jsp
<%@ page contentType="text/html; charset=UTF-8" %>
```

**Include Directive**:
```jsp
<%@ include file="header.jsp" %>
```

#### **e. JSTL (JSP Standard Tag Library)**:

Thay vì viết mã Java, bạn có thể sử dụng JSTL để xử lý logic:

```jsp
<%@ taglib uri="http://java.sun.com/jsp/jstl/core" prefix="c" %>
<c:if test="${user != null}">
    Chào, ${user.name}
</c:if>

```

### **5. Mô hình MVC với JSP**

JSP thường được sử dụng trong mô hình MVC (Model-View-Controller):

- **Model**: Xử lý logic nghiệp vụ (JavaBeans, DAO).
- **View**: JSP hiển thị giao diện người dùng.
- **Controller**: Servlet quản lý request/response.

#### **Ví dụ đơn giản**:

1. **Servlet (Controller):**

```java
@WebServlet("/hello")
public class HelloServlet extends HttpServlet {
    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        String message = "Chào mừng bạn đến với JSP!";
        request.setAttribute("msg", message);
        request.getRequestDispatcher("/hello.jsp").forward(request, response);
    }
}
```

2. **JSP (View):**

```jsp
<%@ page language="java" contentType="text/html; charset=UTF-8" %>
<!DOCTYPE html>
<html>
<head>
    <title>Trang Chào</title>
</head>
<body>
    <h1>${msg}</h1>
</body>
</html>
```

