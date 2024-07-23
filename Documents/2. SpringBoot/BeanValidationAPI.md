
---

Bean Validation API (JSR-303 và JSR-380) cung cấp một tập hợp các annotation validation mạnh mẽ, được sử dụng rộng rãi trong các ứng dụng Java. Dưới đây là một số annotation validation phổ biến:

1. **Constraint annotations**:
    
    - `@NotNull`: Đảm bảo giá trị không được null.
    - `@NotEmpty`: Đảm bảo giá trị chuỗi, collection, map hoặc array không được rỗng.
    - `@NotBlank`: Đảm bảo giá trị chuỗi không được null và có ít nhất một ký tự không phải là khoảng trắng.
    - `@Positive`, `@PositiveOrZero`, `@Negative`, `@NegativeOrZero`: Đảm bảo giá trị số là dương, dương hoặc zero, âm, âm hoặc zero.
    - `@Size(min=, max=)`: Đảm bảo giá trị chuỗi, collection, map hoặc array có độ dài nằm trong khoảng cho trước.
    - `@Email`: Đảm bảo giá trị chuỗi có định dạng email hợp lệ.
    - `@Pattern(regexp="")`: Đảm bảo giá trị chuỗi phù hợp với biểu thức chính quy cho trước.
2. **Group annotations**:
    
    - `@GroupSequence`: Xác định thứ tự thực hiện các validation group.
    - `@Valid`: Yêu cầu validation cho một nested object.
3. **Container annotations**:
    
    - `@ElementCollection`: Đánh dấu một collection (như List, Set) để được validate.
    - `@IndexedValidation`: Đánh dấu một indexed collection (như List) để được validate theo index.
4. **Cross-field annotations**:
    
    - `@AssertTrue`, `@AssertFalse`: Đảm bảo một biểu thức boolean luôn đúng/sai.
    - `@ConstraintComposition`: Tạo ra một custom constraint bằng cách kết hợp các constraint khác.
    - `@CrossParameter`: Đảm bảo sự nhất quán giữa các tham số của một phương thức.
5. **Metadata annotations**:
    
    - `@ReportAsSingleViolation`: Báo cáo các lỗi validation như một lỗi duy nhất.
    - `@OverridesAttribute`, `@OverridesAttribute.List`: Ghi đè các thuộc tính của một constraint.

Các annotation này có thể được sử dụng trên các trường, getters, setters, constructors, và các tham số của phương thức. Khi sử dụng chúng trong ứng dụng, các lỗi validation sẽ được tự động xử lý, ví dụ trong các controller của Spring MVC.