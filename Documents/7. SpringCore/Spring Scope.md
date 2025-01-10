
---
Trong **Spring Framework**, **scope** xác định vòng đời và phạm vi của một bean, tức là khi nào bean được tạo ra, sống bao lâu, và có thể được sử dụng ở đâu. Spring cung cấp nhiều loại scope khác nhau để phù hợp với các tình huống sử dụng khác nhau.

### 1. **Các loại Spring Scope:**

#### a. **Singleton (Mặc định)**

- **Mô tả**: Một instance duy nhất của bean được tạo và dùng chung trong toàn bộ ứng dụng.
- **Vòng đời**: Bean được tạo khi container khởi động và tồn tại cho đến khi container bị dừng.
- **Sử dụng**: Đây là scope mặc định cho tất cả các bean trong Spring.

**Cách sử dụng:**
```java
@Component
public class MyBean {
    // Bean này sẽ được sử dụng như một singleton
}
```

#### b. **Prototype**

- **Mô tả**: Mỗi lần yêu cầu một bean, một instance mới sẽ được tạo.
- **Vòng đời**: Bean được tạo mới mỗi khi có yêu cầu và không được container quản lý sau khi được cung cấp.
- **Sử dụng**: Khi cần tạo nhiều instance khác nhau, như trong các tác vụ không liên quan hoặc trạng thái không cần chia sẻ.

**Cách sử dụng:**
```java
@Component
@Scope("prototype")
public class MyBean {
    // Mỗi lần yêu cầu sẽ tạo ra một instance mới
}

```

#### c. **Request (Web Context)**

- **Mô tả**: Một instance của bean được tạo mới cho mỗi HTTP request.
- **Vòng đời**: Tồn tại trong suốt thời gian của request HTTP và bị hủy sau khi request kết thúc.
- **Sử dụng**: Dùng cho các bean cần dữ liệu cụ thể của một request.

**Cách sử dụng:**
```java
@Component
@Scope("request")
public class MyBean {
    // Bean này sẽ tồn tại trong suốt một request HTTP
}
```
#### d. **Session (Web Context)**

- **Mô tả**: Một instance của bean được tạo mới cho mỗi session HTTP và được dùng chung trong suốt thời gian của session đó.
- **Vòng đời**: Tồn tại trong suốt thời gian của session HTTP và bị hủy khi session kết thúc.
- **Sử dụng**: Dùng cho các bean cần giữ dữ liệu giữa các request trong cùng một session.

**Cách sử dụng:**
```java
@Component
@Scope("session")
public class MyBean {
    // Bean này sẽ tồn tại trong suốt một session HTTP
}
```
#### e. **Application (Web Context)**

- **Mô tả**: Một instance của bean được tạo và dùng chung trong toàn bộ vòng đời của ứng dụng web.
- **Vòng đời**: Tồn tại cho đến khi ứng dụng web bị dừng.
- **Sử dụng**: Dùng cho các bean cần dùng chung dữ liệu hoặc trạng thái trong toàn bộ ứng dụng.

**Cách sử dụng:**
```java
@Component
@Scope("application")
public class MyBean {
    // Bean này tồn tại trong toàn bộ vòng đời của ứng dụng
}
```

#### f. **WebSocket (Web Context)**

- **Mô tả**: Một instance của bean được tạo mới cho mỗi WebSocket session và tồn tại trong suốt thời gian của session đó.
- **Vòng đời**: Tồn tại trong suốt thời gian của một WebSocket session.
- **Sử dụng**: Dùng cho các ứng dụng sử dụng WebSocket.

**Cách sử dụng:**
```java
@Component
@Scope("websocket")
public class MyBean {
    // Bean này tồn tại trong suốt một WebSocket session
}
```
### **Tại sao cần Spring Scope:**

- **Quản lý hiệu quả tài nguyên**: Tùy vào yêu cầu của ứng dụng, các scope khác nhau giúp sử dụng tài nguyên hiệu quả hơn.
- **Tối ưu hóa hiệu suất**: Bằng cách tạo ra bean mới chỉ khi cần thiết (prototype) hoặc dùng lại các bean (singleton).
- **Hỗ trợ các ngữ cảnh khác nhau**: Các scope như request, session, và application giúp quản lý các bean trong các ứng dụng web một cách hiệu quả, phù hợp với ngữ cảnh sử dụng.