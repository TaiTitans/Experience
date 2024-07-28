
---

1. **@Cacheable**:
    
    - Chú thích này được sử dụng để đánh dấu một phương thức cần được cache kết quả.
    - Khi một phương thức được đánh dấu bởi `@Cacheable`, Spring sẽ lưu kết quả trả về của phương thức này vào cache.
    - Khi gọi phương thức này lần sau, nếu kết quả đã có trong cache, Spring sẽ trả về kết quả từ cache thay vì gọi lại phương thức.
    - Ví dụ: `@Cacheable(value = "product", key = "#id")` sẽ lưu kết quả của phương thức theo id.
2. **@CacheEvict**:
    
    - Chú thích này được sử dụng để xóa dữ liệu trong cache.
    - Khi một phương thức được đánh dấu bởi `@CacheEvict`, Spring sẽ xóa các mục tương ứng trong cache.
    - Ví dụ: `@CacheEvict(value = "product", key = "#id")` sẽ xóa dữ liệu product theo id.
3. **@CachePut**:
    
    - Chú thích này được sử dụng để cập nhật dữ liệu trong cache.
    - Khi một phương thức được đánh dấu bởi `@CachePut`, Spring sẽ cập nhật các mục tương ứng trong cache.
    - Ví dụ: `@CachePut(value = "product", key = "#id")` sẽ cập nhật dữ liệu product theo id.
4. **@Caching**:
    
    - Chú thích này được sử dụng khi cần kết hợp các chú thích cache khác với nhau.
    - Ví dụ: `@Caching(cacheable = @Cacheable("product"), evict = { @CacheEvict("product"), @CacheEvict(cacheNames = "all-products", allEntries = true) })` sẽ áp dụng cả `@Cacheable` và `@CacheEvict` cho phương thức.

---
Example:

Giả sử bạn có một ứng dụng quản lý sản phẩm, và bạn muốn sử dụng caching để cải thiện hiệu suất khi truy xuất thông tin sản phẩm.

Ví dụ, bạn có một phương thức `getProductById()` trong `ProductService` như sau:

```Java
@Service
public class ProductService {
    private final ProductRepository productRepository;

    public ProductService(ProductRepository productRepository) {
        this.productRepository = productRepository;
    }

    @Cacheable(value = "products", key = "#id")
    public Product getProductById(Long id) {
        // Truy xuất sản phẩm từ database
        return productRepository.findById(id).orElse(null);
    }
}
```

Trong ví dụ này, chúng ta sử dụng `@Cacheable` để đánh dấu phương thức `getProductById()`. Khi phương thức này được gọi, Spring sẽ kiểm tra xem kết quả đã có trong cache chưa. Nếu có, nó sẽ trả về kết quả từ cache thay vì gọi lại phương thức.

Các thuộc tính của `@Cacheable` bao gồm:

- `value`: Chỉ định tên của cache được sử dụng.
- `key`: Chỉ định khóa (key) để lưu trữ kết quả trong cache. Trong ví dụ này, chúng ta sử dụng `#id` để lưu trữ kết quả theo id của sản phẩm.

Khi một request gọi đến `getProductById(1)`, Spring sẽ kiểm tra cache xem sản phẩm với id 1 đã có trong cache chưa. Nếu có, nó sẽ trả về kết quả từ cache. Nếu không, nó sẽ gọi `productRepository.findById(1)` để lấy sản phẩm từ database và lưu kết quả vào cache.

---
Khi sử dụng `@Cacheable` trong môi trường đa luồng, cần lưu ý một số điểm sau:

1. **Đồng bộ hóa (Synchronization)**: Trong môi trường đa luồng, có thể xảy ra tình huống hai luồng cùng gọi đến phương thức được đánh dấu bởi `@Cacheable`. Để tránh việc hai luồng cùng gọi lại phương thức và lưu kết quả vào cache, bạn cần đảm bảo rằng chỉ có một luồng được phép truy cập và lưu trữ kết quả vào cache.
    
     
    
    Cách đơn giản để đảm bảo điều này là sử dụng annotation `@CacheEvict` để xóa cache khi có sự thay đổi dữ liệu. Ví dụ:
```Java
@CacheEvict(cacheNames = "products", allEntries = true)
public void updateProduct(Product product) {
    // Cập nhật sản phẩm
}
```

1. **Cache Provider**: Lựa chọn cache provider phù hợp là rất quan trọng. Các cache provider phổ biến như Ehcache, Caffeine, Redis, v.v. có các đặc điểm và hành vi khác nhau. Cần đánh giá và lựa chọn cache provider phù hợp với yêu cầu của ứng dụng.
    
2. **Cấu hình Cache**: Cần cấu hình cache một cách hợp lý, chẳng hạn như kích thước cache, thời gian sống của mục cache, chính sách xóa cache, v.v. Cấu hình không phù hợp có thể dẫn đến các vấn đề như cache quá nhỏ, cache bị tràn, cache không được cập nhật kịp thời.
    
3. **Granularity của Cache**: Xác định mức độ chi tiết của cache là rất quan trọng. Nếu cache quá chi tiết (ví dụ: cache theo từng ID sản phẩm), có thể dẫn đến cache bị tràn. Ngược lại, nếu cache quá chung chung (ví dụ: cache toàn bộ danh sách sản phẩm), có thể dẫn đến cache không được sử dụng hiệu quả.
    
4. **Xử lý Fallback**: Khi có lỗi xảy ra với cache (ví dụ: cache không sẵn sàng), cần có cơ chế fallback để truy xuất dữ liệu từ nguồn dữ liệu chính (ví dụ: database).

---
1. **Tối ưu hóa Cấu hình Cache**:
    
    - **Kích thước Cache**: Xác định kích thước cache phù hợp dựa trên lượng dữ liệu và số lượng request. Nếu cache quá nhỏ, nó sẽ bị tràn nhanh chóng, dẫn đến nhiều miss cache.
    - **Thời gian sống (Time-to-Live - TTL)**: Thiết lập thời gian sống phù hợp cho các mục cache. Nếu TTL quá ngắn, cache sẽ bị xóa nhanh, dẫn đến nhiều miss cache. Nếu TTL quá dài, cache có thể chứa dữ liệu không còn phù hợp.
    - **Chính sách Xóa Cache**: Lựa chọn chính sách xóa cache phù hợp, ví dụ như Least Recently Used (LRU), Least Frequently Used (LFU), First In First Out (FIFO), v.v.
2. **Kiểm tra Hiệu suất Cache**:
    
    - **Tỷ lệ Hit/Miss Cache**: Theo dõi tỷ lệ hit/miss cache để đánh giá hiệu suất cache. Tỷ lệ hit cache càng cao, hiệu suất cache càng tốt.
    - **Thời gian Truy xuất**: Theo dõi thời gian truy xuất dữ liệu từ cache và từ nguồn dữ liệu chính (ví dụ: database). Nếu thời gian truy xuất từ cache nhanh hơn nhiều so với nguồn dữ liệu chính, cache đang hoạt động hiệu quả.
    - **Công cụ Monitoring**: Sử dụng các công cụ monitoring như Micrometer, Prometheus, Grafana để theo dõi chi tiết hoạt động của cache.
3. **Xử lý Fallback**:
    
    - **Khi Cache không sẵn sàng**: Khi có lỗi xảy ra với cache (ví dụ: cache không sẵn sàng), cần có cơ chế fallback để truy xuất dữ liệu từ nguồn dữ liệu chính (ví dụ: database).
    - **Sử dụng @CacheEvict và @CachePut**: Kết hợp `@CacheEvict` và `@CachePut` để đảm bảo dữ liệu trong cache luôn được cập nhật và không bị lỗi thời.
    - **Chuyển đổi dữ liệu**: Khi truy xuất dữ liệu từ nguồn dữ liệu chính, cần chuyển đổi dữ liệu về định dạng phù hợp với cache.
4. **Một số lỗi thường gặp khi sử dụng @Cacheable**:
    
    - **Không cấu hình cache provider**: Không cấu hình cache provider cho Spring Cache, dẫn đến cache không hoạt động.
    - **Cấu hình cache không phù hợp**: Cấu hình cache không phù hợp (ví dụ: kích thước cache quá nhỏ, TTL quá ngắn) dẫn đến hiệu suất cache kém.
    - **Sử dụng key không phù hợp**: Sử dụng key không phù hợp (ví dụ: sử dụng đối tượng phức tạp làm key) dẫn đến cache không hoạt động như mong đợi.
    - **Không xử lý Fallback**: Không có cơ chế Fallback khi cache không sẵn sàng, dẫn đến ứng dụng bị lỗi.


