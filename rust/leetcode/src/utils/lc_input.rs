use std::fs::File;
use std::io::{ BufRead, BufReader };
use std::path::PathBuf;
use anyhow::{ Context, Result };
// use leetgo_rs::deserialize;
use serde::Serialize;
use serde_json::Value;
use serde_json;

pub struct TestCase {
    pub input: Vec<serde_json::Value>,
    pub output: serde_json::Value,
}

pub struct ProblemPaths {
    pub source_file: PathBuf,
    pub testcase_file: PathBuf,
}

pub fn get_problem_paths(problem_number: &str) -> Result<ProblemPaths> {
    let base_path = PathBuf::from("src");
    let mut problem_dir = None;

    for entry in base_path.read_dir()? {
        let entry = entry?;
        let path = entry.path();
        if path.is_dir() {
            if let Some(file_name) = path.file_name() {
                if file_name.to_string_lossy().starts_with(problem_number) {
                    problem_dir = Some(path);
                    break;
                }
            }
        }
    }

    let problem_dir = problem_dir.ok_or_else(|| anyhow::anyhow!("Problem directory not found"))?;
    let source_file = problem_dir.join("solution.rs");
    let testcase_file = problem_dir.join("testcases.txt");

    if !source_file.exists() {
        return Err(anyhow::anyhow!("Source file not found"));
    }
    if !testcase_file.exists() {
        return Err(anyhow::anyhow!("Testcase file not found"));
    }

    Ok(ProblemPaths {
        source_file,
        testcase_file,
    })
}

pub fn read_testcases(testcase_file: &PathBuf) -> Result<Vec<TestCase>> {
    let file = File::open(testcase_file).with_context(||
        format!("Unable to open file: {:?}", testcase_file)
    )?;
    let reader = BufReader::new(file);
    let mut testcases = Vec::new();
    let mut lines = reader.lines();

    while let Some(Ok(_)) = lines.next() {
        // Skip "input:" line
        let mut input = Vec::new();
        while let Some(Ok(line)) = lines.next() {
            if line.trim() == "output:" {
                break;
            }
            let value: serde_json::Value = serde_json
                ::from_str(&line)
                .with_context(|| format!("Unable to parse input line: {}", line))?;
            input.push(value);
        }

        let output = if let Some(Ok(line)) = lines.next() {
            serde_json
                ::from_str(&line)
                .with_context(|| format!("Unable to parse output line: {}", line))?
        } else {
            return Err(anyhow::anyhow!("Missing output line"));
        };

        testcases.push(TestCase { input, output });

        // Skip empty line
        lines.next();
    }

    Ok(testcases)
}

pub fn get_testcases(problem_number: &str) -> Result<Vec<TestCase>> {
    let problem_paths = get_problem_paths(problem_number)?;
    read_testcases(&problem_paths.testcase_file)
}

pub trait IntoValue {
    fn into_value(self) -> Value;
}

impl<T: Serialize> IntoValue for T {
    fn into_value(self) -> Value {
        serde_json::to_value(self).unwrap_or(Value::Null)
    }
}

pub fn run_testcases<T>(problem_number: &str, solve: impl Fn(&[Value]) -> T) where T: IntoValue {
    let testcases = get_testcases(problem_number).unwrap();
    for (i, case) in testcases.iter().enumerate() {
        let result = solve(&case.input);
        let result_value = result.into_value();
        println!("\nCurrent test case {}:", i + 1);
        println!("Input: {}", &serde_json::to_string(&case.input).unwrap());
        println!("Output: {}", &serde_json::to_string(&result_value).unwrap());
        println!("Expected: {}", &serde_json::to_string(&case.output).unwrap());
        assert!(result_value == case.output, "Test case {} failed", i);
    }
}