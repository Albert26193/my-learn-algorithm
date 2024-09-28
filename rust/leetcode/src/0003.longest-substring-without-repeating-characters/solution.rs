#![allow(unused_imports)]
#![allow(dead_code)]
use lc::utils::lc_input::run_testcases;
use serde_json::Value;

use anyhow::Result;
use leetgo_rs::*;

struct Solution;

// @lc code=begin

impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        0
    }
}

// @lc code=end

fn solve(input: &[Value]) -> i32 {
    let s = input[0].as_str().unwrap().to_string();
    Solution::length_of_longest_substring(s)
}

#[allow(dead_code)]
mod tests {
    use super::*;
    #[test]
    fn test_solution() {
        run_testcases("0003", solve)
    }
}

fn main() -> Result<()> {
    let s: String = deserialize(&read_line()?)?;
    let ans: i32 = Solution::length_of_longest_substring(s).into();

    println!("\noutput: {}", serialize(ans)?);
    Ok(())
}
