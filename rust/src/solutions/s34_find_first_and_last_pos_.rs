use crate::solution::Solution;

/**
* @date 2024-07-10
*/

#[allow(dead_code)]
impl Solution {
  pub fn search_range(nums: Vec<i32>, target: i32) -> Vec<i32> {
    for (i, num) in nums.iter().enumerate() {
      if *num == target {
        let starting = i;
        let mut ending = i;

        while ending + 1 < nums.len() && nums[ending + 1] == target {
          ending += 1;
        }

        return vec![starting as i32, ending as i32];
      }
    }

    vec![-1, -1]
  }
}

#[test]
fn test() {
  assert_eq!(Solution::search_range(vec![1], 1), vec![0, 0]);
}

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
