use crate::solution::Solution;

/**
* @date 2024-07-08
*/
use std::collections::HashMap;

#[allow(dead_code)]
impl Solution {
  pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut hash_map = HashMap::with_capacity(nums.len());

    for (idx, num) in nums.iter().enumerate() {
      hash_map.insert(num, idx);
    }

    for (idx, num) in nums.iter().enumerate() {
      let dif = target - *num;

      if hash_map.contains_key(&dif) && *hash_map.get(&dif).unwrap() != idx {
        return vec![*hash_map.get(&dif).unwrap() as i32, idx as i32];
      }
    }

    vec![0, 0]
  }
}

#[test]
fn test() {
  assert_eq!(Solution::two_sum(vec![2, 7, 11, 15], 9), vec![0, 1]);
}

// https://leetcode.com/problems/two-sum/
