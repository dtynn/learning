#### Quick sort

1. divide: 分成两个子列, 分别小于和大于一个选定的元素
2. conquer: 递归地划分并排序
3. combine: trivial



##### worst-case time

input (reverse) sorted
$$
\begin{align}
T(n) & = T(0) + T(n-1) + \Theta(n) \\
& = \Theta(1) + T(n-1) + \Theta(n) \\
& = T(n-1)+ \Theta(n) \\
& = \Theta(n^2)
\end{align}
$$
由递归树可见, 是一颗极端不平衡的递归树



##### best-case time

每次划分都可以二等分
$$
\begin{align}
T(n) & = 2T(n/2) + \Theta(n) \\
& = \Theta(n\lg n)
\end{align}
$$


###### 如果每次划分结果都是 1: 9

$$
\begin{align}
T(n) & = T(n/10) + T(9n/10) + \Theta(n) \\
\end{align}
$$

由递归树, 最长路径长度为 $\log_{10/9} n$, 最短路径为 $\log_{10} n$

每一层的叶子和, 在 最短路径前为 $cn$, 其后则均 $\le cn$

由此可知 $cn\log_{10} n \le T(n) \le cn \cdot \log_{10/9}n + \Theta(n) \le \Theta(n\lg n)$



###### 如果划分时, 好的情况和坏的情况间隔发生

$$
L(n) = 2U(n/2) + \Theta(n) \\
U(n) = L(n-1) + \Theta(n) 
$$

可得
$$
\begin{align}
L(n) & = 2[L(n/2 - 1) + \Theta(n/2)] + \Theta(n) \\
& = 2L(n/2 - 1) + \Theta(n) \\
& = \Theta(n \lg n)
\end{align}
$$

##### Randomized quicksort

当我们不依赖于输入的顺序时, 通常都可以得到好的情况

使用一个随机主元,  此时没有一个特定的输入会导致最差情形, 此时仅由随机数生成器决定结果