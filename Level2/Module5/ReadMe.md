Recursion
-----------------
1. Recursion happens when a function call itself on different set of input parameters.
2. Used when the solution of current problem involves first solving small sub problem.
3. Recursion helps us to solve bigger problem by solving smaller problem.

Recursion Tree
--------------
1. A recursion tree is similar like a mind map of the function call
2. It helps us to understand how function acts.


Basic Structure of Recursion
---------------------------------
1. Parameters to start the function
2. Appropiate base case(s) to end the recursion.
3. Recursively solve sub problems
4. Process the result and return the value

Recursion call stack
-----------------------
1. When recursion is executed, the function calls are placed in a stack
2. The first call is at the bottom of stack, and last call is at top of the stack
3. After call is executed, it is popped from stack and returned value is passed to the current top element.


Time Complexity 
-----------------
Time Complexity = Total number of nodes in Recursion tree * Time complexity of a single node