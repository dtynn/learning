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





#### An Algorithm for Finding $\mat{A} ^{-1}$

- Row reduce the augmented matrix $\begin{bmatrix} \mat{A} & \mat{I} \end{bmatrix}$. If  $\mat{A}$ is row equivalent to $\mat{I}$, then $\begin{bmatrix} \mat{A} & \mat{I} \end{bmatrix}$ is row equivalent to $\begin{bmatrix} \mat{I} & \mat{A} ^{-1} \end{bmatrix}$. Otherwise, $\mat{A}$ does not have inverse.



#### Another View of Matrix Inversion

从 $\begin{bmatrix} \mat{A} & \mat{I} \end{bmatrix}$ 到 $\begin{bmatrix} \mat{I} & \mat{A} ^{-1} \end{bmatrix}$ 的行化简过程可以视作同时求解

$\mateq{e_1}, \mateq{e_2}, \dots , \mateq{e_n}$

当某些问题只要求算出 $\mat{A} ^{-1}$ 的特定列时, 只需要求解对应的方程即可.



### 2.3 Characterizations of Invertible Matrices

##### Theorem 8

Let $\mat{A}$ be a square $n \times n$ matrix. Then the following statements are equivalent.

That is, for a given $\mat{A}$, the statements are either all true or all false.

1. $\mat{A}$ is an invertible matrix
2. $\mat{A}$ is row equivalent to the $n \times n$ identity matrix
3. $\mat{A}$ has n pivot positions
4. The equation $\mateq{0}$ has only the trivial solution
5. The columns of $\mat{A}$ form a linearly independent set
6. The linear transformation $\vec{x} \to \mat{A} \vec{x}$ is one to one
7. The equation $\mateq{b}$ has at least one solution for each $\vec{b}$ in $\real{n}$
8. The colunms of $\mat{A}$ span $\real{n}$
9. The linear transformation $\vec{x} \to \mat{A} \vec{x}$ maps $\real{n}$ to $\real{n}$
10. There is an $n \times n$ matrix $\mat{C}$ such that $\mat{CA} = \mat{I}$
11. There is an $n \times n$ matrix $\mat{D}$ such that $\mat{AD} = \mat{I}$
12. $\mat{A} ^T$ is an invertible matrix



Let $\mat{A}$ and $\mat{B}$ be square matrices. If $\mat{AB} = \mat{I}$, then $\mat{A}$ and $\mat{B}$ are both invertible, with $\mat{B} = \mat{A} ^{-1}$ and $\mat{A} = \mat{B} ^{-1}$



- The power of the Invertible Matrix Theorem lies in the connections is provides among so many important concepts
- However, the theorem applies only to square matrices



#### Invertible Linear Transformations

A linear transformation $\mat{T}: \real{n} \to \real{n} $ is said to be **invertible** if there exists a function $\mat{S}: \real{n} \to \real{n}$ such that

$\mat{S}(\mat{T} (\vec{x})) = \vec{x}$ for all $\vec{x}$ in $\real{n}$

$\mat{T}(\mat{S} (\vec{x})) = \vec{x}$ for all $\vec{x}$ in $\real{n}$

We call $\mat{S}$ the **inverse** of $\mat{T}$ and write it as $\mat{T} ^{-1}$



##### Theorem 9

Let $\mat{T}: \real{n} \to \real{n}$ be a linear transformation and let $\mat{A}$ be the standard matrix for $\mat{T}$. Then $\mat{T}$ is invertible if and only if $\mat{A}$ is an invertible matrix. In that case, the linear transformation $\mat{S}$ given by $\mat{S}(\vec{x}) = \mat{A} ^{-1} \vec{x}$ is the unique function that satisfying the equatils above.



#### 2.4 Partitioned Matrices

#### Addition and Scalar Multiplication



#### Muliplication of Partitioned Matrices

##### Theorem10

Column-Row Expansion of $\mat{AB}$

If $\mat{A}$ is $m \times n$ and $\mat{B}$ is $n \times p$, then
$$
\mat{AB} = \begin{bmatrix} col_1(\mat{A}) & col_2(\mat{A}) & \dots & col_n(\mat{A})  \end{bmatrix} \begin{bmatrix}  row_1(\mat{B}) \\ row_2(\mat{B}) \\ \vdots \\ row_n(\mat{B}) \end{bmatrix} \\
= col_1(\mat{A}) row_1(\mat{B}) + \dots + col_n(\mat{A}) row_n(\mat{B})
$$


#### Inverse of Partitioned Matrices



#### 2.5 Matrix Factorizations 矩阵因式分解

#### The LU Factorization

- 令 $\mat{A} = \mat{L} \mat{U}$, 其中 $\mat{L}$ 是一个 $m \times m$ 对角线上元素均为1的下三角矩阵, 而$\mat{U}$ 是 $\mat{A}$ 的阶梯形, 大小为 $m \times n$ 
- $\mat{L}$ 是 不可逆的, 称为 单位下三角矩阵
- 当 $\mat{A} = \mat{LU}$ 时, $\mateq{b}$ 可以写成 $\mat{L} (\mat{U} \vec{x}) = \vec{b}$. 将 $\mat{U} \vec{x}$ 写成 $\vec{y}$, 则可以通过先后求解下列方程来求得 $\vec{x}$

$$
\mat{L} \vec{y} = \vec{b}
\\
\mat{U} \vec{x} = \vec{y}
$$



#### An LU Factorization Algorithm

1. Reduce $\mat{A}$ to an echelon form $\mat{U}$ by a sequence of row replacement operations, if possible.
2. Place entries in $\mat{L}$ such that the *same sequence of row operations* reduces $\mat{L}$ to $\mat{I}$



#### A Matrix Factorization in Electrical Engineering





### 2.6 The Leontief Input-Output Model

$$
\left \{
\begin{array}{c} amout
\\ produced
\\ \vec{x}
\end{array}
\right \}
=
\left \{
\begin{array}{c}
intermeiate
\\
demand
\end{array}
\right \}
+
\left \{
\begin{array}{c}
final
\\
demand
\\
\vec{d}
\end{array}
\right \}
$$

##### The Leontief Input-Output Model, Or Production Equation

$$
\vec{x} = \mat{C} \vec{x} + \vec{d}
$$

can also be written as $\mat{I} \vec{x} - \mat{C} \vec{x} = ( \mat{I} - \mat{C} ) \vec{x} = \vec{d}$



##### Theorem 11

Let $\mat{C}$ be the consumption matrix for an economy, and let $\vec{d}$ be the final demand. If $\mat{C}$ and $\vec{d}$ have nonnegative entries and if each column sum of $\mat{C}$ is less than 1, then $( \mat{I} - \mat{C} ) ^{-1}$ exists and the production vector
$$
\vec{x} = ( \mat{I} - \mat{C} ) ^{-1} \vec{d}
$$
has nonnegative entries and is the unique solution of
$$
\vec{c} = \mat{C} \vec{x} + \vec{d}
$$


#### A Formula for $(\mat{I} - \mat{C}) ^{-1}$



#### The Economic Importance of Entries in $( \mat{I} - \mat{C} ) ^{-1}$



### 2.7 Applications To Computer Graphics





### 2.8 Subspaces of $\real{n}$



##### Definition

A **subspace** of $\real{n}$ is any set $\mat{H}$ in $\real{n}$ that has three properties:

1. The zero vector is in $\mat{H}$
2. For each $\vec{u}$ and $\vec{v}$ in $\mat{H}$, the sum $\vec{u} + \vec{v}$ is in $\mat{H}$.
3. For each $\vec{u}$ in $\mat{H}$ and each scalar c, the vector $c \vec{u}$ is in $\mat{H}$



#### Column Space and Null Space of a matrix

- The **column space** of a mtraix $\mat{A}$ is the set Col $\mat{A}$ of all linear combinations of the columns of $\mat{A}$
- The **nul space** of a matrix $\mat{A}$ is the set Nul $\mat{A}$ of all solutions of the homogeneous equation $\mateq{0}$



##### Theorem 12

The null space of an $m \times n$ matrix $\mat{A}$ is a subspace of $\real{n}$. Equivalently, the set of all solutions of a system $\mateq{0}$ of m homogeneous linear equations in n unknowns is a subspace of $\real{n}$



#### Basis for a Subspace

- A **basis** for a subspace $\mat{H}$ of $\real{n}$ is a linearly independent set in $\mat{H}$ that spans $\mat{H}$.
- The set $\{ \vectors{e}{n} \}$ is called the **standard basis** for $\real{n}$



##### Theorem 13

The pivot columns of a matrix $\mat{A}$ form a basis for the column space of $\mat{A}$



### 2.9 Dimension and Rank

#### Coordinate Systems

##### Definition

Suppose the set $ß = \{ \vectors{b}{p} \}$ is a  basis for a subspace $\mat{H}$. For each $\vec{x}$ in $\mat{H}$, the **coordinates of x relative the basis ß** are the weights $\weights{c}{p}$ such that $\vec{x} = \lcomb{c}{b}{p}$, and the vector in $\real{p}$
$$
\begin{bmatrix} \vec{x} \end{bmatrix}_ß = \begin{bmatrix} c_1 \\ c_2 \\ \vdots \\ c_p \end{bmatrix}
$$
is called the **coordinate vector of x (relative to ß)** or the **ß-coordinate vector of x**.



#### The Dimension of a Subspace

- The **dimension** of a nonzero subspace $\mat{H}$, denoted by $dim \mat{H}$, is the number of vectors in any basis for $\mat{H}$. The dimension of the zero subspace $\{ \vec{0} \}$ is defined to be zero.

- The **rank** of a matrix $\mat{A}$, denoted by rank $\mat{A}$, is the dimension of the column space of $\mat{A}$

  Since the pivot columns of $\mat{A}$ form a basis for Col $\mat{A}$, the rank of $\mat{A}$ is just the number of pivot columns in $\mat{A}$



##### Theorem 14

The Rank Theorem

If a matrix $\mat{A}$ has n columns, then rank $\mat{A}$ + dim Nul $\mat{A}$ = n



##### Theorem 15

The Basis Theorem

Let $\mat{H}$ be a p-dimension subspace of $\real{n}$. Any linear independent set of exactly p elements in $\mat{H}$ is automatically a basis for $\mat{H}$. Also, any set of p elements of $\mat{H}$ tha spans $\mat{H}$ is automatically a basis for $\mat{H}$





#### Rank and the Invertible Matrix Theorem

##### Theorem

The invertible Matrix Theorem (continued)

Let $\mat{A}$ be an $n \times n$ matrix. Then the following statements are each equivalent to the statement that $\mat{A}$ is an invertible matrix.

1. The columns of $\mat{A}$ form a basis of $\real{n}$
2. Col $\mat{A}$ = $\real{n}$
3. dim Col $\mat{A}$ = n
4. rank $\mat{A}$ = n
5. Nul $\mat{A}$ = $\{ \vec{0} \}$
6. dim Nul $\mat{A}$ = 0