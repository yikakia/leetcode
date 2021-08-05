# 743.网络延迟时间
[https://leetcode-cn.com/problems/network-delay-time](https://leetcode-cn.com/problems/network-delay-time) 
## 原题
有 `n` 个网络节点，标记为  `1`  到 `n` 。

给你一个列表  `times` ，表示信号经过 **有向** 边的传递时间。  `times[i] = (u<sub>i</sub>, v<sub>i</sub>, w<sub>i</sub>)` ，其中  `u<sub>i</sub>`  是源节点， `v<sub>i</sub>`  是目标节点， `w<sub>i</sub>`  是一个信号从源节点传递到目标节点的时间。

现在，从某个节点  `K`  发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回  `-1` 。

 

 **示例 1：** 

<img alt="" src="https://assets.leetcode.com/uploads/2019/05/23/931_example_1.png" style="height: 220px; width: 200px;" />

```

输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
输出：2

```
 **示例 2：** 

```

输入：times = [[1,2,1]], n = 2, k = 1
输出：1

```
 **示例 3：** 

```

输入：times = [[1,2,1]], n = 2, k = 2
输出：-1

```
 

 **提示：** 
-  `1 <= k <= n <= 100` 
-  `1 <= times.length <= 6000` 
-  `times[i].length == 3` 
-  `1 <= u<sub>i</sub>, v<sub>i</sub> <= n` 
-  `u<sub>i</sub> != v<sub>i</sub>` 
-  `0 <= w<sub>i</sub> <= 100` 
- 所有 `(u<sub>i</sub>, v<sub>i</sub>)` 对都 **互不相同** （即，不含重复边）
 
**标签**
`深度优先搜索` `广度优先搜索` `图` `最短路` `堆（优先队列）` 


## 
```go

```
>
