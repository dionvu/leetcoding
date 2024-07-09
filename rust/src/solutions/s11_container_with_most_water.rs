use crate::solution::Solution;

/**
* @date 2024-07-08
*/

#[allow(dead_code)]
impl Solution {
  pub fn max_area(height: Vec<i32>) -> i32 {
    let len = height.len();

    let mut max_area = 0;
    let mut left = 0;
    let mut right = len - 1;

    while left < right {
      if height[left] > height[right] {
        max_area = max_area.max(height[right] * (right - left) as i32);
        right -= 1;
      } else {
        max_area = max_area.max(height[left] * (right - left) as i32);
        left += 1;
      }
    }

    max_area
  }
}

#[test]
fn test() {}

// https://leetcode.com/problems/container-with-most-water/description/
