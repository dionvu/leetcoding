use crate::solution::Solution;

/**
* @date 2024-07-13
*/

#[allow(dead_code)]
impl Solution {
  pub fn plus_one(digits: Vec<i32>) -> Vec<i32> {
    let mut sum = 0;

    let mut digits = digits;
    digits.reverse();

    for (idx, digit) in digits.iter().enumerate() {
      sum += digit * (idx as i32 * 10 + 1);
    }

    sum += 1;

    let chars: Vec<char> = sum
      .to_string()
      .chars()
      .rev()
      .collect();

    let digits = chars
      .iter()
      .map(|digit| digit.to_digit(10).unwrap() as i32)
      .collect();

    digits
  }
}

#[test]
fn test() {}

// https://leetcode.com/problems/plus-one/description/
