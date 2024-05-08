#include <algorithm>
#include <iostream>
#include <string>
#include <vector>

using namespace std;

struct comb {
    int index;
    string scores;
};

struct cp {
    int total;
    int seq;
    vector<int> lost;
};

cp getTags(string str) {
    cp ans;
    int n = str.size();
    int cnt = 0, maxSeq = 0;
    int seq = 0;
    vector<int> lost;

    for (int i = 0; i < n; i++) {
        if (str[i] == '1') {
            cnt++;
            seq++;
            if (seq > maxSeq) {
                maxSeq = seq;
            }
        } else if (str[i] == '0') {
            lost.push_back(i);
            seq = 0;
        }
    }

    ans.total = cnt;
    ans.seq = maxSeq;
    ans.lost = lost;
    return ans;
}

int main() {
    int n, m;
    cin >> n >> m;

    vector<comb> persons(n);

    for (int i = 0; i < n; i++) {
        cin >> persons[i].scores;
        persons[i].index = i + 1;
    }

    sort(persons.begin(), persons.end(), [](const comb &a, const comb &b) {
        cp tag1 = getTags(a.scores), tag2 = getTags(b.scores);

        if (tag1.total > tag2.total) {
            return true;
        }

        if (tag1.seq > tag2.seq) {
            return true;
        }

        for (int i = 0; i < min(tag1.lost.size(), tag2.lost.size()); i++) {
            if (tag1.lost[i] > tag2.lost[i]) {
                return true;
            }
        }
        return a.index < b.index;
    });

    for (const auto &p : persons) {
        cout << p.index << " ";
    }

    return 0;
}
