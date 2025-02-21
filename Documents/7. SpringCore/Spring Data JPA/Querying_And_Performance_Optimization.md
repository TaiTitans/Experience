
---
### **1. JPQL & Criteria API**

#### **Viết JPQL Queries hiệu quả**

- *_Tránh SELECT _:__ Chỉ chọn các cột cần thiết thay vì toàn bộ entity để giảm overhead.
- **Dùng `JOIN FETCH`:** Khi cần lấy dữ liệu liên quan, dùng `JOIN FETCH` thay vì `JOIN` để tránh N+1 Query Problem.
- **Dùng `COUNT(*)` riêng biệt:** Nếu cần phân trang, hãy dùng một query `COUNT(*)` riêng để tối ưu hiệu suất.

#### **Khi nào nên dùng Criteria API thay vì JPQL?**

- Khi truy vấn **động**, không cố định điều kiện tìm kiếm.
- Khi cần viết query theo cách **type-safe** thay vì string-based (giúp tránh lỗi cú pháp).
- Khi query phức tạp và cần kết hợp nhiều điều kiện linh hoạt.

### **2. Specifications & QueryDSL**

#### **Cách sử dụng `JpaSpecificationExecutor`**

- Dùng `Specification<T>` để xây dựng truy vấn động.
- Kết hợp `CriteriaBuilder` để tạo điều kiện linh hoạt.
- Ví dụ tìm sản phẩm theo tên hoặc giá:
```java
public static Specification<Product> hasName(String name) {
    return (root, query, criteriaBuilder) -> 
        criteriaBuilder.equal(root.get("name"), name);
}

public static Specification<Product> hasPriceGreaterThan(BigDecimal price) {
    return (root, query, criteriaBuilder) -> 
        criteriaBuilder.greaterThan(root.get("price"), price);
}
```

Sau đó kết hợp như sau:
```java
Specification<Product> spec = Specification.where(hasName("Laptop")).and(hasPriceGreaterThan(new BigDecimal(500)));
productRepository.findAll(spec);
```

#### **QueryDSL là gì? So sánh với JPQL & Criteria API**

- **QueryDSL** là API giúp viết query type-safe, dễ đọc, hỗ trợ truy vấn động mạnh mẽ hơn Criteria API.
- **Ưu điểm:**
    - Truy vấn ngắn gọn hơn Criteria API.
    - Có autocomplete trong IDE.
- **So sánh:**
    - **JPQL:** Đơn giản, dễ dùng nhưng không linh hoạt khi query động.
    - **Criteria API:** Linh hoạt nhưng verbose, khó đọc.
    - **QueryDSL:** Linh hoạt, mạnh mẽ, dễ đọc hơn Criteria API.


### **3. Pagination & Sorting**

#### **Tối ưu Pagination khi dữ liệu lớn (`Pageable`)**

- Dùng `PageRequest` với `Sort`:
```java
Page<Product> page = productRepository.findAll(PageRequest.of(0, 10, Sort.by("price").descending()));
```

- Tránh `ORDER BY RAND()` vì tốn tài nguyên.

#### **Offset Pagination vs Keyset Pagination**

- **Offset Pagination (`LIMIT OFFSET`)**:
    - Dễ dùng nhưng chậm khi offset lớn (phải duyệt qua nhiều dòng dữ liệu).
- **Keyset Pagination (`WHERE id > ? LIMIT ?`)**:
    - Nhanh hơn vì dùng index, chỉ lấy phần tử tiếp theo thay vì duyệt toàn bộ bảng.

👉 **Nên dùng Keyset Pagination khi xử lý danh sách lớn, như infinite scroll.**


### **4. Projection & DTO Mapping**

#### **Cách sử dụng `@Query` với JPQL & Native Queries để map kết quả vào DTO**

```java
@Query("SELECT new com.example.dto.ProductDTO(p.id, p.name, p.price) FROM Product p WHERE p.category.id = :categoryId")
List<ProductDTO> findProductsByCategory(@Param("categoryId") Long categoryId);
```

**Native Query** với DTO:
```java
@Query(value = "SELECT p.id, p.name, p.price FROM products p WHERE p.category_id = ?1", nativeQuery = true)
List<Object[]> findProductsByCategoryNative(Long categoryId);
```
- Sau đó map kết quả vào DTO.

#### **Khi nào dùng Interface Projection, Class Projection, hay Tuple?**

- **Interface Projection:** Dùng khi chỉ cần đọc dữ liệu đơn giản.
- **Class Projection:** Khi cần xử lý dữ liệu trước khi trả về.
- **Tuple:** Dùng khi cần query linh hoạt nhưng không muốn tạo DTO.
### **5. Cache & Performance Optimization**

#### **First-level Cache vs Second-level Cache**

- **First-level Cache:** Mặc định có trong Hibernate, cache theo **session**.
- **Second-level Cache:** Dùng **EhCache, Redis, Hazelcast** để cache entity giữa các session.
Sử dụng Redis Cache trong Spring Boot
```java
@Cacheable(value = "products", key = "#id")
public Product getProduct(Long id) {
    return productRepository.findById(id).orElse(null);
}
```

**Khi nào nên disable Hibernate auto-flush?**

- Khi batch update hoặc insert nhiều bản ghi để tránh flush từng entity một.


	### **6. Batch Processing & Bulk Operations**

#### **Cách batch insert, update, delete trong Spring Data JPA**

- **Batch Insert:**
- **Dùng JDBC batching thay vì JPA khi nào?**
    - Khi cần hiệu suất cao, vì JPA có overhead quản lý entity.
    - Khi insert/update hàng loạt bản ghi.


