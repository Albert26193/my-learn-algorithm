#include <iostream>

#include "../cpp_headers/lc_headers.h"

using namespace std;

int main() {
    int a = 0;
    // for (int i = 0; i < 100; i++) {
    //   cout << i << endl;
    // }

    vector<int> mm = {1, 3, 4, 5};
    vector<string> ss = {"hello", "world"};
    map<int, int> mp;
    map<string, int> mps;
    mps["hello"] = 1;
    mp[0] = 10;
    // mp[1] = 5;
    printIterable(mm);
    printIterable(ss[1]);
    printIterable(ss);
    printIterable(mps);
    printIterable(mp);
    unordered_map<int, int> ump = {{1, 5}};
    printIterable(ump);
    std::vector<std::vector<int>> matrix = {
    {1, 2, 3},
    {4, 5, 6}
};
    printIterable(matrix);
    return 0;
}
