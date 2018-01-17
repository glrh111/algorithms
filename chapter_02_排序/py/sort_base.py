#! coding: utf-8

from abc import abstractmethod
import random
import functools
import time


class DataInterface(object):

    @abstractmethod
    def less(self, i, j):
        """i < j ?"""
        pass

    @abstractmethod
    def swap(self, i, j):
        """交换元素"""
        pass

    @abstractmethod
    def len(self):
        """含有多少元素"""
        pass

    @abstractmethod
    def cmp(self, i, j):
        pass

    @abstractmethod
    def is_sorted(self):
        """是否排定顺序"""
        pass

    @abstractmethod
    def shuffle(self):
        pass


class IntData(DataInterface):

    def __init__(self, lst):
        self.data = lst # 用列表存储数据

    def less(self, i, j):
        return self.data[i] < self.data[j]

    def swap(self, i, j):
        self.data[i], self.data[j] = self.data[j], self.data[i]

    def len(self):
        return len(self.data)

    def cmp(self, i, j):
        if self.data[i] > self.data[j]:
            return 1
        elif self.data[i] == self.data[j]:
            return 0
        else:
            return -1

    def is_sorted(self):
        for i in xrange(self.len()-1):
            if self.data[i] > self.data[i+1]:
                return False
        return True

    def shuffle(self):
        random.shuffle(self.data)

    def __str__(self):
        return ', '.join(map(lambda x: str(x), self.data))


def generateIntData(count, max_value):
    lst = []
    for i in xrange(count):
        lst.append(random.randint(0, max_value))
    return IntData(lst)

def time_it(f):
    @functools.wraps(f)
    def wrapper(*args, **kwargs):
        from_sec = time.time()
        re = f(*args, **kwargs)
        to_sec = time.time()
        print '耗费的时间：', to_sec - from_sec
        return re
    return wrapper
