# 1178.猜字谜
[https://leetcode-cn.com/problems/number-of-valid-words-for-each-puzzle](https://leetcode-cn.com/problems/number-of-valid-words-for-each-puzzle) 
## 原题
外国友人仿照中国字谜设计了一个英文版猜字谜小游戏，请你来猜猜看吧。

字谜的迷面 `puzzle` 按字符串形式给出，如果一个单词 `word` 符合下面两个条件，那么它就可以算作谜底：
- 单词 `word` 中包含谜面 `puzzle` 的第一个字母。
- 单词 `word` 中的每一个字母都可以在谜面 `puzzle` 中找到。

	例如，如果字谜的谜面是 "abcdefg"，那么可以作为谜底的单词有 "faced", "cabbage", 和 "baggage"；而 "beefed"（不含字母 "a"）以及 "based"（其中的 "s" 没有出现在谜面中）。
返回一个答案数组 `answer` ，数组中的每个元素 `answer[i]` 是在给出的单词列表 `words` 中可以作为字谜迷面 `puzzles[i]` 所对应的谜底的单词数目。

 

 **示例：** 

```
输入：
words = ["aaaa","asas","able","ability","actt","actor","access"], 
puzzles = ["aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"]
输出：[1,1,3,2,4,0]
解释：
1 个单词可以作为 "aboveyz" 的谜底 : "aaaa" 
1 个单词可以作为 "abrodyz" 的谜底 : "aaaa"
3 个单词可以作为 "abslute" 的谜底 : "aaaa", "asas", "able"
2 个单词可以作为 "absoryz" 的谜底 : "aaaa", "asas"
4 个单词可以作为 "actresz" 的谜底 : "aaaa", "asas", "actt", "access"
没有单词可以作为 "gaswxyz" 的谜底，因为列表中的单词都不含字母 ';g';。

```


 **提示：** 
-  `1 <= words.length <= 10^5` 
-  `4 <= words[i].length <= 50` 
-  `1 <= puzzles.length <= 10^4` 
-  `puzzles[i].length == 7` 
-  `words[i][j]` , `puzzles[i][j]` 都是小写英文字母。
- 每个 `puzzles[i]` 所包含的字符都不重复。

**标签**
`位运算` `哈希表` 


## 二进制压缩+枚举子集
```go
func convertByte(b byte) (res uint32) {
	return 1 << uint32(b-'a')
}
func convertStr(str string) (res uint32) {
	for i := range str {
		res |= convertByte(str[i])
	}
	return
}
func findNumOfValidWords(words []string, puzzles []string) []int {
	const puzzleLen int = 7
	wordCnt := map[uint32]int{}
	for _, word := range words {
		mask := convertStr(word)
		if bits.OnesCount(uint(mask)) <= puzzleLen {
			wordCnt[mask]++
		}
	}
	res := make([]int, len(puzzles))
	for i, puzzle := range puzzles {
		first := convertByte(puzzle[0])
		mask := convertStr(puzzle[1:])
		subset := mask
		for ok := true; ok; ok = subset != mask {
			res[i] += wordCnt[subset|first]
			subset = (subset - 1) & mask
			// uint32(0)-1 等于 uint32(1111111..)
			// 此时相与得到的即 mask 本身。
		}
	}
	return res
}
```
>执行用时: 80 ms
内存消耗: 11.4 MB

这题本来想到了二进制压缩的，但是没想到怎么枚举子集，于是只用二进制压缩的时候 `TLE` 了。

先说二进制压缩吧，二进制压缩在我的理解里面就是用二进制位的 `0,1` 来表示该位对应信息。这个在用 `OpenGL` 或者其他的一些开源库的时候会有体会，因为一些参数是通过相与来组合而成的。

比如表示一个数组这一位是不是偶数，用 0 表示不是偶数，1表示是偶数。那么就可以这样：

```go
nums := []int{1, 2, 3, 4, 5}
p := 0
for i , num := range nums {
	if num & 1 == 0 {
		p |= 1<<i 
	}
}
```

得到的 `p` 的第 `i` 位就表示 `nums[i]` 是不是偶数。

当然这里可能不太好理解，因为二进制数的表示从右到左的，而数组我们常认为是从左到右的。这个理解了就好了。

在知道了二进制压缩的方法之后，就是利用信息了。

下面用 `b(x)` 表示将字符串 `x` 转换为对应的二进制状态。这里的二进制状态中第 `i` 位表示第 `i` 个字母在这个字符串中是否出现过。于是要确定

- 单词 `word` 中包含谜面 `puzzle` 的第一个字母，可以通过下面的方式确定

  ```go
  if b(puzzle[:1])&b(word) == b(puzzle[:1]){
      return true
  }
  ```

  

- 单词 `word` 中的每一个字母都可以在谜面 `puzzle` 中找到的话就可以通过下面的方式来确定

	```go
	if b(word) & b(puzzle) == b(word){
		return true
	}
	```

于是综合以上两种方法就可以确定这个 `word` 是不是这个 `puzzle` 对应的了。

值得注意的是，不同的字符串可能对应同一个二进制状态，比如 `aaaa` 和 `a` 就都是`(1)` 表示这个字符串中出现过 `a` 。因此需要开一个`map` 来统计这个二进制位出现过多少次。

这里解决了二进制压缩只是完成了一半，因为题目要求

- `1 <= words.length <= 10^5`
- `1 <= puzzles.length <= 10^4`

如果用遍历的方法的话，最差就是 $10^9$ 次的循环。这铁定超时，不用看了。

于是就要用枚举子集的方式，查看 `puzzle` 对应的谜底们是否存在

简单地说就是 `(10110)` 这个 `puzzle` 可以有多少个对应的谜底存在呢？

先谈枚举子集的方法。

`(10110)`对应的子集一定小于它。因为表示的状态只能少于等于它，不能比它多，不然就不是它的子集了。要主要这个数本质是状态的集合。而 `(10110)` 的子集是哪些呢？ `(10110)` `(10100)` `(10010)` `(10000)` `(00000)`

发现了，它本身也是自己的子集，而 `0` 也是它的子集。

那么怎么求呢？一个简单的思路是不断地减一，直到子集为 `0` 就停止。

```go
res := []int{}
subset := mask 
for subset!=0{
	res = append(res,subset)
	subset = (subset-1)&mask
}
res = append(res,0)
return res
```

好了，求子集的方法找到了，接下来就是应用了。

这里先假设`puzzle = (10110)` 那么 `(10000)` 是 `puzzle` 的第一个字母，那么就要求谜底的形式为`(10XX0)` 这里的`X` 表示这个位上可以为 `0` 或 `1` 。

注意到题目要求 `puzzle` 的第一个字母要在 `word` 中出现，这就意味着第一位是一定要满足的。那么就先求出 `puzzle` 除了第一个字母，剩下的 `b(x)` ，然后将这部分子集 `subset` 与之相与，得到的就是可能出现的状态。然后将其与 `puzzle` 的第一个字母相与即可得到这个状态的全部。

