
---
Stack is kind of data structure which allows operations on data only at one end. It allows access to the last inserted data only. Stack is also called LIFO (Last In First Out) data structure and Push and Pop operations are related in such a way that only last item pushed (added to stack) can be popped (removed from the stack).

---
## Stack Representation

![](https://www.tutorialspoint.com/dsa_using_java/images/stack.jpg)


We're going to implement Stack using array in this article.


## Basic Operations

Following are two primary operations of a stack which are following.

- **Push** − push an element at the top of the stack.
    
- **Pop** − pop an element from the top of the stack.
    

There is few more operations supported by stack which are following.

- **Peek** − get the top element of the stack.
    
- **isFull** − check if stack is full.
    
- **isEmpty** − check if stack is empty.

---

## Push Operation

![](https://www.tutorialspoint.com/dsa_using_java/images/stack_push.jpg)

```java
// push item on the top of the stack public void push(int data) {

   if(!isFull()){
      // increment top by 1 and insert data 
      intArray[++top] = data;
   }else{
      System.out.println("Cannot add data. Stack is full.");
   }      
}
```
## Pop Operation

![](https://www.tutorialspoint.com/dsa_using_java/images/stack_pop.jpg)


```java
// pop item from the top of the stack public int pop() {
   // retrieve data and decrement the top by 1 
   return intArray[top--];        
}
```

