# Socket Networking

Socket là khái niệm trong lập trình mạng. Sử dụng để thiết lập và quản lý kết nối giữ các ứng dụng chạy trên các máy tính khác nhau qua internet. → Socket cho phép ứng dụng giao tiếp thông qua mạng, bằng cách truyền và nhận dữ liệu qua kết nối mạng.

→ Coi socket như một đầu nối ảo giữa 2 thiết bị.

Có 2 loại socket:

- Socket TCP: Kênh kết nối có thứ tự và đáng tin cậy. Dữ liệu được chia làm các gói và đảm bảo đến đích một cách có thứ tự và không bị mất. → Thích hợp ứng dụng cần truyền dữ liệu chính xác tin cậy như game, trình duyệt web….
- Socket UDP: giao thức ko kết nối, k thứ tự, ko đảm bảo → Thích hợp cho các ứng dụng truyền dữ liệu nhanh chóng như tin nhắn, streaming ….

![image-20240402160930712](img\image-20240402160930712.png)
