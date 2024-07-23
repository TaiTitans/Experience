
---

```Java
@ExceptionHandler(BindException.class)
@ResponseStatus(HttpStatus.BAD_REQUEST)  // Nếu validate fail thì trả về 400
public String handleBindException(BindException e) {
    // Trả về message của lỗi đầu tiên
    String errorMessage = "Request không hợp lệ";
    if (e.getBindingResult().hasErrors())
        e.getBindingResult().getAllErrors().get(0).getDefaultMessage();
    return errorMessage;
}

```