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
    - log2(1000*1000*1000) = log2(1000^3) = 10 * 3 = 30
    - loga(b) + loba(c) = loga(b*c)
    - log2(1000*1000) = log2(1000) + log2(1000) = 10 + 10 = 20
    - log2(1000^k) = 10 * k
16. Intersection of two ranges (x1,y1) and (x2,y2) = max(x1,x2) <= min(y1,y2)
17. Two way of binary serach
    - Keep maintain ans variable and left, right updated as mid-1, mid+1 respectively
    - Do not keep ans variable and keep updating left, right as mid and in end right will your answer. Here left points to false and right points to true in monotonic predicate function. 
18. In decimal serach space, if L and R are too close and just relative error apart then its not required to serach mid in between them.


Binary search and Interactive Problems
---------------------------
1. In some questions we need to transform the given array to find the solution 
    Example : 
        - F(x) is true if median of subarray length k is greater than equal to x.
            - transformed equation = ai-x + ai+1 -x ...... ai+k-1 - x >= 0
        - F(x) is true if ratio of k elements is greater than equal to x. 
            - Transformed equation= (ai - bi*x) + (ai+1 - bi+1*x) ....... >=0


Interactive Problems
--------------------------
1. Interactive programming is a style of programming problems where your program communicates with a judge during execution, instead of reading all input at once and printing all output at the end.
2. The interaction looks like a conversation:
    - Your program asks the judge a question.
    - The judge immediately replies.
    - Your program uses that reply to decide the next question.
    - Eventually, your program outputs the final answer.
3. Unlike normal competitive programming:
    * You must flush the output after every query.
    * There is usually a limit on the number of queries.
    * Reading input before sending the required output can cause your program to hang (waiting for input that never arrives).
4. In Interactive problems, we ask questions and get outputs and in end print complete result.
5. In interactive problems, we cant use '/n' becoz we need output for every question to get final result.
6. endl flushes the stream and print whatever supposed to be done but /n it wait to everything build up and once its done it prints everything.
7. fmt.Println() directly writes to os.Stdout. But we should use bufio and explicitly flush()
8. To test interactive problems use custom interactor that will acts as judge.
9. Interactive problems are based on Binary/ternary search or randamization (non-deterministic)
10. Interactive problems are just different way of taking input.
11. For interative problems:
    - Write query functions
    - Implement logic



Sliding Window
-------------------
1. Types
    - Fixed size sliding window
    - Dynamic size sliding window or two pointer

2. Time complexity -> It reduces from O(n^2) to O(n)
3. Useful for array based problems with constant subarray size.
4. When to use - calculating some infi for every fixed length subarray in an array
5. Use of 2 Pointers
6. Useful for interview too
7. If there is subarray of fixed size - then use sliding window.
8. The elements added first in the window will be removed first, so sometimes we can optimize our codes by using queue instead of sets or map.
9. In some questions we will use monotonic increasing or descreasing stack or queue to get the answer.


Two Pointers or Variable size sliding window
-----------------------------------------------
1. Good segment technique 1
    - A segment is called good if sum of its elements is <= k where ai >1
    - It means any subarray of segment is also satisfying condition sum <= k
    - In question usually asked largest
    - If segment [L:R] is good then all segements enclosed in it will be good.
    - Try to keep incresing segment size until it is good
2. Size [l, r] = r-l+1
3. Good Segment technique 2
    - Define a good segment as a subarray that follows a particular property
    - Now, all subarray containing that good segment is also good
    - example define a good segment as subarray whose sum is > k
    - now every subarray containing good segment will also have sum > k
    - In question usually asked shortest.
    - If segment [L:R] is good then all segments enclosing it will be good
    - Try to keep decresing segement size until it is good.



Number Theory
--------------
1. Factorization
    - Factors occurs in pair : p*q = n
    - min(p,q) <= square root of n
    - Now, with above two observation we can iterate over smaller number -> i -> 1 to square root of n
    - 
2. n not equal m it means prime representation of n is not equal to prime representation of m.
3. The smallest factor of a number > 1 is always prime
4. If x is a factor of n, and y is a factor of x then y is a factor of n.
5. Every number can be broken down into prime numbers.
6. First factor of a number is prime. 
7. There can be only 1 prime factor > square root of 1
8. Eratosthenes
    - lets say want to check 10^6 elements are prime or not
    - Create an array of size 10^6 and fill all indexes as 1
    - now start from 2nd index and check if its value is 1 then mark all its square numbers as 0.
9. Harmonic series = sum (n/2+n/3....n/n) => nlogn
10. SPF (Smallest prime factor)
    - lets say want to check 10^6 elements are prime or not
    - Create an array of size 10^6 and fill all indexes as 1
    - now start from 2nd index and check if its value is 1 then mark all its square numbers as 0. Plus store smallest prime factor of that number

11. sum of powers of prime numbers = O(logN)
12. Number & sum of divisors
    - N = p1^e1 * p2^e2 .... pk^ek
    - count of divisors = (e1+1)*(e2+1)*.....(ek+1)
    - sum of divisor = multiplication of sum of all possible powers of prime number (geometric progression formula)
13. x^1+x^2+x^3+.....x^k => Geometric progression







