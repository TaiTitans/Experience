
---
### **Các cách để dừng một luồng trong Java**

Có một số cách để dừng một luồng trong Java:

#### **1. Sử dụng phương thức `stop()` của luồng**

Phương thức `stop()` có thể được sử dụng để buộc một luồng dừng lại ngay lập tức. Tuy nhiên, phương thức này đã bị **loại bỏ (deprecated)** và không được khuyến nghị sử dụng.  
Lý do là vì khi sử dụng `stop()`, ngoại lệ `ThreadDeath` sẽ được lan truyền lên toàn bộ luồng, khiến luồng mục tiêu giải phóng **tất cả các khóa đối tượng** mà nó đang giữ. Điều này có thể dẫn đến tình trạng **không đồng bộ dữ liệu**, gây ra các lỗi khó lường trong chương trình.

#### **2. Sử dụng phương thức `interrupt()` để ngắt luồng**

Phương thức `interrupt()` chỉ đơn thuần gửi tín hiệu yêu cầu dừng luồng, nhưng **không đảm bảo luồng sẽ ngay lập tức dừng lại**. Khi gọi `interrupt()`, nó chỉ đánh dấu trạng thái của luồng là bị gián đoạn (`interrupted`), chứ không thực sự dừng luồng.  
Sau đó, chúng ta có thể gọi phương thức `Thread.currentThread().isInterrupted()` để kiểm tra xem luồng có bị gián đoạn hay không. Nếu phương thức này trả về `true`, ta có thể xử lý logic nghiệp vụ cần thiết, thường là **ném một ngoại lệ `InterruptedException`**, sau đó bắt ngoại lệ này bằng khối `try-catch` để xử lý.

#### **3. Sử dụng biến đánh dấu (Flag)**

Một cách tiếp cận an toàn hơn là sử dụng một **biến đánh dấu** (flag) để kiểm soát việc thoát khỏi vòng lặp của luồng. Khi giá trị của biến đạt đến một điều kiện nhất định, luồng sẽ tự động thoát ra một cách tự nhiên.  
Để đảm bảo tính nhất quán của biến dùng chung trong bộ nhớ, có thể sử dụng từ khóa `volatile`, giúp đảm bảo mỗi lần truy cập, biến luôn được lấy giá trị mới nhất từ bộ nhớ chính (main memory).

Tuy nhiên, phương pháp sử dụng `volatile` có một hạn chế lớn: **nó không thể xử lý tình huống khi luồng đang bị chặn (blocked)**. Ví dụ, nếu luồng đang ở trạng thái ngủ (`Thread.sleep()`), nó sẽ không kiểm tra lại biến cờ ngay lập tức, dẫn đến việc luồng không thể bị dừng đúng cách.

👉 **Vì vậy, cách tiếp cận chính xác nhất để dừng một luồng đang chạy là sử dụng `interrupt()` kết hợp với việc ném ngoại lệ một cách thủ công.**

---

📌 **Tóm tắt:**

- Không nên sử dụng `stop()`, vì nó có thể gây mất đồng bộ dữ liệu.
- `interrupt()` là cách tốt hơn nhưng chỉ gửi tín hiệu, không đảm bảo luồng dừng ngay lập tức.
- Sử dụng một biến `volatile` để đánh dấu dừng luồng là cách tiếp cận an toàn hơn nhưng không hoạt động nếu luồng bị chặn.
- **Giải pháp tối ưu:** Dùng `interrupt()` kết hợp với việc bắt và xử lý ngoại lệ `InterruptedException`. 