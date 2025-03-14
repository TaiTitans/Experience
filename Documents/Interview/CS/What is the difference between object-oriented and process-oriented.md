

---
**Hướng đối tượng và hướng thủ tục là các ý tưởng phát triển phần mềm.**

- **Hướng thủ tục** là quá trình phân tích các bước cần thiết để giải quyết vấn đề, sau đó sử dụng các hàm để thực hiện từng bước này và gọi chúng theo thứ tự khi sử dụng.
- **Hướng đối tượng** là quá trình phân tách các vấn đề thành từng đối tượng riêng lẻ, thiết kế chúng một cách độc lập, sau đó lắp ráp chúng lại thành một hệ thống hoạt động hoàn chỉnh. Hướng thủ tục chỉ được thực hiện bằng các hàm, trong khi hướng đối tượng sử dụng các lớp để triển khai các mô-đun chức năng riêng lẻ.

**Ví dụ với trò chơi cờ caro (backgammon):**

- **Ý tưởng thiết kế hướng thủ tục**: Trước tiên, phân tích vấn đề theo các bước:
    1. Bắt đầu trò chơi,
    2. Quân đen đi trước,
    3. Vẽ bàn cờ,
    4. Xác định thắng thua,
    5. Đến lượt quân trắng,
    6. Vẽ bàn cờ,
    7. Xác định thắng thua,
    8. Quay lại bước 2,
    9. Xuất kết quả cuối cùng.  
        Mỗi bước trên được giải quyết bằng một hàm riêng biệt để hoàn thành vấn đề.
- **Thiết kế hướng đối tượng**, ngược lại, tiếp cận vấn đề theo cách khác. Toàn bộ trò chơi cờ caro có thể được chia thành:
    - **Hai bên đen và trắng**: Chịu trách nhiệm nhận đầu vào từ người dùng và thông báo cho hệ thống bàn cờ rằng bố cục quân cờ đã thay đổi.
    - **Hệ thống bàn cờ**: Chịu trách nhiệm vẽ bàn cờ. Khi nhận được thông tin về sự thay đổi của quân cờ, hệ thống này hiển thị thay đổi đó trên màn hình.
    - **Hệ thống luật chơi**: Chịu trách nhiệm xác định các yếu tố như vi phạm luật, thắng thua, v.v.

Hai bên đen và trắng sẽ thông báo cho hệ thống bàn cờ về sự thay đổi, hệ thống bàn cờ hiển thị thay đổi đó, đồng thời sử dụng hệ thống luật chơi để đánh giá ván cờ.