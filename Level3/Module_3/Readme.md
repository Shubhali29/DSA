Questions
------------
1. Find max sum of subarray of size k
2. Find max distinct element in subarray of size k.
3. find maximum sun and all distict element of subarray of size k
 Note- return entire array or index for above questions
4. For each subarray of k, find index of minimum element in that subarray
5. Find first negative element index from each subarray of size k in an array
6. Find median of k subarray. 
    hint - 
    small → a max-heap holding the smaller half of the current window
    large → a min-heap holding the larger half
    Keep them balanced: len(small) == len(large) or len(small) == len(large) + 1
    Median = top of small (odd k), or average of both tops (even k)

The tricky part: removing elements from a heap
A heap only lets you pop from the top efficiently — you can't cheaply remove an arbitrary element in the middle when the window slides. The trick is lazy deletion:

When an element leaves the window, don't remove it immediately. Just mark it in a delayed map (count of pending removals).
Before you look at (or pop) a heap's top, prune it: if the top is marked for deletion, pop it for real and decrement the map.
Track smallSz/largeSz as the effective size (excluding pending deletions) — this is what rebalancing logic uses.

   --> We need to maintain ceil of k/2 in left and floor of k/2 in right.


7. Sliding window cost
    - If you want to make all elements equal, then you change them to the median
    - cost of 1 -> x + cost of 7 -> x = constant
    - Example to make 5 -> (5-1) + (7-5) = 6
    - To make 3 -> (3-1) + (7-3) = 6
    - From 1 to 7 cost does not matter its always same, but to choose between lets say (2,4) best optimal is median 
    - Formula -> left.size()*median - leftSum + rightSum - right.size()*median
8. Maximum subarray sum 2
    - Prefix sum
    - max(prefixSum[j] - prefixSum[i-1], prefixSum[k] - prefixSum[i-1], prefixSum[l]-prefixSum[i-1]) here [a,b] = [j,k,l]



