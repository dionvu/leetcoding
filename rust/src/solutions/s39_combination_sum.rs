use crate::solution::Solution;

/**
* @date 2024-07-10
*/

#[allow(dead_code)]
impl Solution {
  pub fn combination_sum(candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
    let mut combo: Vec<i32> = Vec::new();
    let mut res = Vec::new();

    Self::backtrace_combination_sum(&mut combo, &mut res, &candidates, target);

    res
  }

  fn backtrace_combination_sum(
    combo: &mut Vec<i32>,
    res: &mut Vec<Vec<i32>>,
    candidates: &Vec<i32>,
    target: i32,
  ) {
    let sum = combo.iter().sum::<i32>();

    if sum > target {
      return;
    }
    if sum == target {
      res.push(combo.clone());
      return;
    }

    for i in 0..candidates.len() {
      combo.push(candidates[i]);
      let new_candidates = &candidates[i..];
      Self::backtrace_combination_sum(combo, res, &new_candidates.to_vec(), target);
      combo.pop();
    }
  }
}
#[test]
fn test() {}

// https://leetcode.com/problems/combination-sum/
