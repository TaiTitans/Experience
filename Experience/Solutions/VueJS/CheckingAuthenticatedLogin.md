# CheckingAuthenticated Login

> Kiểm tra người dùng đang ở trạng thái đăng nhập hay chưa.

```vue
router.beforeEach((to, from, next)=>{
  const isAuthenticated = localStorage.getItem('isAuthenticated')
  const isLoginPage = to.path === "/login"
  const isRegisterLinkClicked = from.path === '/login' && to.path === '/register';
  if(!isAuthenticated && !isLoginPage && !isRegisterLinkClicked){
    localStorage.setItem('redirectPath', to.path)
    next('/login')
  }else{
    next()
  }
})
```

