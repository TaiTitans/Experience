
---
# [@EnableWebSecurity in Spring?](https://stackoverflow.com/questions/44671457/what-is-the-use-of-enablewebsecurity-in-spring)


> Add this annotation to an `@Configuration` class to have the Spring Security configuration defined in any `WebSecurityConfigurer` or more likely by extending the `WebSecurityConfigurerAdapter` base class and overriding individual methods:

Or As this `@EnableWebSecurity` depicts, is used to enable SpringSecurity in our project.



The Result:

The `@EnableWebSecurity` is a marker annotation. It allows Spring to find (it's a `@Configuration` and, therefore, `@Component`) and automatically apply the class to the global `WebSecurity`.

> If I don't annotate any of my class with `@EnableWebSecurity` still the application prompting for username and password.

Yes, it is the default behavior. If you looked at your classpath, you could find other classes marked with that annotation (depends on your dependencies):

- `SpringBootWebSecurityConfiguration`;
- `FallbackWebSecurityAutoConfiguration`;
- `WebMvcSecurityConfiguration`.

Consider them carefully, turn the needed configuration off, or override its behavior.