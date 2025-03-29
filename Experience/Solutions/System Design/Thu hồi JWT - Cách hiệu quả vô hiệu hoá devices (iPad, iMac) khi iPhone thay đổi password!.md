
---
Lưu blacklist khi logout bao gồm:
Trên redis: TOKEN_BLACK_LIST_(UID)_JIT(trong body kèm id device lấy từ FE)


---
Lúc đổi pass xong sẽ tạo lấy thời gian token ấy lưu vào redis. kèm UID
VD: TOKEN_IAT_AVAILABLE_UID
=> nghĩa là các token thiết bị khác có thời gian sớm hơn thời gian token mới đổi mật khẩu thì bị từ chối.