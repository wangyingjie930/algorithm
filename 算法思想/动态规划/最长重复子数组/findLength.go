/**
  @author: wangyingjie
  @since: 2023/3/27
  @desc: https://leetcode.cn/problems/maximum-length-of-repeated-subarray/description/
**/

package 最长重复子数组

// findLength
//  @Description: 思考流程
/**
1. 我确定了状态, dp[i][j]表示了当前比对的位置, i指向了num1, j指向了num2
2. 一般的动态规划数组比较都是由后往前的, 所以dp[i][j]分别表示nums1[i]和nums[j]的公共子数组的最后一个数
3. 接下来我以leetcode给出的例子进行推导
		nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
		首先dp数组肯定要初始化, 所以凡是i, j指向相等的都有默认的1
		那么可得dp[2][0] = 1, 由此可以根据例子等于3进行人脑推算
		dp[3][1] = dp[2][0] + 1 // 为什么我要选dp[2][0]
		dp[4][2] = dp[3][1] + 1 // 为什么我要选dp[3][1]
		根据上面的为什么, 就可以进一步探究if的条件是什么, 从而让我可以推出 dp[i][j] = dp[i-1][j-1] + 1
*/
//  @param nums1
//  @param nums2
//  @return int
func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	maxNum := 0

	for i := 0; i < len(nums1); i++ {
		dp[i] = make([]int, len(nums2))
		for j := 0; j < len(nums2); j++ {

			if nums1[i] == nums2[j] {
				dp[i][j] = 1

				if i-1 >= 0 && j-1 >= 0 && nums1[i-1] == nums2[j-1] {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			}

			if maxNum < dp[i][j] {
				maxNum = dp[i][j]
			}
		}
	}
	return maxNum
}
