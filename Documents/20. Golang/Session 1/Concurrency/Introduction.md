
---
Golang có hỗ trợ mạnh mẽ về concurrency (đồng thời) với **goroutines** và **channel**, giúp tận dụng tối đa tài nguyên CPU.

## 1️⃣ **Goroutines: Chạy hàm đồng thời với `go func()`**

🔹 **Goroutine** là một lightweight thread được quản lý bởi Golang runtime.  
🔹 Để chạy một hàm dưới dạng goroutine, chỉ cần thêm từ khóa `go` trước khi gọi hàm.

### 🔥 **Ví dụ cơ bản về Goroutine**
```go
package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine!")
}

func main() {
	go sayHello() // Chạy hàm này dưới dạng Goroutine

	fmt.Println("Hello from main function!")

	// Đợi một chút để goroutine có thời gian chạy
	time.Sleep(1 * time.Second)
}
```

📌 **Lưu ý**: Nếu `main()` kết thúc trước khi goroutine hoàn thành, goroutine cũng sẽ bị dừng.

## 2️⃣ **Channel: Giao tiếp giữa các Goroutines**

🔹 **Channel** giúp các goroutine trao đổi dữ liệu một cách an toàn.  
🔹 Có 2 loại channel:

- **Buffered Channel** (`make(chan int, size)`) có thể lưu trữ giá trị.
- **Unbuffered Channel** (`make(chan int)`) yêu cầu gửi và nhận đồng thời.

### 🔥 **Ví dụ sử dụng Channel**
```go
package main

import "fmt"

func worker(ch chan string) {
	ch <- "Data from worker" // Gửi dữ liệu vào channel
}

func main() {
	ch := make(chan string) // Tạo channel

	go worker(ch)           // Chạy goroutine
	msg := <-ch             // Nhận dữ liệu từ channel

	fmt.Println(msg) // Output: Data from worker
}
```

### **`select {}`: Lắng nghe nhiều channel cùng lúc**

🔹 **`select`** giúp xử lý nhiều channel đồng thời, tương tự như `switch-case` nhưng dành cho channels.
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}
```

📌 **Lưu ý**: `select` sẽ chọn bất kỳ channel nào sẵn sàng đầu tiên.

## 3️⃣ **Đồng bộ hóa với `sync.WaitGroup`, `sync.Mutex`**

🔹 **`sync.WaitGroup`** giúp đợi tất cả goroutines hoàn thành.  
🔹 **`sync.Mutex`** giúp tránh **race condition** khi nhiều goroutines truy cập dữ liệu chung.

### **🔥 Sử dụng `sync.WaitGroup` để đợi Goroutines hoàn thành**

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Giảm số lượng goroutine cần chờ
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Tăng số lượng goroutines cần chờ
		go worker(i, &wg)
	}

	wg.Wait() // Chờ tất cả goroutines hoàn thành
	fmt.Println("All workers finished")
}
```

📌 **Lưu ý**:

- **`wg.Add(n)`**: Đăng ký số lượng goroutines cần chờ.
- **`wg.Done()`**: Báo hiệu rằng một goroutine đã hoàn thành.
- **`wg.Wait()`**: Chờ tất cả goroutines kết thúc.

### **🔥 Sử dụng `sync.Mutex` để tránh Race Condition**

Nếu nhiều goroutines cùng đọc/ghi dữ liệu chung, có thể xảy ra **race condition**.
```go
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()   // Khóa dữ liệu
	defer c.mu.Unlock() // Mở khóa sau khi xong
	c.value++
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter.value) // Output: 10
}
```

📌 **Lưu ý**:

- **`mu.Lock()`**: Khóa dữ liệu để chỉ một goroutine truy cập.
- **`mu.Unlock()`**: Giải phóng khóa sau khi xong.

## 4️⃣ **Context API: Quản lý Goroutines**

🔹 **`context.WithCancel`**: Hủy goroutine khi không cần thiết.  
🔹 **`context.WithTimeout`**: Hủy goroutine sau một khoảng thời gian nhất định.


🔥 `context.WithCancel()` – Hủy Goroutines khi không cần nữa
```go
package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return
		default:
			fmt.Println("Worker running...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(2 * time.Second)
	cancel() // Gửi tín hiệu hủy goroutine
	time.Sleep(1 * time.Second)
}
```

📌 **Lưu ý**:

- **`ctx.Done()`**: Khi `cancel()` được gọi, `ctx.Done()` sẽ nhận giá trị, giúp goroutine thoát.

🔥 `context.WithTimeout()` – Tự động hủy sau thời gian nhất định

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker timeout, stopping...")
			return
		default:
			fmt.Println("Worker processing...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go worker(ctx)

	time.Sleep(3 * time.Second)
}
```

📌 **Lưu ý**:

- **`context.WithTimeout()`** sẽ tự động hủy sau **2 giây** mà không cần gọi `cancel()`.

## 🔥 **Tóm Tắt**

✅ **Goroutines** (`go func()`) giúp chạy đồng thời nhiều tác vụ.  
✅ **Channel** (`chan`, `select`) giúp goroutines giao tiếp với nhau.  
✅ **WaitGroup & Mutex** giúp đồng bộ hóa goroutines để tránh race condition.  
✅ **Context API** (`context.WithCancel`, `context.WithTimeout`) giúp quản lý thời gian sống của goroutines.



