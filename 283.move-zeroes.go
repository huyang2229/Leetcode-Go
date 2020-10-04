/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 *
 * https://leetcode-cn.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (59.45%)
 * Likes:    515
 * Dislikes: 0
 * Total Accepted:    118K
 * Total Submissions: 197K
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 *
 * 示例:
 *
 * 输入: [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 *
 * 说明:
 *
 *
 * 必须在原数组上操作，不能拷贝额外的数组。
 * 尽量减少操作次数。
 *
 *
 */

// @lc code=start

func swap(arr *[]int, i, j int) {
	if len(*arr) == 0 || i < 0 || i > len(*arr) || j < 0 || j > len(*arr) {
		return
	}

	tmp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = tmp
}

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// nums[i], nums[j] = nums[j], nums[i]
			swap(&nums, i, j)
			j += 1
		}
	}
}

// @lc code=end

