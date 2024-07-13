use crate::solution::Solution;

/**
* @date 2024-07-13
*/
use std::collections::HashMap;

#[allow(dead_code)]
impl Solution {
  pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
    let mut map: HashMap<String, Vec<String>> = HashMap::with_capacity(strs.len());

    for str in strs.iter() {
      let mut sorted: Vec<char> = str.chars().collect();
      sorted.sort();
      let sorted: String = sorted.iter().collect();

      map
        .entry(sorted)
        .or_insert(vec![])
        .push(str.clone());
    }

    map.into_values().collect()
  }
}

#[test]
fn test() {}

// https://leetcode.com/problems/group-anagrams/
