Questions

1. Given an array, find subset of k elements with maximum sum.
    Hint
        - Write claim - sort the array and pick last k elements
        - other solution - swap any parameter from picked k elements
        - proof mathematically that GA is better than OA.
2. Given an array, find maximum difference between two elements
3.  Given an array, find minimum difference between two elements
4. Given an array, reorder the array so that sum of a[i]*i is maximized.
5. Coin Change problem 
6. Maximum product problem
    - floor(n/2)*ceil(n/2)
    - we are considering only first half where A <= B.
7. Minimum Dot product
8. Activity selection problem
9. Fractional Knapsack problem
    - maximize value V
    - take vi/wi ratio
10. Kadane's Algo
    - max_sum subarray will always start and end with positive number.
    - max sum subarray has every prefix sum containing a positive sum and every suffix containing positive sum.
11. Job Sequencing Problem
12. Find original array from doubled array
13. Arranging the sheep
    - find median block to place all sheeps
14. bracket colouring
    - Hint - RBS and RRBS - on graph prefix sum of RBS is +ve and prefix sum of RRBS is -ve.
15. Duff and weight lifting
    - Assume binary 10101 to make last bit as 1 we need to sum with 1.
    - To make a power of 2, pick minimum power of 2 and at least 2 minimum power of 2 should be present
    - Create map of frequency of each power of 2
    - take half of the frequency, add it in next power of 2 frequency value
    - if left behind with 1 increase count
    - In short, we can create binary of each number add them after that final answer will be count of set bit.
16. Polycarp at the radio
    - Binary serach on frequency [ Binary serach on answer]
    - Maxmimum minimum frequency will be n/m.
17. k complete word
    - Each chain is palindrome
    - cost of changing is = for each chain -> size of chain - max frequency
18. Candy box problem
19. Increasing Subsequence
20. Triangle colouring - Greedy + Combinatorics
    - From each trid, we can select exactly 2 edges
    - ways of selecting 2 edges is depends upon max weight
21. Tape
22. Minimize the error
23. Recover an RBS
