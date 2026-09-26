Dynamic Programming
----------------------
1. Divide and Conquer strategy
    - Divide big problem into small sub problems. Continue this until you end up in trivial subproblem [Easier subproblem]
    - Divide in a such a way that it is possible to combine subproblems solution to get bigger problem solution
2. Solve the subproblem once and then store the result 
    - In real life its a caching technique
3. For storing result
    - Use map or array
    - Array can be used if input is +ve integer and range is not very large
        - default value will be outside the range of output of function
    - If output range from -infinity to +infinity then 
        - either go with map
        - or create two array one for storing calculated subproblem values and other bool array to indicate it is solved before or not.
4. Every subproblem will be called as a State - Meaning of a subproblem
5. Transition - Calculate answer of big problem from answer of small subproblems.
6. Every DP problem has four terms
    - State
    - Transition
    - Trivial problem
    - Final answer
7. Time complexity of subproblem will not encounter in time complexity of bigger problem
8. Space complexity is equal to number of unique subproblems.
    - Number of states * space required for each state
9. Time complexity
    - Estimate = Number of unique states * Transition time for each state
    - Exact = Total transition time for all states

10. 1+1/2+1/3+1/4 .... + 1/n -> harmonic series = its sum is less than log(n)
11. No of primes upto n is approximate n/log(n).
12. number of factors of i is sqrt(i) but in actual its cuberoot(i)
13. Recursive vs Iterative DP
    - Use recursive DP when most of subproblems are invalid
    - DP optimizations dont work on recursive DP
14. Converting recursive DP to iterative DP
    - all the states that a particular state depends must be evaluated before that state

15. Tip -> divide & conquer -> state -> transition -> base case -> final answer



