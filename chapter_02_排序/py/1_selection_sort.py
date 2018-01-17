# coding: utf-8

import sys
import time
from sort_base import DataInterface, IntData, generateIntData, time_it


class SelectionSort(object):
    """每次排定一个最小的元素, 放在列表最前面"""

    @staticmethod
    @time_it
    def selection_sort(data):
        l = data.len()
        for i in range(l-1):
            min_index = i
            for j in range(i, l): # 找到最小的元素
                if data.less(j, min_index):
                    min_index = j
            data.swap(min_index, i)


class BubbleSort(object):
    """每两个元素做比较，将较大的放在后边; 这样每次可以排定最大的一个元素"""

    @staticmethod
    @time_it
    def bubble_sort(data):
        l = data.len()
        for i in range(l-1, 0, -1):
            for j in range(0, i):
                if data.less(j+1, j):
                    data.swap(j+1, j)


if __name__ == '__main__':
    d = generateIntData(10000, 10)

    print d.is_sorted()

    SelectionSort.selection_sort(d)
    print d.is_sorted()

    d.shuffle()
    BubbleSort.bubble_sort(d)
    print d.is_sorted()