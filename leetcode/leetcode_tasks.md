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

## 03. [Leetcode Problem: 2999. Count the Number of Powerful Integers](https://leetcode.com/problems/count-the-number-of-powerful-integers/description/?envType=daily-question&envId=2025-04-10)

**Topics:** Math, String, Dynamic Programming | **Difficulty:** Hard


### Problem Statement

You are given three integers `start`, `finish`, and `limit`, and a string `s` representing a positive integer.

A **positive integer x** is called **powerful** if:
- `s` is a **suffix** of `x` (i.e., x ends with `s`), and
- **each digit** in `x` is **at most `limit`**.

Return the total number of powerful integers in the range `[start..finish]`. Return `0` if none exist.


**Example 1:**  
Input: `start = 1`, `finish = 6000`, `limit = 4`, `s = "124"`  
Output: `5`  
Explanation: Valid integers: 124, 1124, 2124, 3124, 4124.

**Example 2:**  
Input: `start = 15`, `finish = 215`, `limit = 6`, `s = "10"`  
Output: `2`  
Explanation: Valid integers: 110, 210.

**Example 3:**  
Input: `start = 1000`, `finish = 2000`, `limit = 4`, `s = "3000"`  
Output: `0`  
Explanation: 3000 is beyond the finish boundary.


#### Constraints
- `1 <= start <= finish <= 10^15`
- `1 <= limit <= 9`
- `1 <= s.length <= floor(log10(finish)) + 1`
- `s` only contains numeric digits and each digit `<= limit`
- `s` does not have leading zeros



#### Approach

This problem uses **digit dynamic programming**. Key observations:

- Any valid powerful number must be of the form `p * 10^L + S`, where:
  - `S` is the integer form of suffix `s`
  - `L = len(s)`
  - `p` is a prefix (possibly 0)

#### Steps:
1. **Determine valid range of p**:
   - Minimum `p`: Smallest value such that `p * M + S >= start`
   - Maximum `p`: Largest value such that `p * M + S <= finish`
2. For each `p`, ensure that **all digits of p are ≤ limit**.
3. Use digit DP (`countValid`) to count numbers `p` in valid range whose digits are all ≤ `limit`
4. Result = `countValid(p_max) - countValid(p_min - 1)`


```go
func numberOfPowerfulInt(start, finish int64, limit int, s string) int64 {
    S, _ := strconv.ParseInt(s, 10, 64)
    L := len(s)
    M := int64(math.Pow10(L))

    if finish < S {
        return 0
    }

    var p_min int64
    if start <= S {
        p_min = 0
    } else {
        p_min = (start - S + M - 1) / M
    }
    p_max := (finish - S) / M

    if p_min > p_max {
        return 0
    }

    return countValid(p_max, limit) - countValid(p_min-1, limit)
}

func countValid(X int64, limit int) int64 {
    if X < 0 {
        return 0
    }

    strX := strconv.FormatInt(X, 10)
    n := len(strX)
    digits := make([]int, n)
    for i, ch := range strX {
        digits[i] = int(ch - '0')
    }

    dp := make([][][]int64, n+1)
    for i := range dp {
        dp[i] = make([][]int64, 2)
        for j := 0; j < 2; j++ {
            dp[i][j] = make([]int64, 2)
            for k := 0; k < 2; k++ {
                dp[i][j][k] = -1
            }
        }
    }

    var rec func(pos, tight, started int) int64
    rec = func(pos, tight, started int) int64 {
        if pos == n {
            return 1
        }
        if dp[pos][tight][started] != -1 {
            return dp[pos][tight][started]
        }
        var res int64 = 0
        upper := limit
        if tight == 1 && digits[pos] < upper {
            upper = digits[pos]
        }
        for d := 0; d <= upper; d++ {
            if d > limit {
                break
            }
            newTight := 0
            if tight == 1 && d == digits[pos] {
                newTight = 1
            }
            newStarted := started
            if started == 0 && d != 0 {
                newStarted = 1
            }
            res += rec(pos+1, newTight, newStarted)
        }
        dp[pos][tight][started] = res
        return res
    }

    return rec(0, 1, 0)
}
```
- **Time Complexity:** `O(D * limit)` where D is the number of digits in `p_max`
- **Space Complexity:** `O(D * 2 * 2)` for the DP table

---
