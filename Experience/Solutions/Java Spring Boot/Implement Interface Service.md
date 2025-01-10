
---
```java
public interface ProductService {
    Product getProductById(Long id);
    List<Product> getAllProducts();
    void saveProduct(Product product);
    void deleteProduct(Long id);
}
```

```java
import org.springframework.stereotype.Service;

@Service
public class ProductServiceImpl implements ProductService {

    @Override
    public Product getProductById(Long id) {
        // Logic để lấy thông tin sản phẩm
        return new Product(id, "Product Name");
    }

    @Override
    public List<Product> getAllProducts() {
        // Logic để lấy danh sách sản phẩm
        return List.of(new Product(1L, "Product 1"), new Product(2L, "Product 2"));
    }

    @Override
    public void saveProduct(Product product) {
        // Logic để lưu sản phẩm
        System.out.println("Product saved: " + product.getName());
    }

    @Override
    public void deleteProduct(Long id) {
        // Logic để xóa sản phẩm
        System.out.println("Product deleted: " + id);
    }
}

```

