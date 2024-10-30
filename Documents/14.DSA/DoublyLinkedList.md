
---
Doubly Linked List is a variation of Linked list in which navigation is possible in both ways either forward and backward easily as compared to Single Linked List. Following are important terms to understand the concepts of doubly Linked List

- **Link** − Each Link of a linked list can store a data called an element.
    
- **Next** − Each Link of a linked list contain a link to next link called Next.
    
- **Prev** − Each Link of a linked list contain a link to previous link called Prev.
    
- **LinkedList** − A LinkedList contains the connection link to the first Link called First and to the last link called Last.


## Doubly Linked List Representation

![](https://www.tutorialspoint.com/dsa_using_java/images/dsa_doublylinkedlist.jpg)

---
## Insertion Operation

Following code demonstrate insertion operation at beginning in a doubly linked list.

```java
//insert link at the first location
public void insertFirst(int key, int data){
   //create a link
   Link link = new Link(key,data);

   if(isEmpty()){
      //make it the last link
      last = link;
   }else {
      //update first prev link
      first.prev = link;
   }

   //point it to old first link
   link.next = first;
   //point first to new first link
   first = link;
}
```

## Deletion Operation

Following code demonstrate deletion operation at beginning in a doubly linked list.

```java
//delete link at the first location
public Link deleteFirst(){
   //save reference to first link
   Link tempLink = first;
   //if only one link
   if(first.next == null){
      last = null;
   }else {
      first.next.prev = null;
   }
   first = first.next;
   //return the deleted link
   return tempLink;
}
```


## Insertion at End Operation

Following code demonstrate insertion operation at last position in a doubly linked list.


```java
//insert link at the last location
public void insertLast(int key, int data){
   //create a link
   Link link = new Link(key,data);

   if(isEmpty()){
      //make it the last link
      last = link;
   }else {
      //make link a new last link
      last.next = link;     
      //mark old last node as prev of new link
      link.prev = last;
   }

   //point last to new last node
   last = link;
}
```