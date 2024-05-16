
---

	In Java, Access modifiers help to restrict the scope of a class, constructor, variable, method, or data member. It provides security, accessibility, etc to the user depending upon the access modifier used with the element.

## Types of Access Modifiers in Java

There are four types of access modifiers available in Java: 

1. Default – No keyword required
2. Private
3. Protected
4. Public

![](https://media.geeksforgeeks.org/wp-content/uploads/20230409122917/Access-Modifiers-in-Java-1.webp)

## 1. Default

When no access modifier is specified for a class, method, or data member – It is said to be having the **default*** access modifier by default. The data members, classes, or methods that are not declared using any access modifiers i.e. having default access modifiers are accessible ***only within the same package****.

## 2. Private

The private access modifier is specified using the keyword ***private****. The methods or data members declared as private are accessible only **within the class** in which they are declared.

- Top-level classes or interfaces can not be declared as private because
    - private means “only visible within the enclosing class”.
    - protected means “only visible within the enclosing class and any subclasses”

## 3. Protected

The methods or data members declared as protected are ***accessible within the same package or subclasses in different packages.***

VN: Protected có thể được sử dụng khi được kế thừa từ một class khác còn Private thì không.
## 4. Public


- The public access modifier has the **widest scope** among all other access modifiers.
- Classes, methods, or data members that are declared as public are ***accessible from everywhere*** in the program. There is no restriction on the scope of public data members.

---
### More

![](https://media.geeksforgeeks.org/wp-content/uploads/20240516114231/Access-Modifiers-in-Java-2.webp)

