

---

EntityManager là giao diện cốt lõi của JPA, nó cung cấp các phương thức để thao tác với các thực thể (entities) trong cơ sở dữ liệu. Một số nhiệm vụ chính của EntityManager bao gồm:

1. Quản lý vòng đời của các thực thể:
    
    - Lưu trữ (persist), cập nhật (merge), xóa (remove) các thực thể.
    - Tìm kiếm (find) và truy vấn (query) các thực thể.
2. Quản lý giao dịch (transaction):
    
    - Bắt đầu, xác nhận (commit) hoặc hủy bỏ (rollback) một giao dịch.
3. Quản lý cache:
    
    - Theo dõi và quản lý trạng thái của các thực thể trong bộ nhớ cache.

Khi sử dụng @PersistenceContext, một thể hiện của EntityManager sẽ được tự động tiêm vào trường hoặc phương thức được đánh dấu bởi chú thích này. Điều này giúp các lớp của bạn có thể truy cập và thao tác với các thực thể mà không cần phải tạo ra một thể hiện của EntityManager riêng.