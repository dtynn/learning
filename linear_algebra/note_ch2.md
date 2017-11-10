###### formula defines

real $\newcommand{\real}[1]{\mathbb{R}^{#1}}$

vecsym $\newcommand{\vec}[1]{\mathbf{#1}}$

matsym $\newcommand{\mat}[1]{\boldsymbol{#1}}$

vectors $\newcommand{\vectors}[2]{\vec{#1}_1, \vec{#1}_2, \dots , \vec{#1}_{#2} }$

weights $\newcommand{\weights}[2]{{#1}_1, {#1}_2, \dots , {#1}_{#2} }$

colvecs $\newcommand{\colvecs}[2]{\vec{#1}_1 &  \vec{#1}_2  & \dots & \vec{#1}_{#2}}$

dupcols $\newcommand{\dupcols}[1]{{#1} &  {#1}  & \dots & {#1}}$

lcomb(linear combination) $\newcommand{\lcomb}[3]{{#1}_1\vec{#2}_1 + {#1}_2\vec{#2}_2 + \dots + {#1}_{#3}\vec{#2}_{#3}  }$

mateq(matrix equation) $\newcommand{\mateq}[1]{\mat{A} \vec{x} = \vec{#1}}$

ltran(linear transformation) $\newcommand{\ltran}[1]{\textit{T} ( \vec{#1} ) }$

###### formula defines end

## Chapter2 Matrix Algebra

### 2.1 Matrix Operations

- entry $a_{ij}$
- colums $\vectors{a}{n}$
- matrix $\mat{A} = \begin{bmatrix} \colvecs{a}{n} \end{bmatrix}$
- **diagonal entries** 对角线元素: $a_{11}, a_{22}, a_{33}, …$ , they form the **main diagonal** of A
- **diagonal matrix**  a n*n matrix whose nondiagonal entries are all zero. Example $\mat{I}_n$
- **zero matrix**

#### Sums and Scalar Multiples

- two matrices are **equal**:
  1. have the same size
  2. corresponding colums are equal
- **sum** $\mat{A} + \mat{B}$  only when A and B are the same size



##### Theorem 1

Let $\mat{A}$ and $\mat{B}$ and $\mat{C}$ be matrices of the same size, and let r and s be scalars

1. $\mat{A} + \mat{B} = \mat{B} + \mat{A}  $
2. $(\mat{A} + \mat{B} ) + \mat{C} = \mat{A} + ( \mat{B} + \mat{C} )$
3. $\mat{A} + 0 = \mat{A}$
4. $r ( \mat{A} + \mat{B} ) = r \mat{A} + r \mat{B}$
5. $ (r + s) \mat{A} = r \mat{A} + s \mat{A}$
6. $r ( s \mat{A} ) = ( rs ) \mat{A}$



#### Matrix Multiplication

- $\mat{A} ( \mat{B} \vec{x} ) = ( \mat{A} \mat{B} ) \vec{x}$

###### definition

If $\mat{A}$ is an $m \times n$ matrix, and if $\mat{B}$ is an $n \times p$ matrix with columns $\vectors{b}{p}$ , then the product $\mat{AB}$ is the $m \times p$ matrix whose columns are $\mat{A}\vec{b}_1, … , \mat{A}\vec{b}_p$ . That is
$$
\mat{AB} = \mat{A} \begin{bmatrix} \colvecs{b}{p} \end{bmatrix} = \begin{bmatrix} \mat{A}\vec{b}_1 & \mat{A}\vec{b}_2 & \dots & \mat{A}\vec{b}_p \end{bmatrix}
$$
Multiplication of matrices corresponds to composition of linear transformations.

矩阵乘法对应线性变换的复合



Each column of $\mat{AB}$ is a linear combination of the columns of $\mat{A}$ using weights from the corresponding column of $\mat{B}$



##### Row-Column Rule for Computing $\mat{AB}$

If the product $\mat{AB}$ is defined, the the entry in row *i* and column *j* of $\mat{AB}$ is the sum of the products of corresponging entries from row *i* of $\mat{A}$ and column *j* of $\mat{B}$. if $( \mat{AB} )_{ij}$ denotes the (i, j)-entry in $\mat{AB}$, and if $\mat{A}$ is an $m \times n$ matrix, then
$$
(\mat{AB})_{ij} = a_{i1} b_{1j} + a_{i2} b_{2j} + \dots + a_{in} b_{nj}
$$

$$
row_i ( \mat{AB} ) = row_i( \mat{A} ) \cdot  \mat{B}
$$


#### Properties of Matrix Multiplication

##### Theorem 2

Let $\mat{A}$ be an $m \times n$ matrix, and let $\mat{B}$ and $\mat{C}$ have sizes for which the indecated sums and products are defined.

1. $\mat{A ( \mat{BC} )} = ( \mat{AB} ) \mat{C}$  associative law of mutiplication  乘法结合律
2. $\mat{A} ( \mat{B} + \mat{C} ) = \mat{AB} + \mat{AC}$ left distributive law 左分配率
3. $( \mat{B} + \mat{C} ) \mat{A} = \mat{BA} + \mat{CA}$ right distibutive law 右分配率
4. $r ( \mat{AB} ) = ( r \mat{A} ) \mat{B} = \mat{A} ( r \mat{B} )$ for any scalar r
5. $\mat{I}_m \mat{A} = \mat{A} = \mat{A} \mat{I}_n$ identity for matrix multiplication



###### Warnings:

1. In general, $\mat{AB} \ne \mat{BA}$
2. The cancellation laws do not hold for matrix multiplication. That is, if $\mat{AB} = \mat{AC}$, then it is not true in general that $\mat{B} = \mat{C}$
3. If a product $\mat{AB}$ is the zero matrix, you cannot conclude in general the either $\mat{A} = \mat{0}$ or $\mat{B} = 0$





#### Powers of a Matrix

- for $n \times n$ matrix $\mat{A}$ and positive integer k,  $\mat{A}^k$ denotes the product of k copies of $\mat{A}$



#### The Transpose of a Matrix

- the **transpose** of an $m \times n$ matrix $\mat{A}$ is the $n \times m$ matrix, denoted by $\mat{A}^T$, whose columns are formed from the corresponding rows of $\mat{A}$



##### Theorem 3

Let $\mat{A}$ and $\mat{B}$ denote matrices whose sizes are appropriate for the following sums and products.

1. $( \mat{A} ^T ) ^T = \mat{A}$
2. $ ( \mat{A} ^T + \mat{B} ^T ) = ( \mat{A} + \mat{B} ) ^T $
3. $ ( r \mat{A} ) ^T = r ( \mat{A} ^T ) $ for any scalar r
4. $ ( \mat{AB} ) ^T = \mat{B} ^T \mat{A} ^T $



- The transpose of a product of matrices equals the product of their transpose in the *reverse* order



### 2.2 The Inverse of a Matrix 矩阵的逆

- **invertible** for an $n \times n$ matrix $\mat{A}$, there is an $n \times n$ matrix $\mat{C}$ such that

  $\mat{CA} = \mat{I}$ and $\mat{AC} = \mat{I}$

  in this case , $\mat{C}$ is an inverse of $\mat{A}$, and $\mat{C}$ is uniquely determined by $\mat{A}$

  this unique inverse is denoted by $\mat{A}^{-1}$ so that

  $\mat{A}^{-1} \mat{A} = \mat{I}$ and $\mat{A} \mat{A} ^{-1} = \mat{I}$

- **singular matrix** 奇异矩阵: not invertible matrix

- **nonsingular matrix** 非奇异矩阵:  invertible matrix



##### Therom 4

Let $\mat{A} = \begin{bmatrix} a & b \\ c & d \end{bmatrix} $, if $ad - bc \ne 0$, then $\mat{A}$ is invertible and
$$
\mat{A}^{-1} = \dfrac{1}{ad-bc} \begin{bmatrix} d & -b \\ -c & a \end{bmatrix}
$$
if $ad - bc = 0$, then $\mat{A}$ is not invertible

The quantity $ad - bc$ is called the determinat of $\mat{A}$ and we write

$\text{det} \mat{A} = ad - bc$



##### Therom 5

If $\mat{A}$ is an invertible $n \times n$ matrix , then for each $\vec{b}$ in $\real{n}$, the equation $\mateq{b}$ has the unique solution $\vec{x} = \mat{A} ^{-1} \vec{b}$



##### Theorem 6

1. If $\mat{A}$ is an invertible matrix, then $\mat{A}^{-1}$ is invertible and

   $ ( \mat{A} ^{-1} ) ^{-1} = \mat{A} $

2. if $\mat{A}$ and $\mat{B}$ are $n \times n$ invetible matrices, then so is $\mat{AB}$, and inverse of $\mat{AB}$ is the product of the inverses of $\mat{A}$ and $\mat{B}$ in the reverse order. That is

   $( \mat{AB} ) ^{-1} = \mat{B} ^{-1} \mat{A} ^{-1}$

3.  If $\mat{A}$ is an invertible matrix, then so is $\mat{A} ^T$, and the inverse of $\mat{A} ^T$ is the trnaspose of $\mat{A} ^{-1}$. That is

   $( \mat{A} ^T) ^{-1} = ( \mat{A} ^{-1} ) ^T$



The product of $n \times n$ invertible matrices is invertible, and the inverse is the product of their inverses in the reverse order.



#### Elementary Matrices 初等矩阵

- An **elementary matrix** is one that is obtained by performing a single elementary row operation on an identity matrix.



If an elementary row operation is performed on an $m \times n$ matrix $\mat{A}$, the resulting matrix can be written as $\mat{EA}$, where the $m \times m$ matrix $\mat{E}$ is created by performing the same row operation on $\mat{I}_m$



Each elementary matrix $\mat{E}$ is invertible. The inverse of $\mat{E}$ is the elementary matrix of the same type that transforms $\mat{E}$ back into $\mat{I}$



##### Theorem 7

An $n \times n$ matrix $\mat{A}$ is invertible if and only if $\mat{A}$ is row equivalent to $\mat{I}_n$, and in this case, any sequence of elementary row operations that reduces $\mat{A}$ to $\mat{I}_n$ also transforms $\mat{I}_n$ into $\mat{A} ^{-1}$

