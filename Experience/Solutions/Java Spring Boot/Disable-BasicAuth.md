
---

Config/SecurityConfig:
Adding Code:

```Java
@Bean  
public SecurityFilterChain securityFilterChain(HttpSecurity http) throws Exception {  
   http.authorizeRequests().anyRequest().permitAll();  
  
    http.csrf().disable();  
    return http.build();  
}
```