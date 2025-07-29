Bit Manipulation
-----------------------
1. Number systems:
    - Base 10 (0-9 digits are used to represent a number)
    - Base 2 (0-1 digits are used to represent a number)
    - Base 8 (0-7 digits are used to represent a number)
    - Base 16 (0-E digits are used to represent a number)
    - Convert base 10 to other base - divide by other base util got quotient 0 and write remainder in reverse order
    - Convert other base to 10 base - multiple digit with 10 power as per place of digit
    - Convert other base to other base - convert to base 10 then to other base
2. Bitwise Operator
    - | operator - at least one bit is set then answer is 1
    - & operator - all bit should be set then answer is 1
    - ^ operator - parity of set bits - if odd then answer is 1. if even then answer is 0.
    - Shift operator
        - Left - A << B = A * 2^B --> shift B bits to the left and add B zeroes.
        - Right - A >> B = A/2^B --> shift B bits to the right

3. Bitwise operator properties:
    - &, | and ^ are associative and commulative
        - Associative = (A&B)&C = A&(B&C)
        - Commulative = A&B = B&A
    - A^0 = A
      - XOR with zero will not affect the bit
      - XOR with 1 will toggle the bit.
    - A^A = 0
    - if A ^ B = C then A ^ C = B and B ^ C = A
    - A & B <= MIN(A, B) i.e final answer will have 1 only when both number bit is 1
    - A | B >= MAX(A, B) i.e final answer will have 1 for all set bit in A or B
    - A & 1 is 1 if A is odd else zero
        - LSB (Least significant bit) should be 1 for odd number and 0 for even number
    - A & (A-1) is 0 if A is power of 2
        - A is power of 2 if MSB (Most significant bit) is 1 and all other bits 0