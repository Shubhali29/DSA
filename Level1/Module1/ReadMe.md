<!-- How to take input from user -->
1. fmt.Scan(&a) - here delimeter is space
2. fmt.Scanln(&b) - here delimeter is new line.

<!-- Find sum of digits, last digit of a number -->
1. num%10 used to fetch digits from right to left
2. num/10 gives remaining number after removing right most digit.
3. Mod 2 meaning % 2
4. If you multiple 5 to n number of times you will always get 25 as last two digit 


<!-- Important String Funtions -->
1. strconv.Atoi() - string to integer conversion
2. strconv.Itoa() - integer to string conversion
3. strings.Join([]byte, delimeter) - Join strings array to one string with provided delimeter
4. strings.ToUpper() - Lower to upper conversion
5. strings.ToLower() - Upper to lower conversion
6. math.Ceil() - calculate ceil of given number
