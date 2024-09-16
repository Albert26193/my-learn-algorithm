// Created by Bob at 2024/09/15 22:36
// leetgo: dev
// https://leetcode.cn/problems/two-sum/
// use crate::utils::oj_input::IO;
use anyhow::Result;
use leetgo_rs::*;

struct Solution;

// @lc code=begin
use std::collections::HashMap;
impl Solution {
  pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut mp = HashMap::new();
    for (i, &num) in nums.iter().enumerate() {
      let rest = target - num;
      if let Some(&index) = mp.get(&rest) {
        return vec![index as i32, i as i32];
      }
      mp.insert(num, i);
    }
    vec![]
  }
}

// @lc code=end
fn main() -> Result<()> {
  let nums: Vec<i32> = deserialize(&read_line()?)?;
  let target: i32 = deserialize(&read_line()?)?;
  let ans: Vec<i32> = Solution::two_sum(nums, target).into();

  println!("\noutput: {}", serialize(ans)?);
  Ok(())
}