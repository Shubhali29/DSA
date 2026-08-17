Questions
----------
1. Given a list of strings, now perform following operations
    - check X is present in list of strings or not
    - Find number of strings that has prefix X 
There will be Q queries.

    - Hint : Use binary serach and hashing

2. Given N strings and Q queries, in each query check a string X can be created from at least 2 strings concatenation
     - Create prefix trie and prefix array
     - create suffix trie and suffix array
     - find i in X string where prefix[i] = suffix[i+1] = T
     - TC = O(Q.X)


3. There are N integers each ai <  10^9, find maximum XOR of 2 number
4. There are N integers each ai <  10^9, find maximum AND of 2 numbers
    - Hint cannot be solved by binary trie as it does not have deterministic approach.