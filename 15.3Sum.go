/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (25.57%)
 * Likes:    2644
 * Dislikes: 0
 * Total Accepted:    338.6K
 * Total Submissions: 1.1M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？请你找出所有满足条件且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 * 示例：
 *
 * 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 *
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */

// @lc code=start
import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	l := len(nums)

	if l == 0 {
		return res
	}
	// 双指针收敛——左右夹逼，前提是数组是有序数组
	sort.Ints(nums)
	for i := 0; i < l-2; i++ {
		j, k := i+1, l-1
		for j < k {
			if (nums[j] + nums[k] + nums[i]) < 0 {
				j++
			} else if (nums[i] + nums[j] + nums[k]) == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j < l-1 && nums[j] == nums[j+1] {
					j++
					continue
				}
				j++
			} else {
				k--
			}
		}

		for i < l-1 && nums[i] == nums[i+1] {
			i++
			continue
		}
	}

	return res

}

// @lc code=end

