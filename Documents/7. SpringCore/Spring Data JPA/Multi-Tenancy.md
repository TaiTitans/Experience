
---
Multi-Tenancy là một kiến trúc quan trọng giúp **một ứng dụng có thể phục vụ nhiều khách hàng (tenant) cùng lúc** mà vẫn đảm bảo dữ liệu của từng tenant được cô lập.



### **Multi-Tenancy là gì?**

Multi-Tenancy (đa khách thuê) là một **kiến trúc phần mềm** cho phép **một ứng dụng duy nhất phục vụ nhiều khách hàng (tenants) riêng biệt**.

Mỗi **tenant** có thể là:

- **Công ty, tổ chức khác nhau** sử dụng chung một hệ thống SaaS.
- **Người dùng cá nhân**, mỗi người có dữ liệu riêng biệt.
- **Cửa hàng, doanh nghiệp nhỏ** trên một nền tảng thương mại điện tử.

Ví dụ:

- **Gmail** là một hệ thống multi-tenant, nơi mỗi công ty có email riêng (@company.com).
- **Shopify, Shopee** phục vụ nhiều cửa hàng khác nhau trên cùng một nền tảng.
- **Salesforce, AWS** cung cấp dịch vụ SaaS cho nhiều công ty mà không cần mỗi công ty có server riêng.
---
### **Khi nào nên sử dụng Multi-Tenancy?**

Bạn nên sử dụng Multi-Tenancy khi:

🔹 **1. Cần phục vụ nhiều khách hàng trên cùng một hệ thống**

- Nếu bạn phát triển một nền tảng SaaS, marketplace, hoặc hệ thống dùng chung cho nhiều công ty/cửa hàng.

🔹 **2. Dữ liệu của mỗi tenant cần được phân tách độc lập**

- Mỗi công ty, tổ chức có dữ liệu riêng và không muốn bị lẫn với dữ liệu của khách hàng khác.

🔹 **3. Muốn tối ưu tài nguyên & giảm chi phí hạ tầng**

- Thay vì tạo **một ứng dụng riêng cho từng khách hàng**, bạn có thể dùng một hệ thống duy nhất và chỉ phân tách dữ liệu.

🔹 **4. Yêu cầu bảo mật, truy cập & kiểm soát dữ liệu riêng biệt**

- Mỗi tenant có thể có quyền riêng, chỉ có thể xem dữ liệu của họ.

🔹 **5. Muốn dễ dàng mở rộng (scalability)**

- Multi-Tenancy giúp dễ dàng **thêm khách hàng mới mà không cần triển khai lại ứng dụng**.
## **1. Các Mô Hình Multi-Tenancy trong JPA**

Có **3 cách chính** để triển khai Multi-Tenancy trong Spring Boot & JPA:

|Mô hình|Mô tả|
|---|---|
|**DATABASE**|Mỗi tenant có một database riêng biệt.|
|**SCHEMA**|Mỗi tenant có một schema riêng trong cùng một database.|
|**DISCRIMINATOR**|Tất cả tenants dùng chung DB nhưng có cột `tenant_id` để phân biệt.|

Mỗi mô hình có ưu và nhược điểm riêng:

| Mô hình           | Tính cô lập dữ liệu | Dễ triển khai | Hiệu suất   |
| ----------------- | ------------------- | ------------- | ----------- |
| **DATABASE**      | ✅ Rất cao           | ❌ Khó         | ⚠️ Chậm hơn |
| **SCHEMA**        | ✅ Cao               | ⚖️ Trung bình | ✅ Tốt       |
| **DISCRIMINATOR** | ❌ Thấp              | ✅ Dễ          | ✅ Rất tốt   |
## **2. Multi-Tenancy với JPA**

### **🔹 1. Multi-Tenancy Level: DATABASE**

Mỗi tenant có một database riêng, Spring Boot sẽ chọn **database phù hợp cho từng request**.

🔹 **Khi nào nên dùng?**

- Khi cần **cô lập dữ liệu hoàn toàn** giữa các tenants.
- Khi mỗi tenant có lượng dữ liệu lớn, tránh ảnh hưởng hiệu suất.
- Khi ứng dụng có thể **quản lý nhiều database connections**.

### **🔹 2. Multi-Tenancy Level: SCHEMA**

Mỗi tenant có một **schema riêng trong cùng một database**.

🔹 **Khi nào nên dùng?**

- Khi muốn **cô lập dữ liệu nhưng vẫn dùng chung database**.
- Khi hệ thống cần **quản lý kết nối hiệu quả hơn so với DATABASE**.

### **🔹 3. Multi-Tenancy Level: DISCRIMINATOR**

Dùng **một database chung**, mỗi bảng có thêm cột `tenant_id` để phân biệt tenants.

🔹 **Khi nào nên dùng?**

- Khi tenants có **ít dữ liệu** hoặc **cần truy vấn chéo tenants**.
- Khi muốn **triển khai nhanh** mà không phải tạo nhiều schema/database.

## **Khi nào nên dùng Hibernate Multi-Tenancy Provider?**

Hibernate Multi-Tenancy Provider giúp tự động quản lý kết nối cho từng tenant. Nó hữu ích khi:

1. **Cần thay đổi kết nối dựa trên tenant một cách tự động.**
2. **Muốn tận dụng cơ chế Multi-Tenancy sẵn có của Hibernate thay vì tự quản lý.**
3. **Hệ thống có nhiều tenants cần kết nối nhanh chóng mà không phải quản lý từng connection riêng.**

🔹 **Nhược điểm**:

- Cần cấu hình phức tạp.
- Chỉ thực sự cần thiết với mô hình `DATABASE` hoặc `SCHEMA`.

