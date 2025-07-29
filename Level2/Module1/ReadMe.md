Prefix Sum
================
Prefix sum is a powerful technique that can be used to preprocess an array to facilitate fast subarray sum queries without modifying original array.

Prefix[i] => sum of all elements from index 0 to i.

Differnce Array
=====================
A Difference array can be used to perform multiple range update where we need to find the final state of arrays only after performing all queries.

Difference array helps to solve this in O(n) time and space complexity.

We can process each range update in O(1) timre.

Diff array - by how much every index value will change


Coordinate Compression
========================
Steps:
1. Find important points
2. Map out importants points
3. Solve question using only important points
4. Figure out answer for all remaining points.