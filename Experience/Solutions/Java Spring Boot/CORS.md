
---
Config:

```Java
...
       .cors(cors -> cors.configurationSource(corsConfigurationSource()))  
           ...
  
    return http.build();  
}  
@Bean  
public CorsConfigurationSource corsConfigurationSource() {  
    CorsConfiguration configuration = new CorsConfiguration();  
    configuration.addAllowedOrigin("http://localhost:3000");  
    configuration.addAllowedMethod("*");  
    configuration.addAllowedHeader("*");  
    configuration.setAllowCredentials(true);  
  
    UrlBasedCorsConfigurationSource source = new UrlBasedCorsConfigurationSource();  
    source.registerCorsConfiguration("/**", configuration);  
    return source;  
}
```