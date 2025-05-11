Time Complexity
------------------
1. Elementary Operations:
    - Operations that take constant Time
    - Like Airthmetic Operations, Input & Output, Comparison on primitive data types
    - System can execute 10^8 Elementary operations in 1 Second.

2. Time complexity is a function to describe approximate amount of operations an algorithm requires for a given input.

3. Three ways to define Time complexity:
    - Atleast, Best case, Lower bound -> Omega notation
    - Atmost, Worst case, Upper bound -> Big O Notation
    - Average -> Theta Notation

4. Rules of Big O notation:
    - Should not consider constant (Lessor in value)  -> O(n + k) = O(n)
    - Should not consider factors (multiple and divide) -> O(n/2) = O(n)
    - Should consider fastest growing function for each variable -> O(n2 + n) =  O(n2)
    - Can never be 0. It has to be O(1)
    - Ex : 4n^2 + n + 4 (m^3 + 5) + 10 = O(n^2 + m^3)

5. Expected time CheatSheet

    -----------------------------------------
    Fesable Big O |  Maximum  | Example 
    function      |    N      | Algorithms
    ------------------------------------------
    O(N!)         |   10      | All permutation of list
    O(N3)         |   400     | Multiplication of metrics
    O(N2)         |   5000    | Square grid, Bubble sort, insertion sort
    O(N sqrt(N))  |   10^5    | Usually related to factoring
    O(N log N)    |    10^6   | Merge sort, binary serach for N times
    O(N)          |   10^7    |  Linear serach, reversing an array, string Comparison
    O(sqrt(N))    |  10^12    |  factors of number
    O(log N), O(1)|  10^18    | Binary Search, Constant time formulas