
---
Convert Enity to DTO

Example:

```Java
@Configuration  
public class ModelMapperConfig {  
    @Bean  
    public ModelMapper modelMapper() {  
        ModelMapper modelMapper = new ModelMapper();  
        modelMapper.getConfiguration()  
                .setFieldMatchingEnabled(true)  
                .setFieldAccessLevel(org.modelmapper.config.Configuration.AccessLevel.PRIVATE);  
        return modelMapper;  
    }
```


Convert:

```Java
@Service
public class SupplierService {
    @Autowired
    private SupplierRepository supplierRepository;

    @Autowired
    private ModelMapper modelMapper;

    public SupplierDTO convertToDTO(Supplier supplier) {
        return modelMapper.map(supplier, SupplierDTO.class);
    }

    public Supplier convertToEntity(SupplierDTO supplierDTO) {
        return modelMapper.map(supplierDTO, Supplier.class);
    }
}
```