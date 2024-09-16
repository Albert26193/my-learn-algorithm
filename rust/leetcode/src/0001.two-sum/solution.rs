// Created by Bob at 2024/09/15 22:36
// leetgo: dev
// https://leetcode.cn/problems/two-sum/
use leetcode::utils::scanner::IO;
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

#[cfg(test)]
mod test {
    use super::*;
    use leetcode::utils::test_helper::Tester;

    #[test]
    fn test_solution_1() {
        let tester = Tester::new("src/0001.two-sum/in/", "src/0001.two-sum/out/");
        tester.test_solution(|sc| {
          //let arr:Vec<i32> = sc.read_line_as_vec();
          //let num:i32 = sc.read();
          // println!("{}, {:?}", num, arr);
          //assert_eq!(num, 9);q
          sc.write(format!("{:?}", vec![0, 1]));
        });
    }
}