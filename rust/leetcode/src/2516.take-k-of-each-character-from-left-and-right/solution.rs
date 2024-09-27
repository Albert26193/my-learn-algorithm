// Created by Bob at 2024/09/27 18:40
// leetgo: dev
// https://leetcode.cn/problems/take-k-of-each-character-from-left-and-right/

// use std::{ fs::File, io::{ BufRead, BufReader }, path::PathBuf };
//

use anyhow::Result;
use serde_json::Value;
use lc::utils::lc_input::run_testcases;

struct Solution;

// @lc code=begin
impl Solution {
    pub fn take_characters(s: String, k: i32) -> i32 {
        println!("s: {}", s);
        println!("k: {}", k);
        0
    }
}

// @lc code=end

fn main() -> Result<()> {
    Ok(())
}

fn solve(input: &[Value]) -> Value {
    let s = input[0].as_str().unwrap();
    let k = input[1].as_i64().unwrap() as i32;
    let result = Solution::take_characters(s.to_string(), k);
    Value::Number(serde_json::Number::from(result))
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_solution() {
        run_testcases("2516", solve)
    }
}