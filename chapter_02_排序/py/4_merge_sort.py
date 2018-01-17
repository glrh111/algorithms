# coding: utf-8

"""
归并排序. 自顶向下topDown，自底向上downTop
"""
import sys
import time
from sort_base import DataInterface, IntData, generateIntData, time_it


class MergeSort(object):
    """递归排序
    然后将字数组的排序结果，合并为有序的数组
    """

    def __init__(self, data):
        self.data = data

    def _merge(self, lo, mid, hi):
        # 1/ 备份
        aux = []
        for i in xrange(lo, hi+1):
            aux.append(self.data.data[i])  # 直接访问内部变量，不提倡

        # 2/ 合并
        # [i, mid] 有序 [mid+1, hi] 有序 现在将他们合并
        i, j = lo, mid+1
        for k in xrange(lo, hi+1):
            if i > mid: # 左边用完了,
                self.data.data[k] = aux[j-lo]
                j += 1
            elif j > hi: # 右边用完了
                self.data.data[k] = aux[i-lo]
                i += 1
            elif aux[i-lo] > aux[j-lo]: # 取右边的
                self.data.data[k] = aux[j-lo]
                j += 1
            else:
                self.data.data[k] = aux[i - lo]
                i += 1

    def _top_down(self, lo, hi):
        if lo >= hi:
            return
        mid = (lo+hi)/2
        self._top_down(lo, mid)
        self._top_down(mid+1, hi)
        self._merge(lo, mid, hi)

    @staticmethod
    @time_it
    def top_down(data):
        """自顶向下"""
        o = MergeSort(data)
        o._top_down(0, data.len()-1)

    @staticmethod
    @time_it
    def down_top(data):
        """自底向下
        最开始排定每2个，之后4个，之后8个
        """
        o = MergeSort(data)
        sz = 1
        l = data.len()
        while sz < l:
            for i in range(0, l-sz+1, sz*2):
                o._merge(i, i+sz-1, min(i+2*sz-1, l-1))
            sz *= 2


if __name__ == '__main__':
    sys.setrecursionlimit(1000000)
    d = generateIntData(100000, 100)

    print d.is_sorted()

    MergeSort.top_down(d)
    print d.is_sorted()

    d.shuffle()
    MergeSort.down_top(d)
    print d.is_sorted()

