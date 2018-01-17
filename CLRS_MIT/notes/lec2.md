### 渐进符号

#### $O$ -notation (upper bounds)

$f(n) = O({g(n)})$:

exist constants $c > 0, n_0 > 0$ 

such that $0 \le f(n) \le cg(n)$

for all $n \ge n_0$



理解类似小于等于



##### Macro substitution

1. $f(n) = n^3 + O(n^2)$

   means

   $f(n) = n^3 + h(n)$

   for some 

   $h(n) \in O(n^2)$

2. $n^2 + O(n) = O (n^2)$

   means for any $f(n) \in O(n)$

   there exists $h(n) \in O(n^2)$

   such that $n^2 + f(n) = h(n)$

   ​

解读: 等号代表 "is", 左侧表明任意性 ( any ), 右侧表明存在性 ( exists )



#### $\Omega$ - notation (lower bounds)

$f(n) = \Omega(g(n))$

exists constants $c > 0, n_0 > 0$

such that $0 \le cg(n) \le f(n)$

for all $n \ge n_0$



理解类似大于等于



#### $\Theta$ - notation (tight bounds)

$\Theta(g(n)) = O(g(n)) \cap \Omega(g(n))$



#### $o$ - notation and $\omega$ - notation

##### $o$ - notaion

$f(n) = O({g(n)})$:

for any constant $c > 0$

there is a constant $n_0 > 0$ 

such that $0 \le f(n) < cg(n)$

for all $n \ge n_0$



理解类似小于



##### $\omega$ - notaion

$f(n) = \Omega(g(n))$

for any constant $c > 0$

there is a constant $n_0 > 0$

such that $0 \le cg(n) < f(n)$

for all $n \ge n_0$



理解类似大于



### 解递归式

#### Substitution method

most general method:

1. **Guess** the form of the solution
2. **Verify** by induction
3. **Solve** for constants



#### Recursion-tree method



#### The master method 主方法

applies to recurrences of the form

$T(n) = aT(n/b) + f(n)$

where $a \ge 1, b \ge 1$ and $f$ is asymptotically positive



比较 $f(n)$ 与 $n^{\log_b a}$

$f(n) = O(n^{\log_b a - \epsilon})$ for some constant $\epsilon > 0$

1. $f(n)$ 增长慢, 则结果为 $T(n) = \Theta(n^{\log_ba})$
2. 两者增长的比率接近, 则结果为 $T(n) = \Theta(n^{\log_b a} \lg^{k+1} n)$
3. $f(n)$ 增长快, 则结果为 $T(n) = \Theta(f(n))$

