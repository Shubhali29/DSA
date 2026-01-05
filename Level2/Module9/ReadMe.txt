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
