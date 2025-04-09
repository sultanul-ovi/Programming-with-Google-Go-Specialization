# Leetcode Tasks

---

## 01. [Leetcode Problem: 416. Partition Equal Subset Sum](https://leetcode.com/problems/partition-equal-subset-sum/?envType=daily-question&envId=2025-04-07)

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

## 02. [Leetcode Problem: 3375. Minimum Operations to Make Array Values Equal to K](https://leetcode.com/problems/minimum-operations-to-make-array-values-equal-to-k/?envType=daily-question&envId=2025-04-09)

**Topics:** Array, Hash Table | **Difficulty:** Easy


### Problem Statement

You are given an integer array `nums` and an integer `k`. An integer `h` is called **valid** if all values in the array that are **strictly greater than `h`** are identical. You are allowed to perform the following operation on `nums`:

- Select an integer `h` that is valid for the current values in `nums`.
- For each index `i` where `nums[i] > h`, set `nums[i] = h`.

Return the **minimum number of operations** required to make every element in `nums` equal to `k`. If it is **impossible**, return `-1`.


**Example 1:**  
Input: `nums = [5,2,5,4,5]`, `k = 2`  
Output: `2`  
Explanation: Operations with `h = 4` and then `h = 2`.

**Example 2:**  
Input: `nums = [2,1,2]`, `k = 2`  
Output: `-1`  
Explanation: `1 < k`, so it's impossible.

**Example 3:**  
Input: `nums = [9,7,5,3]`, `k = 1`  
Output: `4`  
Explanation: Operations with `h = 7, 5, 3, 1`.


#### Constraints

- `1 <= nums.length <= 100`
- `1 <= nums[i] <= 100`
- `1 <= k <= 100`



### Approach

1. First, check if there exists any `num < k`.
   - If true, return `-1` since we cannot increase any number.
2. Otherwise, collect all **unique values** that are **greater than `k`**.
3. Each such unique value requires a separate operation.
4. Return the count of those unique values.



```go
func minOperations(nums []int, k int) int {
    // Check for impossible case
    for _, num := range nums {
        if num < k {
            return -1
        }
    }

    // Count unique values greater than k
    uniqueSet := make(map[int]bool)
    for _, num := range nums {
        if num > k {
            uniqueSet[num] = true
        }
    }

    return len(uniqueSet)
}
```


- **Time Complexity:** `O(n)` — One pass to check for invalid numbers and one to collect unique values.
- **Space Complexity:** `O(n)` — At most `n` unique values stored in a set.

---
