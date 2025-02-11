
---
Trong Java SE, các bộ tạo số ngẫu nhiên được gọi chính xác hơn là **bộ tạo số ngẫu nhiên giả (pseudorandom number generators - PRNGs)**, vì chúng tạo ra một chuỗi số dựa trên thuật toán xác định.

[Oracle Docs](https://docs.oracle.com/en/java/javase/21/core/pseudorandom-number-generators.html?utm_source=chatgpt.com)

**Các giao diện và lớp quan trọng:**

- **`RandomGenerator`**: Cung cấp các phương thức để tạo ra các số ngẫu nhiên của các kiểu dữ liệu nguyên thủy khác nhau, như `nextInt()`, `nextLong()`, `nextDouble()`, và `nextBoolean()`.
    
    [Oracle Docs](https://docs.oracle.com/en/java/javase/21/docs/api/java.base/java/util/random/RandomGenerator.html?utm_source=chatgpt.com)
    
- **`RandomGeneratorFactory`**: Cho phép tạo các PRNG dựa trên các đặc điểm cụ thể, không chỉ dựa trên tên của thuật toán.
    
    [Oracle Docs](https://docs.oracle.com/en/java/javase/21/core/pseudorandom-number-generators.html?utm_source=chatgpt.com)
    

**Các loại PRNG trong Java:**

- **Legacy PRNGs**: Bao gồm các lớp như `Random`, `ThreadLocalRandom`, `SplittableRandom`, và `SecureRandom`. Trong đó, `Random` sử dụng thuật toán Linear Congruential Generator (LCG) với chu kỳ ngắn (2^48), nên được khuyến nghị chuyển sang các thuật toán mới hơn.
    
    [Oracle Docs](https://docs.oracle.com/en/java/javase/21/docs/api/java.base/java/util/random/package-summary.html?utm_source=chatgpt.com)
    
- **LXM PRNGs**: Là nhóm các thuật toán mới, kết hợp giữa Linear Congruential Generators (LCG) và Xor-based Generators (XBG), cung cấp chu kỳ dài hơn và chất lượng ngẫu nhiên tốt hơn.
    
    [Oracle Docs](https://docs.oracle.com/en/java/javase/21/docs/api/java.base/java/util/random/package-summary.html?utm_source=chatgpt.com)
    
- **Xoroshiro/Xoshiro PRNGs**: Nhóm thuật toán dựa trên phép toán XOR và các phép dịch vòng (rotate), cung cấp hiệu suất cao và chất lượng ngẫu nhiên tốt.
    
    [Oracle Docs](https://docs.oracle.com/en/java/javase/21/docs/api/java.base/java/util/random/package-summary.html?utm_source=chatgpt.com)
    

**Sử dụng PRNG trong ứng dụng đa luồng:**

Trong các ứng dụng đa luồng, việc sử dụng PRNG cần được xem xét cẩn thận để đảm bảo tính độc lập và hiệu suất. Các PRNG như `SplittableRandom` được thiết kế cho các tính toán song song, cho phép tạo ra các bộ tạo con với trạng thái độc lập, giúp cải thiện hiệu suất trong môi trường đa luồng.

[Oracle Docs](https://docs.oracle.com/en/java/javase/21/core/generating-pseudorandom-numbers-multithreaded-applications.html?utm_source=chatgpt.com)

**Lựa chọn thuật toán PRNG:**

Việc lựa chọn thuật toán PRNG phụ thuộc vào yêu cầu cụ thể của ứng dụng, bao gồm hiệu suất, chất lượng ngẫu nhiên, và tính bảo mật. Đối với các ứng dụng yêu cầu bảo mật cao, nên sử dụng `SecureRandom`. Đối với các ứng dụng yêu cầu hiệu suất cao và chu kỳ dài, các thuật toán trong nhóm LXM hoặc Xoroshiro/Xoshiro có thể là lựa chọn phù hợp.

[Oracle Docs](https://docs.oracle.com/en/java/javase/21/docs/api/java.base/java/util/random/package-summary.html?utm_source=chatgpt.com)

Để biết thêm chi tiết, bạn có thể tham khảo tài liệu chính thức về [Pseudorandom Number Generators](https://docs.oracle.com/en/java/javase/21/core/pseudorandom-number-generators.html).