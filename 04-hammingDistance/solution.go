package hammingDistance

/**
 * https://leetcode.com/problems/hamming-distance/description/
 * The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
 * Given two integers x and y, calculate the Hamming distance.
 * Note:
 *   0 ≤ x, y < 231.
 *
 * Example:
 * Input: x = 1, y = 4
 * Output: 2
 * Explanation:
 * 1   (0 0 0 1)
 * 4   (0 1 0 0)
		  ↑   ↑
 *The above arrows point to positions where the corresponding bits are different.
*/

func hammingDistance(x int, y int) int {
	z := x ^ y
	n := 0
	for z > 0 {
		z &= z - 1
		n = n + 1
	}
	return n
}
