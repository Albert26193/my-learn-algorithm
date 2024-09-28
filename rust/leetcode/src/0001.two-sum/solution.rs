// Created by Bob at 2024/09/28 00:12
// leetgo: dev
// https://leetcode.cn/problems/two-sum/
#![allow(unused_imports)]
#![allow(dead_code)]
use lc::utils::lc_input::run_testcases;

use anyhow::Result;
use serde_json::Value;
use leetgo_rs::*;
struct Solution;

// @lc code=begin

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut map = std::collections::HashMap::new();
        for (i, &num) in nums.iter().enumerate() {
            let complement = target - num;
            if let Some(&j) = map.get(&complement) {
                return vec![j as i32, i as i32];
            }
            map.insert(num, i);
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

fn solve(input: &[Value]) -> Vec<i32> {
    let nums: Vec<i32> = deserialize(&input[0].to_string()).unwrap();
    let target = input[1].as_i64().unwrap() as i32;
    Solution::two_sum(nums, target)
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_solution() {
        run_testcases("0001", solve)
    }
}
