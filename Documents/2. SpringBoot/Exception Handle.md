
---

### 1. **Spring Core Exceptions**

- **`IllegalArgumentException`**: Ném ra khi một phương thức nhận một tham số không hợp lệ.
- **`NullPointerException`**: Xảy ra khi cố gắng truy cập vào một đối tượng `null`.
- **`BeanCreationException`**: Xảy ra khi Spring không thể tạo một bean do lỗi cấu hình hoặc phụ thuộc không hợp lệ.
- **`BeanNotOfRequiredTypeException`**: Ném ra khi bean không phải là kiểu mong muốn.

### 2. **Spring Data Exceptions**

- **`DataAccessException`**: Ngoại lệ gốc cho các lỗi liên quan đến truy cập cơ sở dữ liệu.
- **`EmptyResultDataAccessException`**: Xảy ra khi không có kết quả nào được trả về từ truy vấn, nhưng truy vấn yêu cầu một đối tượng duy nhất.
- **`DataIntegrityViolationException`**: Ném ra khi có sự vi phạm ràng buộc dữ liệu (chẳng hạn như khóa chính hoặc khóa ngoại).
- **`OptimisticLockingFailureException`**: Xảy ra khi cập nhật thất bại do xung đột với phiên bản dữ liệu khác trong cơ sở dữ liệu (trong cơ chế khóa lạc quan).

### 3. **Spring MVC Exceptions**

- **`HttpRequestMethodNotSupportedException`**: Xảy ra khi một HTTP method không được hỗ trợ bởi một endpoint (ví dụ: gửi `POST` đến endpoint chỉ hỗ trợ `GET`).
- **`HttpMediaTypeNotSupportedException`**: Ném ra khi kiểu dữ liệu của request không được hỗ trợ (ví dụ: gửi JSON nhưng endpoint chỉ hỗ trợ XML).
- **`HttpMessageNotReadableException`**: Xảy ra khi Spring không thể đọc dữ liệu từ request (ví dụ: JSON không hợp lệ).
- **`MethodArgumentNotValidException`**: Xảy ra khi tham số của phương thức không hợp lệ (thường xảy ra khi validation không thành công).

### 4. **Security Exceptions (Spring Security)**

- **`AuthenticationException`**: Gốc của các ngoại lệ liên quan đến xác thực.
- **`BadCredentialsException`**: Xảy ra khi thông tin đăng nhập (username/password) không chính xác.
- **`AccessDeniedException`**: Ném ra khi người dùng không có quyền truy cập vào một tài nguyên.
- **`AccountExpiredException`**: Xảy ra khi tài khoản người dùng hết hạn.
- **`LockedException`**: Xảy ra khi tài khoản bị khóa.

### 5. **Transactional Exceptions**

- **`TransactionSystemException`**: Xảy ra khi có lỗi trong quá trình xử lý transaction.
- **`RollbackException`**: Ném ra khi transaction bị rollback.
- **`TransactionRequiredException`**: Xảy ra khi một hoạt động yêu cầu transaction nhưng transaction không tồn tại.

### 6. **File Handling Exceptions**

- **`MaxUploadSizeExceededException`**: Xảy ra khi tải lên tệp có kích thước vượt quá giới hạn cấu hình.
- **`MultipartException`**: Xảy ra khi có lỗi trong việc xử lý multipart file upload.

### 7. **Custom Exception (Ngoại lệ tùy chỉnh)**

- Spring Boot cho phép tạo các exception tùy chỉnh để xử lý các tình huống đặc biệt trong ứng dụng. Bằng cách sử dụng `@ControllerAdvice` kết hợp với `@ExceptionHandler`, bạn có thể quản lý các ngoại lệ tùy chỉnh và trả về các phản hồi phù hợp.

### 8. **Validation Exceptions**

- **`ConstraintViolationException`**: Ném ra khi một ràng buộc validation không thành công (thường xảy ra với các annotation như `@NotNull`, `@Size`, `@Min`, `@Max`).
- **`BindException`**: Xảy ra khi có lỗi trong việc bind dữ liệu từ request vào đối tượng (thường gặp trong trường hợp validation thất bại).