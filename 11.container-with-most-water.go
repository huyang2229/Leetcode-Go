/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 *
 * https://leetcode-cn.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (61.37%)
 * Likes:    1888
 * Dislikes: 0
 * Total Accepted:    294.3K
 * Total Submissions: 458.3K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
 * (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 *
 * 说明：你不能倾斜容器，且 n 的值至少为 2。
 *
 *
 *
 *
 *
 * 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 *
 *
 *
 * 示例：
 *
 * 输入：[1,8,6,2,5,4,8,3,7]
 * 输出：49
 *
 */

// @lc code=start
// 辅助工具类
func getMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 解法一：暴力解法
// func maxArea(height []int) int {
// 	if len(height) == 0 {
// 		return 0
// 	}

// 	max := 0
// 	len := len(height)

// 	for i := 0; i < len-1; i++ {
// 		for j := 0; j < len; j++ {
// 			smaller := getMin(height[i], heignt[j])

// 			area := (j - i) * smaller
// 			max = getMax(max, area)
// 		}

// 	}
// 	return max
// }

// 解法二： 收敛法
func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}

	max := 0
	left, right := 0, len(height)-1
	for left < right {
		// 谁更小，谁就挪动，虽然 x 轴的宽度减小了，但不排除 y 轴出现大值
		if height[left] < height[right] {
			max = getMax(max, (right-left)*height[left])
			left++
		} else {
			max = getMax(max, (right-left)*height[right])
			right--
		}
	}
	return max
}

// @lc code=end

