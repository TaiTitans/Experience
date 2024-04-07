#  Verify User Login

```vue
Vue.use(VueRouter);
```

```vue
router.beforeEach((to, from, next) => {
  const isAuthenticated = localStorage.getItem('isAuthenticated'); // Kiểm tra xem người dùng đã đăng nhập chưa
  if (!isAuthenticated && to.path !== '/login') { // Nếu chưa đăng nhập và không phải trang đăng nhập
    localStorage.setItem('redirectPath', to.path); // Lưu đường dẫn hiện tại
    next('/login'); // Chuyển hướng đến trang đăng nhập
  } else {
    next(); // Tiếp tục chuyển hướng đến đích
  }
```

