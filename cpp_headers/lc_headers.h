#ifndef LC_HEADERS_H
#define LC_HEADERS_H

#include <algorithm>
#include <bitset>
#include <cassert>
#include <cmath>
#include <cstdio>
#include <cstdlib>
#include <deque>
#include <forward_list>
#include <fstream>
#include <functional>
#include <iomanip>
#include <iostream>
#include <iterator>
#include <map>
#include <numeric>
#include <queue>
#include <random>
#include <set>
#include <sstream>
#include <stack>
#include <string>
#include <tuple>
#include <type_traits>
#include <typeinfo>
#include <unordered_map>
#include <unordered_set>
#include <utility>
#include <variant>
#include <vector>

template <typename T>
void printIterable(const T& iterable);

template <>
void printIterable<std::string>(const std::string& iterable) {
    std::cout << iterable << " ****Length: " << iterable.size() << std::endl
              << std::endl;
}

template <>
void printIterable<std::vector<int>>(const std::vector<int>& iterable) {
    std::cout << "[ ";
    for (auto it = iterable.begin(); it != iterable.end(); ++it) {
        std::cout << *it;
        if (it + 1 != iterable.end()) {
            std::cout << ", ";
        }
    }
    std::cout << " ] ";
    std::size_t length = std::distance(iterable.begin(), iterable.end());
    std::cout << " ****Length: " << length << std::endl
              << std::endl;
}

template <>
void printIterable<std::vector<std::string>>(const std::vector<std::string>& iterable) {
    std::cout << "[ ";
    for (auto it = iterable.begin(); it != iterable.end(); ++it) {
        std::cout << *it;
        if (it + 1 != iterable.end()) {
            std::cout << ", ";
        }
    }
    std::cout << " ] ";
    std::size_t length = std::distance(iterable.begin(), iterable.end());
    std::cout << " ****Length: " << length << std::endl
              << std::endl;
}

template <>
void printIterable<std::vector<std::vector<int>>>(const std::vector<std::vector<int>>& iterable) {
    std::cout << "{" << std::endl << std:: endl;
    for (auto it = iterable.begin(); it != iterable.end(); ++it) {
        std::cout << "  ";
        printIterable<std::vector<int>>(*it);
    }
    std::cout << "}" << std::endl
              << std::endl;
}

template <typename M>
void printMap(const M& iterable) {
    std::cout << "{ ";
    for (auto it = iterable.begin(); it != iterable.end(); ++it) {
        std::cout << it->first << ": " << it->second;
        if (std::next(it) != iterable.end()) {
            std::cout << ", ";
        }
    }
    std::cout << " } ";
    std::size_t length = std::distance(iterable.begin(), iterable.end());
    std::cout << " ****Length: " << length << std::endl
              << std::endl;
}

template <>
void printIterable<std::map<std::string, int>>(const std::map<std::string, int>& iterable) {
    printMap(iterable);
}

template <>
void printIterable<std::map<int, int>>(const std::map<int, int>& iterable) {
    printMap(iterable);
}

template <>
void printIterable<std::unordered_map<std::string, int>>(const std::unordered_map<std::string, int>& iterable) {
    printMap(iterable);
}

template <>
void printIterable<std::unordered_map<int, int>>(const std::unordered_map<int, int>& iterable) {
    printMap(iterable);
}
#endif  // LC_HEADERS_H
