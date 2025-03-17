
---
#### 1.2. Hiểu về GOPATH, GOROOT, module (go mod init, go mod tidy)

Go có cách tổ chức code khá độc đáo, bạn cần nắm mấy khái niệm này:

- **GOROOT**:
    - Là thư mục cài đặt Go (nơi chứa các file hệ thống của Go, như src, pkg, bin).
    - Bạn không cần đụng tới nó, chỉ cần biết mặc định là nơi bạn cài Go (ví dụ: /usr/local/go).
- **GOPATH**:
    - Là thư mục làm việc của bạn, nơi chứa code, dependency và file thực thi.
    - Mặc định từ Go 1.13, bạn không cần set GOPATH nữa (nó dùng ~/go nếu không chỉ định).
    - Trong GOPATH, có 3 thư mục chính:
        - src: Chứa source code.
        - pkg: Chứa các file đã compile.
        - bin: Chứa các file thực thi (sau khi go build hoặc go install).
- **Module (Go Modules)**:
    - Từ Go 1.11, Go giới thiệu module để quản lý dependency, thay thế dần GOPATH.
    - Tạo một module mới:
```bash
mkdir myproject
cd myproject
go mod init myproject
```
- Lệnh này tạo file go.mod, định nghĩa tên module và version.
- Thêm dependency (ví dụ: một package bên thứ 3):
```bash
go get github.com/gin-gonic/gin
```
Dọn dẹp và đồng bộ dependency:
```bash
go mod tidy
```

- - Lệnh này đảm bảo go.mod và go.sum khớp với code của bạn.

**Lưu ý**: Ngày nay, bạn nên dùng module thay vì dựa vào GOPATH. Nó hiện đại, linh hoạt và được cộng đồng ủng hộ mạnh mẽ.

---
#### Chạy chương trình đầu tiên: Hello, World!

Bây giờ, bạn sẽ viết và chạy chương trình Go đầu tiên.

- **Bước 1: Tạo file**
    - Tạo thư mục dự án:

```bash
mkdir hello
cd hello
```

Tạo file main.go với nội dung sau:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

**Bước 2: Chạy chương trình**

- Dùng lệnh:
```bash
go run main.go
```

- Kết quả: Hello, World! sẽ in ra màn hình.
- Nếu muốn build thành file thực thi:
```bash
go build main.go
```
Sau đó chạy file thực thi: ./main (Linux/macOS) hoặc main.exe (Windows).

**Giải thích code**:

- package main: Mọi chương trình Go đều thuộc một package. main là package đặc biệt để tạo ứng dụng thực thi.
- import "fmt": Import package fmt để in ra màn hình.
- func main(): Hàm chính, nơi chương trình bắt đầu chạy.


