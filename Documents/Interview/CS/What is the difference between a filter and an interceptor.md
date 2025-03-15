
---

### **Sự khác biệt giữa Filter và Interceptor trong Java**

1. **Nguyên tắc triển khai khác nhau**
    
    - **Filter** được triển khai dựa trên **cơ chế callback function**, sử dụng phương thức `doFilter()` với tham số `FilterChain`, thực chất là một **giao diện callback**.
    - **Interceptor** sử dụng **cơ chế phản xạ (Reflection) của Java**, được triển khai dựa trên **cơ chế proxy động (Dynamic Proxy)**.
2. **Thời điểm kích hoạt khác nhau**
    
    - **Filter**: Được thực thi **ngay sau khi yêu cầu (request) vào container**, nhưng **trước khi vào Servlet**, và kết thúc **sau khi Servlet xử lý xong**.
    - **Interceptor**: Được thực thi **sau khi request vào Servlet**, nhưng **trước khi vào Controller**, và kết thúc **sau khi Controller render xong View**.
3. **Ứng dụng khác nhau**
    
    - **Interceptor** gần với tầng **business logic**, chủ yếu dùng để **xác thực quyền truy cập, ghi log, kiểm tra điều kiện nghiệp vụ**, v.v.
    - **Filter** thường được dùng để **xử lý các chức năng chung**, như **lọc từ ngữ nhạy cảm, nén dữ liệu phản hồi, xử lý encoding**, v.v.
4. **Phạm vi áp dụng khác nhau**  
    **Trình tự xử lý request trong hệ thống**:
```java
Request -> Container -> Filter -> Servlet -> Interceptor -> Controller -> View
```

. Như vậy:
    
    - **Filter được thực thi trước** Interceptor.
    - **Interceptor được thực thi sau Servlet nhưng trước Controller**.

---

### **Khi nào nên dùng Filter và khi nào nên dùng Interceptor?**

✅ **Dùng Filter** khi cần xử lý **các tác vụ chung** ở tầng thấp như:

- Xử lý **mã hóa request/response** (Encoding).
- **Lọc nội dung** request trước khi đến Servlet.
- **Ghi log hoặc theo dõi request** ở tầng thấp.

✅ **Dùng Interceptor** khi cần can thiệp vào **luồng xử lý của ứng dụng** như:

- **Xác thực quyền truy cập (Authorization)**.
- **Ghi log các thao tác của người dùng**.
- **Thêm hoặc chỉnh sửa dữ liệu request/response trước khi đến Controller**.

➡ **Tóm lại**, Filter hoạt động ở tầng thấp hơn, còn Interceptor được sử dụng nhiều hơn ở tầng business logic.