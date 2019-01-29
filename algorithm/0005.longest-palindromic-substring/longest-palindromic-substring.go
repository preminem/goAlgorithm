package problem0005

/*
Topic:
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example:
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example:
Input: "cbbd"
Output: "bb"

Thought：
回文的特点是：假定回文的长度为l，x为任意字符
当l为奇数时，回文的正中间段只会是，“x”，或“xxx”，或“xxxxx”，或...
当l为偶数时，回文的正中间段只会是，“xx”，或“xxxx”，或“xxxxxx”，或...
同一个字符串的任意两个回文substring的正中间段，不会重叠。

Summary:
充分利用查找对象的特点，可以加快查找速度。
*/
