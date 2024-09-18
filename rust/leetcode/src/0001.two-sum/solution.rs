// Created by Bob at 2024/09/15 22:36
// leetgo: dev
// https://leetcode.cn/problems/two-sum/
// use lc::utils::scanner::IO;
use std::env;
use anyhow::Result;
use serde::Deserialize;
use std::fs::File;
use std::io::{ BufRead, BufReader };
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
    // 打印当前工作目录
    let current_dir = env::current_dir()?;
    println!("Current directory: {:?}", current_dir);

    let file = File::open("src/0001.two-sum/testcases.txt")?;
    let reader = BufReader::new(file);

    // iter
    let mut lines = reader.lines();
    lines.next();
    let arr: Vec<i32> = deserialize(&lines.next().unwrap().unwrap()).unwrap();
    let target: i32 = deserialize(&lines.next().unwrap().unwrap()).unwrap();

    println!("{:?}", Solution::two_sum(arr, target));
    Ok(())
}

#[cfg(test)]
mod test {
    use super::*;
    use lc::utils::test_helper::Tester;

    #[test]
    fn test_solution() {
        let tester = Tester::new("src/0001.two-sum/in/", "src/0001.two-sum/out/");
        tester.test_solution(|sc| {
            let arr: Vec<i32> = sc.read_line_as_vec();
            let num: i32 = sc.read();
            let ans = Solution::two_sum(arr, num);
            println!("{:?}", ans);
            sc.write_arr(&ans);
        });
    }
}
