# coding: utf-8

"""
插入排序 希尔排序(h-sort)
"""

import sys
import time
from sort_base import DataInterface, IntData, generateIntData, time_it


class InsertionSort(object):
    """将元素插入已经排定的列表里边"""
    @staticmethod
    @time_it
    def insertion_sort(data):
        for i in xrange(1, data.len()):
            for j in xrange(i, 0, -1):
                if not data.less(j, j-1):
                    break
                data.swap(j, j-1)


class ShellSort(object):
    """希尔排序：插入排序的改进版本

    """
    @staticmethod
    @time_it
    def shell_sort(data):
        # 1/ find h
        h = 1
        l = data.len()
        while h <= l / 3:
            h = 3 * h + 1

        # 2/ desc h
        while h >= 1:
            for i in xrange(1, l-h+1):
                for j in xrange(i, h-1, -h): # 每次往前步进 h
                    if not data.less(j, j-h):
                        break
                    data.swap(j, j-h)

            h /= 3


if __name__ == '__main__':

    d = generateIntData(10000, 10)

    print d.is_sorted()

    InsertionSort.insertion_sort(d)
    print d.is_sorted()

    d.shuffle()
    ShellSort.shell_sort(d)
    print d.is_sorted()