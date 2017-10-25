package judgeRouteCircle

/**
 * https://leetcode.com/problems/judge-route-circle/description/
 *
 * Initially, there is a Robot at position (0, 0). Given a sequence of its moves, judge if this robot makes a circle, which means it moves back to the original place.
 * The move sequence is represented by a string. And each move is represent by a character. The valid robot moves are R (Right), L (Left), U (Up) and D (down). The output should be true or false representing whether the robot makes a circle.
 *
 * Example 1:
 *   Input: "UD"
 *   Output: true
 * Example 2:
 *   Input: "LL"
 *   Output: false
 */

/**
 * @param {string} moves
 * @return {bool}
 */
func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, char := range moves {
		switch char {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		case 'R':
			x++
		}
	}
	return x == 0 && y == 0
}
