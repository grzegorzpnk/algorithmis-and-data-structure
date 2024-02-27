Linked List

A linked list is a linear data structure, in which the elements are not stored at contiguous memory locations. The elements in a linked list are linked using pointers as shown in the below image:

![img.png](img.png)

Head and Tail: The linked list is accessed through the head node, which points to the first node in the list. The last node in the list points to NULL or nullptr, indicating the end of the list. This node is known as the tail node.


Types of LinkedList:

**1. Single-linked list:**

In a singly linked list, each node contains a reference to the next node in the sequence. Traversing a singly linked list is done in a forward direction.

![img_1.png](img_1.png)

**2. Double-linked list:**

In a doubly linked list, each node contains references to both the next and previous nodes. This allows for traversal in both forward and backward directions, but it requires additional memory for the backward reference.

![img_2.png](img_2.png)

**3. Circular linked list:**

In a circular linked list, the last node points back to the head node, creating a circular structure. It can be either singly or doubly linked.

![img_3.png](img_3.png)