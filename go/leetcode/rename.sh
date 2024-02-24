rename_go() {
	local root_dir="$(git rev-parse --show-toplevel)"
	for dir in ${root_dir}/leetcode/go/*/; do
		local dirname="$(basename ${dir})"
		local solution_file="${dir}solution.go"
		local question_file="${dir}question.md"
		if [[ ${dirname} =~ ^([0-9]{4}).*$ ]]; then
			local prefix="${BASH_REMATCH[1]}"
			local new_name="${root_dir}/leetcode/go/${dirname}/lc-${prefix}"
			if [[ -f "${question_file}" ]]; then
				mv "${question_file}" "${new_name}.md"
				echo "${new_name}.md"
			fi
		fi
	done
}

rename_go
