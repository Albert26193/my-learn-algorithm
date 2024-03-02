var buf = '';
process.stdin.on('readable', function () {
  var chunk = process.stdin.read();
  if (chunk) buf += chunk.toString();
});
let getInputArgs = line => {
  return line.split(' ').filter(s => s !== '').map(x => parseInt(x));
}
process.stdin.on('end', function () {
  buf.split('\n').forEach(function (line, lineIdx) {
    if (lineIdx === 1) {
      let arr = getInputArgs(line);
      margeSort(arr, 0, arr.length - 1);
      console.log(arr.join(' '))
    }
  });
});

/*
作者：gaobowen
链接：https://www.acwing.com/activity/content/code/content/147105/
*/