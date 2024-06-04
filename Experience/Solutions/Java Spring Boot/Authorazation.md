

---
`JwtProvider`

```Java
@Component  
public class JwtTokenProvider {  
  
    @Value("${jwt.secret}")  
    private String secret;  
  
    @Value("${jwt.access-token-expiration-in-ms}")  
    private long accessTokenExpirationInMs;  
  
    @Value("${jwt.refresh-token-expiration-in-ms}")  
    private long refreshTokenExpirationInMs;  
  
    public String getUsernameFromToken(String token) {  
        return getClaimFromToken(token, Claims::getSubject);  
    }  
  
    public Date getExpirationDateFromToken(String token) {  
        return getClaimFromToken(token, Claims::getExpiration);  
    }  
  
    public <T> T getClaimFromToken(String token, Function<Claims, T> claimsResolver) {  
        final Claims claims = getAllClaimsFromToken(token);  
        return claimsResolver.apply(claims);  
    }  
  
    private Claims getAllClaimsFromToken(String token) {  
        return Jwts.parser().setSigningKey(secret).parseClaimsJws(token).getBody();  
    }  
  
    private Boolean isTokenExpired(String token) {  
        final Date expiration = getExpirationDateFromToken(token);  
        return expiration.before(new Date());  
    }  
  
    public String generateAccessToken(UserDTO userDTO) {  
        return Jwts.builder()  
                .setSubject(String.format("%s,%s", userDTO.getUser_id(), userDTO.getUsername()))  
                .setIssuer("TitansDev")  
                .claim("roles", userDTO.getRoles())  
                .setIssuedAt(new Date())  
                .setExpiration(new Date(System.currentTimeMillis()+ accessTokenExpirationInMs))  
                .signWith(SignatureAlgorithm.HS512, secret)  
                .compact();  
    }  
  
    public String generateRefreshToken(UserDTO userDTO) {  
        Map<String, Object> claims = new HashMap<>();  
        claims.put("user_id", userDTO.getUser_id());  
        claims.put("roles", userDTO.getRoles());  
        return doGenerateToken(claims, userDTO.getUsername(), refreshTokenExpirationInMs);  
    }  
  
    private String doGenerateToken(Map<String, Object> claims, String subject, long expirationInMs) {  
        return Jwts.builder()  
                .setClaims(claims)  
                .setSubject(subject)  
                .setIssuedAt(new Date(System.currentTimeMillis()))  
                .setExpiration(new Date(System.currentTimeMillis() + expirationInMs))  
                .signWith(SignatureAlgorithm.HS512, secret)  
                .compact();  
    }  
  
    public String getTokenFromRequest(HttpServletRequest request) {  
        String bearerToken = request.getHeader("Authorization");  
        if (StringUtils.hasText(bearerToken) && bearerToken.startsWith("Bearer ")) {  
            return bearerToken.substring(7);  
        }  
        return null;  
    }  
  
    public boolean validateToken(String token) {  
        try {  
            Jwts.parser().setSigningKey(secret).parseClaimsJws(token);  
            return true;  
        } catch (Exception e) {  
            return false;  
        }  
    }  
    public Claims parseClaims(String token) {  
        return Jwts.parser()  
                .setSigningKey(secret)  
                .parseClaimsJws(token)  
                .getBody();  
    }  
}
```

`JwtFilter`

```Java

public class JwtTokenFilter extends OncePerRequestFilter {  
    private static final Logger logger = LoggerFactory.getLogger(JwtTokenFilter.class);  
  
    private final JwtTokenProvider jwtTokenProvider;  
  
    public JwtTokenFilter(JwtTokenProvider jwtTokenProvider) {  
        this.jwtTokenProvider = jwtTokenProvider;  
    }  
  
    @Override  
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain)  
            throws ServletException, IOException {  
        try {  
            String token = jwtTokenProvider.getTokenFromRequest(request);  
            if (token != null && jwtTokenProvider.validateToken(token)) {  
                logger.debug("Valid JWT token found, setting authentication context");  
                setAuthenticationContext(token, request);  
            } else {  
                logger.debug("No valid JWT token found");  
            }  
        }catch (Exception e){  
            logger.error("Error processing JWT token", e);  
            response.setStatus(HttpServletResponse.SC_FORBIDDEN);  
            return;  
        }  
        filterChain.doFilter(request, response);  
    }  
  
    private void setAuthenticationContext(String token, HttpServletRequest request) {  
        UserDetails userDetails = getUserDetails(token);  
  
        UsernamePasswordAuthenticationToken  
                authentication = new UsernamePasswordAuthenticationToken(userDetails, null, userDetails.getAuthorities());  
  
        authentication.setDetails(  
                new WebAuthenticationDetailsSource().buildDetails(request));  
  
        SecurityContextHolder.getContext().setAuthentication(authentication);  
    }  
  
    private UserDetails getUserDetails(String token) {  
        Claims claims = jwtTokenProvider.parseClaims(token);  
        String username = claims.getSubject();  
        Set<String> roles = new HashSet<>(claims.get("roles", List.class));  
  
        User user = new User();  
        user.setUsername(username);  
        roles.forEach(role -> user.addRole(new Role(role)));  
  
        return user;  
    }  
  
}
```

`SercurityConfig`

```Java
@Configuration  
@EnableWebSecurity  
public class SecurityConfig{  
    private static final Logger logger = LoggerFactory.getLogger(JwtTokenFilter.class);  
    private final JwtTokenProvider jwtTokenProvider;  
  
    public SecurityConfig(JwtTokenProvider jwtTokenProvider) {  
        this.jwtTokenProvider = jwtTokenProvider;  
    }  
  
    @Bean  
    public PasswordEncoder passwordEncoder() {  
        return new BCryptPasswordEncoder();  
    }  
  
    @Bean  
    public SecurityFilterChain securityFilterChain(HttpSecurity http) throws Exception {  
        logger.info("Configuring security filter chain");  
        http.authorizeHttpRequests(authorize -> authorize  
                        .requestMatchers("/api/login").permitAll()  
                        .requestMatchers("/api/**").hasRole("ADMIN")  
                        .requestMatchers("/api/manager/**").hasRole("MANAGER")  
                        .requestMatchers("/api/staff/**").hasRole("STAFF")  
                        .requestMatchers("/api/common/**").hasAnyRole("STAFF", "MANAGER")  
                        .anyRequest().authenticated()  
        ).csrf(csrf -> csrf.disable())  
                .addFilterBefore(new JwtTokenFilter(jwtTokenProvider), UsernamePasswordAuthenticationFilter.class);  
  
        return http.build();  
    }  
  
}
```