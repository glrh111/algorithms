## 3.1 符号表

符号表：一张抽象的表格，我们会将信息存储在其中，然后按照指定的键来搜索并获取这些信息。有时称为字典，有时称为索引。

三种实现高效符号表的经典类型：二叉查找树，红黑树，散列表。

典型的符号表应用

应用|查找的目的|键|值
---|---|---|---
字典|找出单词的释义|单词|释义
图书索引|找出相关的页码|术语|一串页码
文件共享|找出歌曲的下载地址|歌曲名|url之类的
账户管理|处理交易|账户号码|交易详情
网络搜索|找出相关网页|关键字|网页名称
编译器|找出符号的类型和值|变量名|类型和值

### 3.1.1 API

符号表API

返回类型 | 方法名 | desc
---|---|---|---
|ST()创建一张符号表|
void | put(Key key, Value value) | 存入键值对
Value | get(Key key) | 获取键key对应的值
void | delete(Key key) | 删除键值对
boolean | contains(Key key) | 键key在表中是否有对应的值  get(key) != nul
boolean | isEmpty() | 表是否为空  size() == 0
int | size() | 键值对数量
Iterable<Key> | keys() | 所有键的集合

为了保持代码的一致、简洁和实用，以下是几个设计决策
+ 范型
+ 重复的键
  + 每个键只对应一个值
  + 新存入的键值対和表中已有的冲突时，新值会替换旧值
+ 空键null 键不能位null
+ 空值null 值不能有null
  + 键不存在时可以返回null
  + 导致可以使用 get() 的返回值判断键是否存在
  + 导致可以将值设置为 null(put) 来实现删除 
+ 删除操作
  + 延时删除 将键对应的值置为空，然后在某个时候删除所有值为空的键 put(key, null) 是一种实现
  + 即时删除 立刻从表中删除指定的键 delete(key) 是这种实现
+ 便捷方法 ：用已经实现的方法实现
+ 迭代 使用keys返回包含所有键的迭代器
+ 键的等价性 怎么判断键是否存在的等价

### 3.1.2 有序符号表

在典型应用中，键都是Comparable对象，许多符号表的实现，都利用Comparable带来的键的有序性，更好得实现put get方法。

以下是有序的泛型符号表API

返回值类型 | 方法名 | 说明
---|---|---
| ST() | 创建一张有序符号表
void | put(Key key, Value value) | 将键值対存入表中
Value | get(Key key) | 获取键key对应的值
void | delete(Key key) | 从表中删除键key
boolean | contains(Key key) | 表中是否有某个键
boolean | isEmpty() | 表是否为空
int | size() | 表中键值対数量
Key | min() | 最小的键
Key | max() | 最大的键
Key | floor(Key key) | 小于等于key的最大键
Key | ceiling(Key key) | 大于等于key的最小键
int | rank(Key key) | 小于key的键的数量
Key | select(int k) | 排名为k的键
void | deleteMin() | 删除最小的键 delete(min())
void | deleteMax() | 删除最大的键 delete(max())
int | size(Key lo, Key hi) | \[lo..hi\] 之间键的数量 
Iterable<Key> | keys(Key lo, Key hi) | \[lo..hi\]之间的所有键 
Iterable<Key> | keys() | 表中的所有键的集合，已排序 keys(min(), max())

#### 3.1.2.1 最大键和最小键

#### 3.1.2.2 向下取整和向上取整

floor ceiling

#### 3.1.2.3 排名和选择

+ 0 ~ size()-1 的所有i满足： i == rank(select(i))
+ 所有key都满足: key == select(rank(key))

挑战是，实现插入删除查找的同时，快速实现这两种操作。

#### 3.1.2.4 范围查找

size(), keys() 能够接受两个参数。

能够处理这类查询，是有序符号表在实践中被广泛应用的重要原因。

#### 3.1.2.5 例外情况

如果一个方法需要返回一个键但表中没有合适的键返回时，约定抛出一个异常。

符号表为空时，min(), max(), deleteMin(), deleteMax(), floor(), ceiling() 都会抛出异常。

k < 0, 或者 k >= size() 时，select(k) 也会抛出异常。

#### 3.1.2.6 便捷方法

#### 3.1.2.7 键的等价性

#### 3.1.2.8 成本模型

### 3.1.3 用例举例

#### 3.1.3.1 行为测试用例 

用一串字符串与各自对应的索引，来作为key value。

#### 3.1.3.2 性能测试用例

使用网上知名测试内容，进行单词统计。



  



