
---
Tạo ra Master và Slave
Master: Nhiệm vụ write
Slave: Nhiệm vụ read
Các slave đồng bộ với nhau.

Nếu master chết -> Sẽ bầu 1 slave lên làm master. -> Sentinal giám sát master thông qua thuật toán vote.

P/s: Nếu master cũ sống lại thì vẫn trở thành slave