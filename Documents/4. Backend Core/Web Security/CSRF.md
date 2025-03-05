
---
**CSRF** (Cross-Site Request Forgery) là một loại tấn công bảo mật mà kẻ tấn công lợi dụng quyền xác thực của người dùng trên một trang web để thực hiện các hành động trái phép mà người dùng không hề biết hoặc đồng ý. Đây là một mối đe dọa phổ biến đối với các ứng dụng web sử dụng cookies để xác thực.


### Cách Hoạt Động của CSRF

1. **Người dùng đăng nhập** vào một trang web hợp lệ (ví dụ: `bank.com`) và nhận được một session cookie.
2. Khi người dùng truy cập vào một trang web khác do kẻ tấn công kiểm soát, trang web này có thể chứa mã độc, ví dụ:
    
    html
    
    Copy code
    
    `<img src="http://bank.com/transfer?amount=1000&to=attacker-account">`
    
3. Trình duyệt sẽ tự động gửi cookie kèm theo yêu cầu này đến `bank.com` vì cookie vẫn còn hiệu lực.
4. Nếu không có cơ chế bảo vệ chống lại CSRF, server `bank.com` sẽ xử lý yêu cầu này như thể nó được người dùng hợp lệ thực hiện.

### Tại Sao CSRF Nguy Hiểm?

- Kẻ tấn công có thể thực hiện các hành động trái phép như chuyển tiền, thay đổi thông tin tài khoản mà người dùng không hề hay biết.
- Người dùng không cần thực hiện hành động nào mà chỉ cần truy cập vào một trang web độc hại.

### Cách Phòng Chống CSRF

1. **CSRF Token:**
    
    - Mỗi lần người dùng gửi một yêu cầu có khả năng thay đổi trạng thái (POST, PUT, DELETE), ứng dụng web gửi kèm một token duy nhất và ngẫu nhiên.
    - Token này được kiểm tra trên server để xác minh rằng yêu cầu đến từ người dùng hợp lệ và không phải từ một trang độc hại.
2. **SameSite Cookies:**
    
    - Cookie có thể được cấu hình với thuộc tính `SameSite` để hạn chế trình duyệt gửi cookie trong các yêu cầu cross-site.
        - `Strict`: Cookie chỉ được gửi khi người dùng truy cập cùng origin.
        - `Lax`: Cookie được gửi trong các yêu cầu "safe" như GET, nhưng không được gửi trong các yêu cầu gây thay đổi trạng thái từ một trang khác.
        - `None`: Cookie luôn được gửi trong các yêu cầu cross-site (chỉ nên dùng với HTTPS).
3. **Kiểm Tra Referrer Header:**
    
    - Server kiểm tra `Referer` hoặc `Origin` header để xác nhận rằng yêu cầu đến từ một nguồn đáng tin cậy.