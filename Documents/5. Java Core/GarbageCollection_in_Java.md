
---

Dưới đây là một số khái niệm cơ bản và cách Garbage Collection hoạt động trong Java:

1. **Đối tượng trong Java**: Trong Java, các đối tượng được tạo bằng từ khóa `new`. Mỗi khi tạo một đối tượng, Java cấp phát bộ nhớ cho đối tượng đó trong vùng Heap.
2. **Vùng nhớ Heap**: Đây là nơi lưu trữ các đối tượng Java. Các đối tượng trong Heap tồn tại cho đến khi không còn tham chiếu nào trỏ đến chúng.
3.  **Tham chiếu**: Một đối tượng được coi là "không sử dụng" khi không còn tham chiếu nào trỏ đến nó. Các loại tham chiếu trong Java bao gồm tham chiếu biến local, tham chiếu biến instance, tham chiếu biến static và tham chiếu từ một đối tượng đến đối tượng khác.
4. **Garbage Collection**: Là quá trình tự động thu gom các đối tượng không còn sử dụng để giải phóng bộ nhớ. Garbage Collector (GC) là một thành phần quan trọng của Java Runtime Environment (JRE) và hoạt động sau các bước sau:
    
    - **Mark**: GC đi qua từng đối tượng trong Heap và đánh dấu (mark) các đối tượng mà không có tham chiếu nào trỏ đến chúng.
        
    - **Sweep**: Sau khi đánh dấu, GC sẽ thu gom (sweep) các đối tượng đã được đánh dấu này và giải phóng bộ nhớ của chúng.
        
    - **Compact**: Trong một số thuật toán GC, có thể có bước Compact để tối ưu không gian bộ nhớ, dọn dẹp các khoảng trống và tạo ra một không gian bộ nhớ liên tục hơn.


5. **Thực hiện GC**: Trong Java, việc thực hiện Garbage Collection không do người lập trình quyết định mà do JVM quyết định. Tuy nhiên, có thể gợi ý JVM thực hiện GC bằng cách sử dụng các tùy chọn như `-XX:+UseG1GC` để sử dụng Garbage First (G1) Collector.

6. **Phương thức `finalize()`**: Trong Java, có phương thức `finalize()` mà bạn có thể ghi đè để thực hiện các hoạt động cuối cùng trước khi một đối tượng bị thu gom bởi GC. Tuy nhiên, việc sử dụng `finalize()` không được khuyến khích vì nó có thể gây ra các vấn đề về hiệu suất và không chắc chắn khi nào được gọi.

Khi lập trình trong Java, bạn không cần phải quản lý việc thu gom rác một cách trực tiếp. Tuy nhiên, để tối ưu hiệu suất của ứng dụng, bạn có thể thực hiện các biện pháp như tối ưu hóa thời gian sống của đối tượng, giảm số lượng tham chiếu không cần thiết và sử dụng các công cụ giám sát hiệu suất như JVisualVM để theo dõi hoạt động của Garbage Collection.