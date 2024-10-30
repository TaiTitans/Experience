

---


Linked List is a sequence of links which contains items. Each link contains a connection to another link. Linked list the second most used data structure after array. Following are important terms to understand the concepts of Linked List.

- **Link** − Each Link of a linked list can store a data called an element.
    
- **Next** − Each Link of a linked list contain a link to next link called Next.
    
- **LinkedList** − A LinkedList contains the connection link to the first Link called First.


![](https://www.tutorialspoint.com/dsa_using_java/images/dsa_linkedlist.jpg)

## Insertion Operation

Insertion is a three step process:

1. Create a new Link with provided data.
    
2. Point New Link to old First Link.
    
3. Point First Link to this New Link.

![](https://www.tutorialspoint.com/dsa_using_java/images/dsa_linkedlist_insertfirst.jpg)


```java
//insert link at the first location
public void insertFirst(int key, int data){
   //create a link
   Link link = new Link(key,data);
   //point it to old first node
   link.next = first;
   //point first to new first node
   first = link;
}
```

## Deletion Operation

Deletion is a two step process:

1. Get the Link pointed by First Link as Temp Link.
    
2. Point First Link to Temp Link's Next Link.

![](https://www.tutorialspoint.com/dsa_using_java/images/dsa_linkedlist_deletefirst.jpg)
```java
//delete first item
public Link deleteFirst(){
   //save reference to first link
   Link tempLink = first;
   //mark next to first link as first first = first.next;
   //return the deleted link
   return tempLink;
}
```

## Navigation Operation

Navigation is a recursive step process and is basis of many operations like search, delete etc.:

1. Get the Link pointed by First Link as Current Link.
    
2. Check if Current Link is not null and display it.
    
3. Point Current Link to Next Link of Current Link and move to above step.

![](https://www.tutorialspoint.com/dsa_using_java/images/dsa_linkedlist_navigation.jpg)
