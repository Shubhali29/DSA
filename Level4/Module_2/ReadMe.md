Dynamic Programming - Intermediate
--------------------------------------
Space and Transition Optimization
- To reduce Time and space complexity in DP
- They both work only on iterative DP

Space Optimization
    - Identify on what states our current state depends
    - Find do we need answer of all subproblems at all time?
    - then only store answer of relevant state and discard other states for space optimization
    - For 256MB and 512MB memory, we can't store more than 10^7 integers

Transition Optimization
    - Observe transition equation
    - Can you do some pre-computations to evaluate equation faster
    - Using range query data structures
        - Segment Tree
        - Sparse table


State Elimination and Cyclic DP States
---------------------------------------------
State Elimination:
    - Ask youself do we need all the parameters in the DP state
    - If you have dp[a][b][c] and a+b=c, do you need to store c as a parameter or can you just compute it on spot
    - if you can compute a parameter in dp state from other parameters, no need to store it.
    - Always remove lasgest state

Cycling DP States
    - What happens when your state depends on itself
       - dp[i] depending on dp[i] itself
    - It used in DP + Probability problems
    - Direction of moves should be only in one direction