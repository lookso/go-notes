/*
@Time : 2019/5/26 5:37 PM 
@Author : Tenlu
@File : lengthOfLongestSubstring
@Software: GoLand
*/
package main

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

func main(){
	lengthOfLongestSubstring("pwwkew")
}
func lengthOfLongestSubstring(str string){
	//int[] m = new int[256];
	//
	//int res = 0, left = 0;
	//
	//for (int i = 0; i < s.length(); i++) {
	//left = Math.max(left, m[s.charAt(i)]);
	//
	//res = Math.max(res, i - left + 1);
	//
	//m[s.charAt(i)] = i + 1;
	//}
	//return res;

	//maxLength = 0
	//for i,_ in enumerate(s):  #这里没有使用range函数
	//count = 0
	//usedChar = str()
	//for j in s[i:]:
	//if j not in usedChar:
	//usedChar += j
	//count += 1
	//if maxLength < count:#这里没有使用max函数
	//maxLength = count
	//else:
	//break
	//return maxLength
	//---------------------
	//	作者：coordinate_blog
	//来源：CSDN
	//原文：https://blog.csdn.net/qq_17550379/article/details/80547777
	//版权声明：本文为博主原创文章，转载请附上博文链接！


}

