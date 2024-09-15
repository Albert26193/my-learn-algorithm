// Created by Bob at 2024/09/15 22:36
// leetgo: dev
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

#![allow(unused_imports)]
#![allow(unused_variables)]
#![allow(unused_mut)]
#![allow(dead_code)]
use anyhow::Result;
use leetgo_rs::*;
use std::{collections::HashMap};
use std::fs::File;
use std::io::{self, BufRead, BufReader};

struct Solution;

// @lc code=begin
impl Solution {
  pub fn length_of_longest_substring(s: String) -> i32 {
    let s = s.as_bytes();
    let (mut ans, mut left, mut right) = (0, 0, 0);
    let mut window = [false; 128];
    for (right, &cur) in s.iter().enumerate() {
      let c = cur as usize;
      while window[c] {
        window[s[left] as usize] = false;
        left += 1;
      }
      window[c] = true;
      ans = ans.max(right - left + 1);
    }

    ans as _
  }
}

// @lc code=end

fn main() -> Result<()> {
	let file = File::open("./testcases.txt")?;
	let reader = BufReader::new(file);
	let mut lines = reader.lines();
	let _ = lines.next();	
	let input = lines.next().ok_or_else(|| io::Error::new(io::ErrorKind::InvalidData, "Not enough lines in file"))??;

  let ans: i32 = Solution::length_of_longest_substring(s).into();

  println!("\noutput: {}", serialize(ans)?);
  Ok(())
}
