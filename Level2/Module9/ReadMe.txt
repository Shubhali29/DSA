Stack, Queues and Dequeue
---------------------------
1. **Stack**
    - Add and Remove both performed from one end (back) only.
    - Follows LIFO
    - Operations
        - push(type value)
        - pop()
        - top() type
        - empty() bool
        - size() 
    - All operations complexity is O(1)
    - If stack is empty then top and pop will give runtime error so handle it properly.
   
2. **Queue**
    - Add is performed from back end and Remove is performed from front end only.
    - Follows FIFO
    - Operations
        - push(type value)
        - pop()
        - front() type
        - empty() bool
        - size() 
    - All operations complexity is O(1)
    - If stack is empty then front and pop will give runtime error so handle it properly.

2. **DeQueue**
    - Add and Remove is performed from both end (front and back).
    - Operations
        - push_back(type value)
        - push_front(type value)
        - pop_back()
        - pop_front()
        - front() type
        - back() type
        - empty() bool
        - size() 
    - All operations complexity is O(1)

3. **Monotonic Stacks**
    - It is used to find following things in O(1) time after precomputation time of O(n)
        - Next Great element towards Left/Right
        - Next Small element towards Left/Right

4. **Heap**
    - Binary Tree - Each node can have at most 2 child
    - Complete Binary Tree - Nodes are filled from left to right and each node has two children except at last level.
    - Heap is a complete binary tree with some heap order rule.
    - Types of heap 
        - Min - Parent node is smaller than equal to its left and right child
        - Max - Parent node is greater than equal to its left and right child
    - If Heap is represented as an array then 
        - Parent node is at index i
        - Its left child will be at 2i+1
        - Its right child will be at 2i+2
        - If we need to find the parent of any node at index i then i-1/2
    - To build a heap from array
        - Traverse from last non leaf node and call heapify_down on that node to move it to right place.
    - Heap sort
        - Build heap
        - swap first and last element
        - decrease size of heap and call heapify_down on top element.
