
---

Redis supports 5 types of data types.

## Strings

Chuỗi Redis là một chuỗi byte. Các chuỗi trong Redis là chuỗi nhị phân an toàn, nghĩa là chúng có độ dài đã biết và không được xác định bởi bất kỳ ký tự kết thúc đặc biệt nào. Do đó, bạn có thể lưu trữ mọi thứ lên tới 512 megabyte trong một chuỗi.
##### Example

```
redis 127.0.0.1:6379> SET name "tutorialspoint" 
OK 
redis 127.0.0.1:6379> GET name 
"tutorialspoint"
```


In the above example, **SET** and **GET** are Redis commands, **name** is the key used in Redis and **tutorialspoint** is the string value that is stored in Redis.

**Note** − A string value can be at max 512 megabytes in length.


---
## Hashes

Băm Redis là tập hợp các cặp giá trị khóa. Băm Redis là bản đồ giữa các trường chuỗi và giá trị chuỗi. Do đó, chúng được sử dụng để đại diện cho các đối tượng.


```
redis 127.0.0.1:6379> HMSET user:1 username tutorialspoint password 
tutorialspoint points 200 
OK 
redis 127.0.0.1:6379> HGETALL user:1  
1) "username" 
2) "tutorialspoint" 
3) "password" 
4) "tutorialspoint" 
5) "points" 
6) "200"
```

Here **HMSET, HGETALL** are commands for Redis, while **user − 1** is the key.

Every hash can store up to 2^32 - 1 field-value pairs (more than 4 billion).


---

## Lists

Danh sách Redis chỉ đơn giản là danh sách các chuỗi, được sắp xếp theo thứ tự chèn. Bạn có thể thêm các phần tử vào Danh sách Redis ở phần đầu hoặc phần đuôi.

```
redis 127.0.0.1:6379> lpush tutoriallist redis 
(integer) 1 
redis 127.0.0.1:6379> lpush tutoriallist mongodb 
(integer) 2 
redis 127.0.0.1:6379> lpush tutoriallist rabitmq 
(integer) 3 
redis 127.0.0.1:6379> lrange tutoriallist 0 10  

1) "rabitmq" 
2) "mongodb" 
3) "redis"
```

The max length of a list is 2^32 - 1 elements (4294967295, more than 4 billion of elements per list).

---
## Sets

Sets Redis là một tập hợp các chuỗi không có thứ tự. Trong Redis, bạn có thể thêm, xóa và kiểm tra sự tồn tại của các thành viên trong độ phức tạp thời gian O(1).

```
redis 127.0.0.1:6379> sadd tutoriallist redis 
(integer) 1 
redis 127.0.0.1:6379> sadd tutoriallist mongodb 
(integer) 1 
redis 127.0.0.1:6379> sadd tutoriallist rabitmq 
(integer) 1 
redis 127.0.0.1:6379> sadd tutoriallist rabitmq 
(integer) 0 
redis 127.0.0.1:6379> smembers tutoriallist  

1) "rabitmq" 
2) "mongodb" 
3) "redis"
```

**Note** − In the above example, **rabitmq** is added twice, however due to unique property of the set, it is added only once.

The max number of members in a set is 232 - 1 (4294967295, more than 4 billion of members per set).

---
## Sorted Sets

Sorted Sets Redis tương tự như Sets Redis, bộ sưu tập Chuỗi không lặp lại. Sự khác biệt là, mọi thành viên của Tập hợp đã sắp xếp đều được liên kết với một điểm, được sử dụng để sắp xếp tập hợp theo thứ tự, từ điểm nhỏ nhất đến điểm lớn nhất. Mặc dù các thành viên là duy nhất nhưng điểm số có thể bị lặp lại.

```
redis 127.0.0.1:6379> zadd tutoriallist 0 redis 
(integer) 1 
redis 127.0.0.1:6379> zadd tutoriallist 0 mongodb 
(integer) 1 
redis 127.0.0.1:6379> zadd tutoriallist 0 rabitmq 
(integer) 1 
redis 127.0.0.1:6379> zadd tutoriallist 0 rabitmq 
(integer) 0 
redis 127.0.0.1:6379> ZRANGEBYSCORE tutoriallist 0 1000  

1) "redis" 
2) "mongodb" 
3) "rabitmq"
```
