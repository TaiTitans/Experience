
---
-> 2 transaction chạy độc lập ko ảnh hưởng nhau.
-> Cho đến khi COMMIT

Sử dụng SAVEPOINT + ROLLBACK TO để rollback lại nếu transaction có 2 lệnh

---
Cơ chế cách ly transaction trong MYSQL

[10:22](https://www.youtube.com/watch?v=VWXZUWzT46U&t=622s) Read Uncommitted (RU) trong Transaction
[14:20](https://www.youtube.com/watch?v=VWXZUWzT46U&t=860s) Read Committed (RC) trong Transaction 
[17:13](https://www.youtube.com/watch?v=VWXZUWzT46U&t=1033s) Repeatable Read (RR) trong Transaction
[19:26](https://www.youtube.com/watch?v=VWXZUWzT46U&t=1166s) Serializable trong Transaction