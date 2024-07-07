
----

```xml
<dependency>  
    <groupId>redis.clients</groupId>  
    <artifactId>jedis</artifactId>  
    <version>5.1.3</version>  
</dependency>
```

Config:
```Java
@Configuration  
public class RedisConfig {  
  
    private static final Logger logger = LoggerFactory.getLogger(RedisConfig.class);  
  
    @Value("${REDIS_HOST}")  
    private String redisHost;  
  
    @Value("${REDIS_PORT}")  
    private int redisPort;  
  
    @Value("${REDIS_PASSWORD}")  
    private String redisPassword;  
  
    @Bean  
    public JedisPool jedisPool() {  
        logger.info("Connecting to Redis at {}:{}", redisHost, redisPort);  
        GenericObjectPoolConfig<Jedis> poolConfig = new GenericObjectPoolConfig<>();  
        poolConfig.setJmxEnabled(false); // Disable JMX  
        return new JedisPool(poolConfig, redisHost, redisPort, 2000, redisPassword);  
    }}
```

