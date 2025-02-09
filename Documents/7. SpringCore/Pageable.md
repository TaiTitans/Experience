
---

`Pageable` trong JPA là một phần của Spring Data JPA được sử dụng để phân trang dữ liệu và trả về kết quả phân trang từ cơ sở dữ liệu. Nó giúp truy vấn dữ liệu trong các khối nhỏ, dễ quản lý hơn và cải thiện hiệu suất khi làm việc với lượng dữ liệu lớn.


### Các Thành Phần Chính

1. **Pageable Interface**:
    
    - Là giao diện đại diện cho thông tin phân trang (trang hiện tại, kích thước trang) và sắp xếp dữ liệu.
2. **Page Interface**:
    
    - Đại diện cho một trang kết quả, chứa danh sách dữ liệu của trang hiện tại cùng với các thông tin bổ sung như tổng số trang, tổng số phần tử, v.v.
3. **PageRequest Class**:
    
    - Là một triển khai của `Pageable`, cung cấp cách tạo đối tượng `Pageable` với thông tin về trang hiện tại và kích thước trang.

### Cách Sử Dụng Pageable

 1. Khai Báo Repository
```java
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.repository.PagingAndSortingRepository;

public interface ProductRepository extends JpaRepository<Product, Long>, PagingAndSortingRepository<Product, Long> {
}

```
2. Sử Dụng Pageable trong Phương Thức Repository
```java
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;

public interface ProductRepository extends JpaRepository<Product, Long> {
    Page<Product> findByCategory(String category, Pageable pageable);
}

```
3. Tạo Pageable Object
```java
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;

// Tạo Pageable với trang thứ 2, mỗi trang 10 phần tử, sắp xếp theo tên tăng dần
Pageable pageable = PageRequest.of(1, 10, Sort.by("name").ascending());
```
4. Sử Dụng Pageable trong Dịch Vụ
```java
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.stereotype.Service;

@Service
public class ProductService {

    @Autowired
    private ProductRepository productRepository;

    public Page<Product> getProductsByCategory(String category, Pageable pageable) {
        return productRepository.findByCategory(category, pageable);
    }
}

```
5. Sử Dụng Pageable trong Controller
```java
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class ProductController {

    @Autowired
    private ProductService productService;

    @GetMapping("/products")
    public Page<Product> getProducts(@RequestParam String category,
                                     @RequestParam int page,
                                     @RequestParam int size) {
        Pageable pageable = PageRequest.of(page, size);
        return productService.getProductsByCategory(category, pageable);
    }
}

```
### Lợi Ích của Pageable

1. **Quản lý Dữ liệu Tốt hơn**:
    
    - Truy vấn dữ liệu theo từng trang giúp giảm tải cho hệ thống và cải thiện hiệu suất.
2. **Tối Ưu Hóa Hiệu Suất**:
    
    - Truy vấn chỉ lấy dữ liệu cần thiết theo từng trang, giảm băng thông và thời gian xử lý.
3. **Đơn Giản Hóa Code**:
    
    - Tích hợp sẵn các phương thức và lớp hỗ trợ, giúp đơn giản hóa việc triển khai phân trang.

Với JPA `Pageable`, bạn có thể dễ dàng thêm phân trang vào các API của mình, giúp cải thiện trải nghiệm người dùng và hiệu suất hệ thống.

---
### 1. **Cách Hoạt Động Của Pageable**

- **Khi Gửi Yêu Cầu**: Khi bạn gọi phương thức với `Pageable`, Spring Data JPA sẽ tự động chuyển đổi đối tượng `Pageable` thành các câu lệnh SQL với `LIMIT` và `OFFSET`.

- **Truy Vấn SQL Tương Ứng**: Nếu bạn có truy vấn như sau:
```
productRepository.findByCategory("electronics", PageRequest.of(0, 10));
```
Spring Data JPA sẽ tạo câu lệnh SQL như sau:
```
SELECT * FROM products WHERE category = 'electronics' LIMIT 10 OFFSET 0;
```
### 2. **Pageable Chi Tiết**

#### `PageRequest.of(int page, int size)`

- **`page`**: Số thứ tự của trang bắt đầu từ `0` (trang đầu tiên là `0`).
- **`size`**: Số phần tử trong mỗi trang.
#### `Sort` trong Pageable

- Sắp xếp có thể được thêm vào `Pageable` bằng cách sử dụng `Sort`.
```
Pageable pageable = PageRequest.of(0, 10, Sort.by("name").ascending());
```
### 3. **Interface `Page` và `Slice`**

- **Page Interface**:
    
    - Cung cấp thông tin về trang hiện tại, tổng số phần tử, tổng số trang, và danh sách dữ liệu.
    - Phương thức thường dùng:
        - `getTotalPages()`: Tổng số trang.
        - `getTotalElements()`: Tổng số phần tử.
        - `getContent()`: Danh sách phần tử của trang hiện tại.
- **Slice Interface**:
    
    - Tương tự `Page`, nhưng không cung cấp tổng số phần tử và tổng số trang, giúp giảm chi phí tính toán nếu không cần thông tin này.
    - Phù hợp khi bạn chỉ cần biết có trang tiếp theo hay không (`hasNext()`).
### 4. **Custom Pageable**

Bạn có thể tùy chỉnh các yêu cầu phân trang và sắp xếp theo nhu cầu cụ thể.
*Sắp Xếp Nhiều Trường*
```
Pageable pageable = PageRequest.of(0, 10, Sort.by("name").ascending().and(Sort.by("price").descending()));
```
*Phân Trang Động*
Có thể tạo Pageable động dựa trên đầu vào từ người dùng:
```java
public Page<Product> getProducts(String category, int page, int size, String sortField, String sortDirection) {
    Sort sort = sortDirection.equalsIgnoreCase("asc") ? Sort.by(sortField).ascending() : Sort.by(sortField).descending();
    Pageable pageable = PageRequest.of(page, size, sort);
    return productRepository.findByCategory(category, pageable);
}
```
### 5. **Tối Ưu Hóa Phân Trang**

#### Kết Hợp với `@Query` và `@Modifying`

Bạn có thể sử dụng `@Query` để viết truy vấn tùy chỉnh có hỗ trợ phân trang:
```java
@Query("SELECT p FROM Product p WHERE p.category = :category")
Page<Product> findProductsByCategory(@Param("category") String category, Pageable pageable);
```
#### Native Query

Spring Data JPA cũng hỗ trợ native query với phân trang:
```java
@Query(value = "SELECT * FROM products WHERE category = :category", 
       countQuery = "SELECT count(*) FROM products WHERE category = :category", 
       nativeQuery = true)
Page<Product> findProductsByCategoryNative(@Param("category") String category, Pageable pageable);
```
#### Nâng Cao Hiệu Suất

- **Phân trang qua chỉ số (Index Pagination)**: Nếu cơ sở dữ liệu lớn, sử dụng `OFFSET` có thể không hiệu quả vì phải quét toàn bộ dữ liệu trước khi trả về kết quả. Bạn có thể sử dụng "Keyset Pagination" bằng cách giữ thông tin khóa cho phân trang tiếp theo.
### **Các Trường Hợp Sử Dụng Thực Tế**

- **Trang Quản Lý**: Hiển thị danh sách sản phẩm, người dùng, đơn hàng với phân trang.
- **API Công Khai**: Giới hạn số lượng dữ liệu trả về cho mỗi yêu cầu API, giúp giảm tải server và băng thông.
- **Tìm Kiếm**: Kết hợp với `Elasticsearch` hoặc `MongoDB` để phân trang kết quả tìm kiếm nhanh chóng.