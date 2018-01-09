#### 分治法 Divide and Conquer

1. Divide into sub-problems
2. Conquer each sub-problem recursively
3. Combine the solutions



##### Merge sort

$T(n) = 2T(n/2) + \Theta(n) = \Theta(n\lg n)$



##### Binary search

$T(n) = T(n/2) + \Theta(1) = \Theta(\lg n)$



##### Powering a number

###### native algorithm

$T(n) = \Theta(n)$



###### divide-and-conquer

$$
a^n = \left\{ 
\begin{aligned}
\begin{aligned}
& a^{n/2} \cdot a^{n/2} \\
& a^{(n-1) / 2} \cdot a^{(n-1) / 2} \cdot a \\
\end{aligned}
\begin{aligned}
& \text{if n is even} \\
& \text{if n is odd} \\
\end{aligned}
\end{aligned}
\right.
$$



$T(n) = T(n/2) + \Theta(1) = \Theta(\lg n)$



##### Fibonacci numbers

$$
F_n = 
\left\{
\begin{aligned}

\begin{aligned}
& 0 \\
& 1 \\
& F_{n-1} + F_{n-2} \\
\end{aligned}

\begin{aligned}
& \text{if n = 0;} \\
& \text{if n = 1;} \\
& \text{if n } \ge \text{2;} \\
\end{aligned}

\end{aligned}
\right.
$$



###### native recursive algorithm

exponential time 指数时间

$\Omega(\phi^n)$

where $\phi = (1 + \sqrt{5}) / 2$



###### bottom-up

$\Theta(n)$



###### naive recursive squaring

$F_n = \phi / \sqrt{5}$ rounded to the nearest integer.

$\Theta(\lg n)$

unreliable on real machine as floting-point arithmetic is prone to round-off errors



###### recursive squaring

Theorem

$\begin{bmatrix} F_{n+1} & F_n \\ F_n & F_{n-1} \end{bmatrix} = \begin{bmatrix} 1 & 1 \\ 1 & 0 \end{bmatrix} ^ n$

$\Theta(\lg n)$



##### Matrix mutiplication

###### standard algorithm

$\Theta(n^3)$



###### divide-and-conquer algorithm

思路: 矩阵分块

$T(n) = 8 T(n/2) + \Theta(n^2) = \Theta(n^3)$



###### Strassen's algorithm

减少一次递归乘法, 即系数8

$T(n) = 7 T(n/2) + \Theta(n^2) = \Theta(n^{\lg 7})$



#### VLSI layout

