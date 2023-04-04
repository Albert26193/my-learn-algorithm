let res = 0;

function dfs(begin: number, curLen: number, n: number): void {
  if (curLen >= n) {
    res += 1;
    return;
  }

  for (let i = begin; i < 5; i++) {
    dfs(i, curLen + 1, n);
  }
}

function countVowelStrings(n: number): number {
  dfs(0, 0, n);
  return res;
}

console.log(countVowelStrings(2));
