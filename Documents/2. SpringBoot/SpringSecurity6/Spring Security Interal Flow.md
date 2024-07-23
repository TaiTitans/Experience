

---
![](https://images.viblo.asia/3d6e20d9-8e39-4968-bf3a-1b31dd945488.png)


### Bước 1: Spring Security Filters

- Khi user gửi đi thông tin đăng nhập đến backend, SpringBoot sẽ gọi đến class **AuthenticationFilter.java**(class này nằm ở bên trong **package org.springframework.security.web.authentication**) và chạy đến method **doFilterInternal**.

### Bước 2: Authentication

- Tiếp theo, ở bên trong method **doFilterInternal** nó sẽ gọi đến method **attemptAuthentication**, ở trong này tiếp tục gọi đến method **authenticationManager.authenticate(authentication)**


### Bước 3: Authentication Manager

- Ở bên trong **interface AuthenticationManager** chúng ta có một method là **authenticate**, có nhiều class implement từ interface này, nhưng ở đây chúng ta sẽ chỉ đề cập đến class **ProviderManager.java**


### Bước 4: AuthenticationProvider

- Sau khi click vào method **authenticate**, nó gọi đến **interface AuthenticationProvider**, click vào abstract method thì gọi đến class **AbstractUserDetailsAuthenticationProvider**. File này được implement từ **AuthenticationProvider.java**.

Ở bên trong class này, nó sẽ lại tiếp tục gọi đến method **retrieveUser(...)**. method này trỏ đến method bên trong class **DaoAuthenticationProvider.java**


- Ở đây lại tiếp tục gọi đến method **loadUserByUsername(username)**

### Bước 5: UserDetailService

- Khi click vào method **loadUserByUsername(username)**, Intellij gọi đến **interface UserDetailService**

- Ở đây Spring Security sẽ gọi đến class **InMemoryUserDetailsManager** được implement từ **UserDetailService** interface và lấy ra thông tin của user.
- Ở đây Spring Security sẽ gọi đến class **InMemoryUserDetailsManager** được implement từ **UserDetailService** interface và lấy ra thông tin của user.

### Bước 6: PasswordEncoder

- Ngoài ra sau khi **retrieveUser** ở **bước 4** trong class **AbstractUserDetailsAuthenticationProvider** xong thì chạy tiếp đến logic **additionalAuthenticationChecks**

Method này gọi đến class **DaoAuthenticationProvider**, ở đây có nhiệm vụ kiểm tra xem password user nhập vào sau khi encode có khớp với mật của user hay không


### Bước 7, 8, 9: Complete Security Check

- Sau khi đã kiểm tra xong tất cả các bước thì data sẽ trả ngược lại như trên hình và thêm thông tin User vào **SecurityContextHolder**: **SecurityContextHolder.getContext().setAuthentication(auth);**