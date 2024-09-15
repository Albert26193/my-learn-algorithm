rename_go() {
	local root_dir="$(git rev-parse --show-toplevel)"
	if [[ -z "${root_dir}" ]]; then
		echo "Error: Not a git repository."
		return 1
	fi

	local rust_leetcode_dir="${root_dir}/rust/leetcode/src"
	if [[ ! -d "${rust_leetcode_dir}" ]]; then
		echo "Error: Directory ${rust_leetcode_dir} does not exist."
		return 1
	fi

	for dir in "${rust_leetcode_dir}"/*/; do
		local dirname="$(basename "${dir}")"
		local solution_file="${dir}solution.rs"
		local question_file="${dir}question.md"

		if [[ ! -d "${dir}" ]]; then
			echo "Warning: ${dir} is not a directory. Skipping."
			continue
		fi

		if [[ ! -f "${solution_file}" ]]; then
			echo "Warning: ${solution_file} does not exist. Skipping."
			continue
		fi

		if [[ ! -f "${question_file}" ]]; then
			echo "Warning: ${question_file} does not exist. Skipping."
			continue
		fi

		# do action
		if [[ ${dirname} =~ ^([0-9]{4}).*$ ]]; then
			local prefix="${BASH_REMATCH[1]}"
			local new_name="${rust_leetcode_dir}/${dirname}/lc-${prefix}"
			mv "${question_file}" "${new_name}.md"
			echo "${new_name}.md"
		else
			echo "Warning: Directory ${dirname} does not match the expected pattern. Skipping."
		fi

	done
}

rename_go
