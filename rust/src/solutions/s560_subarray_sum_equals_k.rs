use crate::solution::Solution;

/**
* @date 2024-07-15
*/

// Brute force solution, will revisit and learn the prefix sum solution.
#[allow(dead_code)]
impl Solution {
  pub fn subarray_sum(nums: Vec<i32>, k: i32) -> i32 {
    let mut count = 0;

    for i in 0..nums.len() {
      let mut sum = 0;

      sum += nums[i];

      for w in (i + 1)..nums.len() {
        if sum == k {
          count += 1;
        }

        sum += nums[w];
      }

      if sum == k {
        count += 1;
      }
    }

    count
  }
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn hard_test_case() {
    let nums = vec![4, 3, 2, 1];
    let k = 6;
    let result = Solution::subarray_sum(nums, k);
    assert_eq!(result, 1); // Expected result is 9 (you need to count the actual number of subarrays that sum up to 10)
  }
}

// https://leetcode.com/problems/subarray-sum-equals-k/description/
