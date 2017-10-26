package reverseString

/**
 * https://leetcode.com/problems/reverse-string/description/
 *
 * Write a function that takes a string as input and returns the string reversed.
 *
 * Example:
 *   Given s = "hello", return "olleh".
 */

/**
 * @param {string} s
 * @return {string}
 */
func reverseString(s string) string {
	var rs []byte
	for i := len(s); i > 0; i-- {
		rs = append(rs, s[i-1])
	}
	return string(rs)
}
