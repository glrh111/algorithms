# coding: utf-8

"""
快速排序，及其改良版本：三向快速排序
"""
import sys
import time
from sort_base import DataInterface, IntData, generateIntData, time_it


class QuickSort(object):

    def __init__(self, data):
        self.data = data

    def _partition(self, lo, hi):
        i, j = lo, hi + 1
        while True:
            while True:
                i += 1
                if self.data.less(lo, i) or i == hi:
                    break
            while True:
                j -= 1
                if self.data.less(j, lo) or j == lo:
                    break
            if i >= j:
                break
            self.data.swap(i, j)
        self.data.swap(j, lo)
        return j

    def _quick_sort(self, lo, hi):
        if lo >= hi: # 只有1个元素
            return
        mid = self._partition(lo, hi)
        self._quick_sort(lo, mid-1)
        self._quick_sort(mid+1, hi)

    @staticmethod
    @time_it
    def quick_sort(data):
        o = QuickSort(data)
        o._quick_sort(0, data.len() - 1)


class QuickSort3Way(object):
    """三向切分快速排序"""

    def __init__(self, data):
        self.data = data

    def _partition(self, lo, hi):
        lt, gt, i = lo+1, hi, lo+1
        while i <= gt:
            cmp = self.data.cmp(i, lo)
            if cmp < 0:
                self.data.swap(i, lt)
                lt += 1
                i += 1
            elif cmp == 0:
                i += 1
            else:
                self.data.swap(i, gt)
                gt -= 1

        if lt-1 > lo:
            self.data.swap(lt-1, lo)

        lt -= 1

        return lt, gt

    def _quick_sort(self, lo, hi):
        if lo >= hi:  # 只有1个元素
            return
        lt, gt = self._partition(lo, hi)
        self._quick_sort(lo, lt-1)
        self._quick_sort(gt+1, hi)

    @staticmethod
    @time_it
    def quick_sort(data):
        o = QuickSort3Way(data)
        o._quick_sort(0, data.len() - 1)


class QuickSort3WayUnlimitSpace(object):
    """三向切分快速排序
    没有空间限制，不限制原地排序
    实验证明，这个最快。碉堡了。
    """

    @classmethod
    def _quick_sort(self, o_data):
        """以第一个元素为切分元素"""
        if len(o_data) <= 1:
            return o_data
        v = o_data[0] # 切分元素
        smaller = []
        equaller = []
        bigger = []
        for i in o_data:
            if i < v:
                smaller.append(i)
            elif i == v:
                equaller.append(i)
            else:
                bigger.append(i)
        return self._quick_sort(smaller) + equaller + self._quick_sort(bigger)

    @staticmethod
    @time_it
    def quick_sort(data):
        return QuickSort3WayUnlimitSpace._quick_sort(data)


if __name__ == '__main__':
    sys.setrecursionlimit(1000000)
    d = generateIntData(10000, 10)

    print d.is_sorted()

    QuickSort.quick_sort(d)
    print d.is_sorted()

    d.shuffle()
    QuickSort3Way.quick_sort(d)
    print d.is_sorted()

    d.shuffle()
    d.data = QuickSort3WayUnlimitSpace.quick_sort(d.data)
    print d.is_sorted()
