Binary Search
----------------
- Binary search is an efficient algorithm for finding a target value in a sorted array. Instead of checking every element one by one, it repeatedly cuts the search space in half (left + (right-left)/2).
- The array MUST be sorted
- why left + (right-left)/2 is best?
    - (left+right)/2 with large positive number can cause integer overflow
    - right - left is always a small, positive number (just the gap between them)
    - Dividing that by 2 makes it even smaller
    - Adding it back to left never exceeds right

- Time Complexity 
    - How many steps to find target in n array
        - n -> n/2 -> n/4 -> n/8 ......... -> 1
        - n/2^k = 1
        - k = log n
    - Recurrence relation
        T(n) = T(n/2) + 1
        T(1) = 1
    - O(log n)

- Binary serach can not only apply on sorted array but also on
    - Monotonically Increasing Functions
    - Monotonically Decreasing Functions

- Meaning of Monotonic
    - X increases, Y increases
    - X increases, Y Decreases

- Two Concepts (Same)
    - Binary search on sorted array
    - Binary serach on functions (Monotonically Increasing or decreasing)


- Predicate function
    - Function that return true or false on every input.
    - These functions are used to check if our inputs meets some condition or not.


- odd no * odd no = odd no
- even no * odd no = even no
- even no * even no = even

- Some mathematical point of view
    - Why odd * odd = odd and even * even = even
    - if x is odd then x = 2k+1, now x^2 = (2k+1)^2 = 4k^2 + 4k + 1 (odd)
    - if x is even then x = 2k, now x2 = 4k2 (even)

Monotonic Predicate Function
  - Lets assume True is 1 and false is 0
  - TTTTTTFFFFF,  FFFFTTTTTTT  is Monotonic Predicate function output
  
Binary Serach on Answer
------------------------

1. There will be monotonic predicate function defined on ordered set (search space)
2. In competitive programming, the main challenge is to find monotonic predicate function.
3. For decimal problems, there will always be mentioned upto how many places after decimal. 
4. For decimal problems one of those condition will be given
    - Correct upto x decimal places
    - Output must have relative error <= y. (|real answer - output (users print)| <= y)
5. If relative error requirement is for y = 10^3 then if we print any answer which is correct upto 4 decimal places will work. 
6. Binary serach on decimal places - we can fix the infinite serach space of decimals into finite by fixing the number of decimal places in the answer.
7. Precision issue - decimal has precision issue. lets assume low = 0 and high = 0.000001 if we find mid then 0.0000005 and we are considering only 6 decimal places. thats why condition of loop will be (high-low > precision) not low <= high
8. To avoid precision issue then find one more decimal places which is given in question
9. Summary to resolve precision issue:
    - if asks for x decimal places then do it for x+1
    - if asks relative error for 10^x then do it for x+1 decimal places
10. With decimal numbers : decimal number is a.b
    - if a is big then b must be small
    - if a is small, then b must be big 
    Note: This problem cannot be solved
11. Number of iterations with integer serach space = log(high-low)
12. Number of iterations with decimal serach space = log((high-low) * precision (10^k))
13. Issue with decimal serach space in which mid become high value, to resolve this just find out total iterations and run loop only equal to number of iterations.
14. To solve the problem of infinite loop - find out the number of iterations before writing code. No need to handle any low <= high

15. Some nice results to know while calculating number of iterations
    - log2(1000) = 10
    - loga(b) + loba(c) = loga(b*c)
    - log2(1000*1000) = log2(1000) + log2(1000) = 10 + 10 = 20
    - log2(1000^k) = 10 * k
16. Intersection of two ranges (x1,y1) and (x2,y2) = max(x1,x2) <= min(y1,y2)
17. Two way of binary serach
    - Keep maintain ans variable and left, right updated as mid-1, mid+1 respectively
    - Do not keep ans variable and keep updating left, right as mid and in end right will your answer. Here left points to false and right points to true in monotonic predicate function. 




