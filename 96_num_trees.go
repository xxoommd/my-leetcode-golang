package leetcode

func numTrees(n int) int {
	return numTreesCatalan(n)
}

// 卡特兰数
//  g(n)的值就被称为卡特兰数
//  卡特兰数更便捷的定义：
//  C(0) = 1, C(n+1) = (2 * (2 * n + 1) / n + 2) * C(n)
func numTreesCatalan(n int) int {
	var c int64 = 1
	for i := int64(0); i < int64(n); i++ {
		c = c * 2 * (2*i + 1) / (i + 2)
	}

	return int(c)
}

// 动态规划
// * 函数g和f
// 	- g(n)为长度为n的不同二叉搜索树个数（即求解函数）
// 	- f(i, n)为以i为根节点的二叉搜索树（且不重复）
//
// * g(n)可以从f(i, n)得到，而f(i, n)又依赖于g(n):
// 	- g(n) = ∑f(i, n)
// 	- f(i, n) = g(i-1) * g(n-i) (笛卡尔积)
//
// * 特殊情况:
// 	- g(0) = 1
//  - g(1) = 1
func numTreesDP(n int) int {
	g := make([]int, n+1)
	g[0], g[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			g[i] += g[j-1] * g[i-j]
		}
	}

	return g[n]
}
