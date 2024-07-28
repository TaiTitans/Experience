
---
```Java
@Configuration  
public class DataSourceConfiguration {  
    private static final Logger logger = LoggerFactory.getLogger(DataSourceConfiguration.class);  
    @Bean  
    public DataSource dataSource(Environment env) {  
        String url = env.getProperty("DB_URL");  
        String username = env.getProperty("DB_USERNAME");  
        String password = env.getProperty("DB_PASSWORD");  
        try{  
            DataSource dataSource = DataSourceBuilder.create()  
                    .url(url)  
                    .username(username)  
                    .password(password).build();  
            logger.info("==== Database connection successful ====");  
            return dataSource;  
        }catch(Exception e){  
            logger.error("Failed to connect to the database: {}",e.getMessage(), e);  
            throw e;  
        }    }}
```

