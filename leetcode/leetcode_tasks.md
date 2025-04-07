# Leetcode Tasks

---

## [Leetcode Problem: 416. Partition Equal Subset Sum](https://leetcode.com/problems/partition-equal-subset-sum/?envType=daily-question&envId=2025-04-07)

**Topics:** Array, Dynamic Programming | **Difficulty:** Medium

### Problem Statement

Given an integer array `nums`, return `true` if you can partition the array into two subsets such that the sum of the elements in both subsets is equal. Otherwise, return `false`.

**Example 1:**  
Input: `nums = [1,5,11,5]`  
Output: `true`  
Explanation: The array can be partitioned as [1, 5, 5] and [11].

**Example 2:**  
Input: `nums = [1,2,3,5]`  
Output: `false`  
Explanation: The array cannot be partitioned into equal sum subsets.

#### Constraints

- `1 <= nums.length <= 200`
- `1 <= nums[i] <= 100`

### Approach

1. Calculate the total sum of all elements.
2. If the total is odd, return `false` (as it cannot be split equally).
3. Use a 1D dynamic programming array `dp` where `dp[i]` is `true` if a subset sum `i` can be formed.
4. Target subset sum becomes `total / 2`.
5. For each number in `nums`, update the dp array in reverse.

```go
func canPartition(nums []int) bool {
    total := 0
    for _, num := range nums {
        total += num
    }

    if total%2 != 0 {
        return false
    }

    target := total / 2
    dp := make([]bool, target+1)
    dp[0] = true

    for _, num := range nums {
        for j := target; j >= num; j-- {
            dp[j] = dp[j] || dp[j-num]
        }
    }

    return dp[target]
}
```

- **Time Complexity:** `O(n * sum/2)` where `n` is the number of elements in `nums`
- **Space Complexity:** `O(sum/2)` due to the DP array

---
