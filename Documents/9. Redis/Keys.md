
---
## Redis Keys Commands

|Sr.No|Command & Description|
|---|---|
|1|[DEL key](https://www.tutorialspoint.com/redis/keys_del.htm)<br><br>This command deletes the key, if it exists.|
|2|[DUMP key](https://www.tutorialspoint.com/redis/keys_dump.htm)<br><br>This command returns a serialized version of the value stored at the specified key.|
|3|[EXISTS key](https://www.tutorialspoint.com/redis/keys_exists.htm)<br><br>This command checks whether the key exists or not.|
|4|[EXPIRE key](https://www.tutorialspoint.com/redis/keys_expire.htm) seconds<br><br>Sets the expiry of the key after the specified time.|
|5|[EXPIREAT key timestamp](https://www.tutorialspoint.com/redis/keys_expireat.htm)<br><br>Sets the expiry of the key after the specified time. Here time is in Unix timestamp format.|
|6|[PEXPIRE key milliseconds](https://www.tutorialspoint.com/redis/keys_pexpire.htm)<br><br>Set the expiry of key in milliseconds.|
|7|[PEXPIREAT key milliseconds-timestamp](https://www.tutorialspoint.com/redis/keys_pexpireat.htm)<br><br>Sets the expiry of the key in Unix timestamp specified as milliseconds.|
|8|[KEYS pattern](https://www.tutorialspoint.com/redis/keys_keys.htm)<br><br>Finds all keys matching the specified pattern.|
|9|[MOVE key db](https://www.tutorialspoint.com/redis/keys_move.htm)<br><br>Moves a key to another database.|
|10|[PERSIST key](https://www.tutorialspoint.com/redis/keys_persist.htm)<br><br>Removes the expiration from the key.|
|11|[PTTL key](https://www.tutorialspoint.com/redis/keys_pttl.htm)<br><br>Gets the remaining time in keys expiry in milliseconds.|
|12|[TTL key](https://www.tutorialspoint.com/redis/keys_ttl.htm)<br><br>Gets the remaining time in keys expiry.|
|13|[RANDOMKEY](https://www.tutorialspoint.com/redis/keys_randomkey.htm)<br><br>Returns a random key from Redis.|
|14|[RENAME key newkey](https://www.tutorialspoint.com/redis/keys_rename.htm)<br><br>Changes the key name.|
|15|[RENAMENX key newkey](https://www.tutorialspoint.com/redis/keys_renamenx.htm)<br><br>Renames the key, if a new key doesn't exist.|
|16|[TYPE key](https://www.tutorialspoint.com/redis/keys_type.htm)<br><br>Returns the data type of the value stored in the key.|
