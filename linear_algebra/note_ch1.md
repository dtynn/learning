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

## Chapter 1 Linear Equation in Linear Algebra

### 1.1 System of Linear Equations

-   linear equation 线性方程
-   system of linear equation 线性方程组
-   solution 线性方程组的解
-   solution set: the set of all possible solutions 线性方程组的解集
-   equivalent 等价
-   consistent: one or infinitely many solutions 相容的
-   inconsistent: no solution 不相容的

#### Matrix Notation 矩阵记号

-   coefficient matrix 系数矩阵
-   augmented matrix 增广矩阵
-   $m \times n$ matrix

#### Solving a Linear System

basic strategy: replace one system with an equivalent system that is easier to solve.

##### three basic operations (elementary row operations)
1.  Replacement: replace one equation by the sum of itself and a multiple of another equation
2.  Interchange: interchange two equations
3.  Scaling: multiply all the terms in an equation by a nonzero constant

-   row equivalent
-   row operations are revesible

#### Existence and Uniqueness Questions

##### two fundamental questions about a linear system
1.  does at least one solution exists
2.  is the solution unique

### 1.2 Row Reduction And Echelon Forms 行化简和阶梯形矩阵
-   leading entry: the leftmost nonzero entry in a nonzeron row

##### echelon form:
1.  all nonzero rows are above any rows of all zeros
2.  each leading entry of a row is in a column to the right of the leading entry of the row above it
3.  all entries in a column below a leading entry are zeros

reduced echelon form:

4.  the leading entry in each nonzero row is 1
5.  each leading 1 is the only nonzero entry in its column

##### Theorem 1 uniqueness of the reduced echelon form
`each matrix is equaivalent to one and only one reduced echelon form`

#### Pivot Positions 主元位置
the leading entries are always in the same positions in any echelon form obtained from a given matrix

```
A pivot position in a matrix A is a localtion in A that corresponds to leading 1 in the reduced echelon form of A. A pivot column is a column of A that contains a pivot position.
```

#### Solution Of Linear Systems

-   basic variables / leading variables: variables in pivot column
-   free variables: other variables

#### Parametric Descriptions of Solution Sets 解集的参数表示

#### Back-Substitution 回代

#### Existence and Uniqueness Questions
##### Theorem 2
A linear system is consistent if and only if the rightmost column of the augmented matrix is not a pivot column

that is, if and only if an echelon form of the augmented matrix has no row of the form

$\begin{bmatrix}\dupcols{0} & b\end{bmatrix}$ with b *nonzero*



If a linear system is consistent, then the solution set contains either

1.  a unique solution, when there are no free variables, or
2. infinitely many solutions, when there is at least one free variables.

### 1.3 Vector Equations

#### Vector in $\real{2}$

-   column vector / vector
-   sum of vector $\vec{u}$ and $\vec{v}$
-   scalar multiple of vector $\vec{u}$ by c

#### Geometric Descriptions of $\real{2}$

##### Parallelogram Rule for Addition

#### Linear Combinations

- Given vectors $\vectors{v}{p}$ in $\real{n}$ and scalars $\weights{c}{p}$ , the vector $\vec{y}$ deinfe by

  $\vec{y} = \lcomb{c}{v}{p}$

  is called linear combination of $\vectors{v}{p}$ and weights $\weights{c}{p}$

- a vector equation $\lcomb{x}{a}{n} = \vec{b}$

  has the same solution set as the linear system whose augmented matrix is

  $\begin{bmatrix}\colvecs{a}{n} & \vec{b}\end{bmatrix}$

-   $\vec{b}​$ can be generated by a linear combination of $\vectors{a}{n}​$ if and only if there exsit a solution to the linear system corresponding to the matrix above

-   If $\vectors{v}{p}$ are in $\real{n}$, then the set of all linear combination of $\vectors{v}{p}$ is denoted by Span{$\vectors{v}{p}$} and is called the subset of $\real{n}$ spanned (or generated) by $\vectors{v}{p}$

    That is, Span{$\vectors{v}{p}$} is the collection of all vectors that can be written the form

    $\lcomb{c}{v}{p}$

    with $\weights{c}{p}$ scalars

#### Linear Combinations in Applications

### 1.4 The Matrix Equation $\mateq{b}$

##### Theorem 3

1. the matrix equation: $\mateq{b}$
2. the vector equation: $\lcomb{x}{a}{n} = \vec{b}$
3. the system of linear equation whose augmented matrix is $\begin{bmatrix} \colvecs{a}{n} & \vec{b}\end{bmatrix}$

has the same solution sets

#### Existence of Solutions

##### Theorem 4

Let $\mat{A}$ be an  $m \times n$ matrix. Then the following statements are logically equivalent. That is, for particular $\mat{A}$, either they are all true statements or they are all false.

1.  For each $\vec{b}$ in $\real{m}$, the equation $\mateq{b}$ has a solution
2.  Each $\vec{b}$  in  $\real{m}$ is a linear combination of the columns of  $\mat{A}$
3.  The columns of $\mat{A}$ span $\real{m}$
4.  $\mat{A}$ has a pivot position in every row

##### Row-Vector Rule of Computing $\mat{A}\vec{x}$

If the product $\mat{A}\vec{x}$ is defined, then the ith entry in $\mat{A}\vec{x}$ is the sum of the products of corresponding entries from row i of $\mat{A}$ and from the vector $\vec{x}$

- identity matrix $\mat{I}$ 单位矩阵

#### Properties of the Matrix-Vector Product $\mat{A}\vec{x}$

##### Theorem 5

If $\mat{A}$ is an $m \times n$ matrix, $\vec{u}$ and $\vec{b}$ are vectors in $\real{n}$, and c is a scalar, then
1.  $\mat{A} ( \vec{u} + \vec{v} ) = \mat{A} \vec{u} + \mat{A} \vec{v}$
2.  $\mat{A} ( c \vec{u} ) = c  ( \mat{A} \vec{u} ) $

### 1.5 Solutions Sets of Linear Systems

#### Homogeneous Linear System 齐次线性方程组

-   $\mateq{0}$ where $\mat{A}$ is an $m \times n$  matrix and $\vec{0}$ is the zero vector in $\real{m}$
-   trivial solution 平凡解: $\vec{x} = \vec{0}$
-   nontrivial solution 非平凡解: a nonzero vector that satisfies $\mateq{0}$

The homogeneous equation $\mateq{0}$ has a nontrivial solution if and only if the equation has at least one free variable.

#### Parametric Vector Form 参数向量形式

#### Solution of Nonhomogeneous Systems

##### Theorem 6

Suppose the equation $\mateq{b}$ is consistent for some given $\vec{b}$, and let $\vec{p}$ be a solution. Then the solution set of $\mateq{b}$ is the set of all vectors of the form $\vec{w} = \vec{p} + \vec{v}_h$ , where $\vec{v}_h$ is any solution of the homogeneous equation $\mateq{0}$

##### Writing a Solution Set (of a Consistent System) in Parametric Vector From

1.  row reduce the augmented matrix to reduced echelon form.
2.  express each basic variable in terms of any free variables appearing in an equation.
3.  write a typical solution $\vec{x}$ as a vector whose entries depend on the free variables, if any.
4.  decompose $\vec{x}$ into a linear combination of vectors (with numeric entries) using the free variables as parameters.

### 1.6 Appications of Linear Systems

### 1.7 Linear Independence

An indexed set of vectors {$\vectors{v}{p}$} in $\real{n}$ is said to be linearly independent if the vector equation

$\lcomb{x}{v}{p} = \vec{0}$

has only the trivial solution.

The set is said to be linearly dependent if there exists weights $\weights{c}{p}$ , not all zero, such that

$\lcomb{c}{v}{p} = \vec{0}$



-   linear dependence relation

#### Linear Independence of Matrix Columns

The columns of a matrix $\mat{A}$ are linearly independent if and only if the equation $\mateq{0}$ has only the trivial solution

#### Set of One or Two Vectors

-   a set containing only one vector — $\vec{v}$  —  is linearly independent if and only if $\vec{v}$ is not the zero vector
-   the zero vector is linearly dependent

A set of tow vectors {$\vec{v}_1, \vec{v}_2 $} is linearly dependent if at least one of the vectors is a multiple of the other. The set is linearly independent if and only if neither of the vectors is a multiple of the other.

#### Set of Two or More Vectors

##### Theorem 7

Characterization of Linearly Dependent Sets

An indexed set S = {$\vectors{v}{p}$} of two or more vectors is linearly dependent if and only if at least one of the vectors in S is a linear combination of the others. In fact, if S is linearly dependent and $\vec{v}_1$ is not $\vec{0}$, then some ${\vec{v}_j}$  (with j > 1) is a linear combination of the preceding vectors $\vectors{v}{j-1}$

##### Theorem 8

If a set contains more vectors than there are entries in each vector, than the set is linearly dependent.

That is, any set {$\vectors{v}{p}$} in $\real{n}$ is linearly dependent if p > n.

##### Theorem 9

If a set S = {$\vectors{v}{p}$} in $\real{n}$ contains the zero vector, then the set is linearly dependent

### 1.8 Introduction to Linear Transformations 线性变换

$\textit{T} : \real{n} \rightarrow \real{m}$

-   $\textit{T}$ : transformation (or function, or mapping)
-   $\real{n}$:  domain of $\textit{T}$
-   $\real{m}$:  codomain of $\textit{T}$
-   $\ltran{x}$ : the image of $\vec{x}$
-   range of $\textit{T}$ : the set of all images $\ltran{x}$

#### Matrix Transformations

-   projection transformation 投影变换
-   shear transformation 剪切变换

#### Linear Transformations
A transformation (or mapping) $\textit{T}$ is linear if:

1.  $\ltran{u + v} = \ltran{u} + \ltran{v}$ for all $\vec{u}$, $\vec{v}$  in the domain of $\textit{T}$
2.  $\textit{T} ( c \vec{u} ) = c \ltran{u}$ for all scalars c and all $\vec{u}$ in the domain of $\textit{T}$

every matrix transformation is a linear transformation

if $\textit{T}$ is a linear transformation, then

1.  $\ltran{0} = \vec{0}$
2.  $\textit{T} ( c \vec{u} + d \vec{v} ) = c \ltran{u} + d \ltran{v}$

for all vectors $\vec{u}$, $\vec{v}$ in the domain of $\textit{T}$ and all scalars c, d.

-   $\textit{T} ( \lcomb{c}{v}{p} ) = c_1 \ltran{v_1} + c_2 \ltran{v_2} + \dots + c_p \ltran{v_p}$

    superposition principle 叠加原理

### 1.9 The Matrix of a Linear Transformation

##### Theorem 10

Let $\textit{T} : \real{n} \rightarrow \real{m}$ be a linear transformation. Then there exists a unique matrix $\mat{A}$ such that
$\ltran{x} = \mat{A} \vec{x}$ for all $\vec{x}$ in $\real{n}$

in fact, $\mat{A}$ is the $m \times n$ mtraix whose *j*th column is the vector $\ltran{e_j}$ where is $\vec{e_j}$ the *j*th column of the identity matrix in $\real{n}$

$\mat{A} = \begin{bmatrix} \ltran{e_1} & \dots & \ltran{e_n} \end{bmatrix}$  is called the standard matrix for the linear transfomation $\textit{T}$

#### Geometrix Linear Transformations of $\real{2}$

#### Existence and Uniqueness Questions

1.  A mapping $\textit{T} : \real{n} \rightarrow \real{m}$ is said to be onto $\real{m}$ if each $\vec{b}$ in  $\real{m}$ is the image of at least one $\vec{x}$ in  $\real{n}$ 满射

    for each $\vec{b}$ in  the codomain  $\real{m}$, there exists at least one solution of $\ltran{x} = \vec{b}$.

2.  A mapping $\textit{T} : \real{n} \rightarrow \real{m}$ is said to be one-to-one if each $\vec{b}$ in  $\real{m}$ is the image of at most one $\vec{x}$ in  $\real{n}$ 一对一映射, 单射

     for each $\vec{b}$ in  $\real{m}$, the equation $\ltran{x} = \vec{b}$ has either a unique solution or none at all.


##### Theorem 11

Let $\textit{T} : \real{n} \rightarrow \real{m}$ be a linear transformation. Then $\textit{T}$ is one-to-one if and only if the equation $\ltran{x} = \vec{0}$ has only the trivial solution.

##### Theorem 12

Let $\textit{T} : \real{n} \rightarrow \real{m}$ be a linear transformation and let $\mat{A}$ be the standard matrix for $\textit{T}$. Then:

1.  $\textit{T}$ maps  $\real{n}$ onto  $\real{m}$ if and only if the columns of $\mat{A}$ span  $\real{m}$
2.  $\textit{T}$ is one-to-one if and only if the columns of $\mat{A}$ are linearly independent.





### Supplementary Exercises

1. (s), (y)



### 重点关注

- In some cases, it is possible for four vectors to span $\real{5}$
  - 一个 $5 \times 4$ 的矩阵, 对其增广矩阵进行阶梯化, 不可避免会出现 $\begin{bmatrix} 0 & 0 & 0 & 0 & b \end{bmatrix}$  的行