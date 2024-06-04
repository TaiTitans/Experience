
---

`application.properties`

```Java
logging.level.org.springframework=DEBUG
```

**In File:**

```Java
import org.slf4j.Logger;  
import org.slf4j.LoggerFactory;

private static final Logger logger = LoggerFactory.getLogger(MyClass.class);
public void someMethod() { 
logger.info("Starting method someMethod"); 
logger.debug("Processing data..."); 
if (hasError) { logger.error("An error occurred!", exception); } 
else { logger.info("Method someMethod completed successfully."); } }
```

**Logging Messages:**

The `Logger` interface provides various methods for logging messages at different levels:

- `info(String message)`: Logs an informational message.
- `debug(String message)`: Logs a debug message (useful for debugging purposes).
- `warn(String message)`: Logs a warning message.
- `error(String message)`: Logs an error message.
- There are additional methods for logging messages with exceptions (`info(String message, Throwable throwable)`, etc.).