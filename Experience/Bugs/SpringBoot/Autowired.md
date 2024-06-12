

---
Autowired members must be defined in valid Spring bean (@Component|@Service|...)


Fix it:
```
@SuppressWarnings("SpringJavaInjectionPointsAutowiringInspection")
```