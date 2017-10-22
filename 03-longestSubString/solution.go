package longestSubString

/**
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * Given a string, find the length of the longest substring without repeating characters.
 *
 * Examples:
 * Given "abcabcbb", the answer is "abc", which the length is 3.
 * Given "bbbbb", the answer is "b", with the length of 1.
 * Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */

import (
	"strings"
)

/**
 * @param {string} s
 * @return {int}
 */
func lengthOfLongestSubstring(s string) int {
	start, end, max, length := 0, 0, 0, len(s)
	for start < length && end < length {
		if !strings.Contains(s[start:end], string(s[end])) {
			end = end + 1
			if end-start > max {
				max = end - start
			}
		} else {
			start = start + 1
		}
	}
	return max
}
