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

## Chapter 3 Determinants

### 3.1 Introduction to Determinants 行列式

##### Definition

For $n \ge 2$, the **determinant** of an $n \times n$ matrix $\mat{A} = \begin{bmatrix} a_{ij} \end{bmatrix}$  in symbols:
$$
\text{det} \mat{A} = a_{11} \text{det} \mat{A}_{11} - a_{12} \text{det} \mat{A}_{12} + \dots + (-1)^{1+n} a_{1n} \text{det} \mat{A}_{1n}
\\
= \sum_{j=1}^{n} (-1) ^{1+j} a_{1j} \text{det} \mat{A}_{1j}
$$


Given $\mat{A} = \begin{bmatrix} a_{ij} \end{bmatrix}$, the **(i, j)-cofactor** of $\mat{A}$ is the number $C_{ij}$ givet by
$$
C_{ij} = (-1) ^{i+j} \text{det} \mat{A}_{ij}
$$


##### Theorem 1

The determinant of an $n \times n$ matrix $\mat{A}$ can be computed by a cofactor expansion across any row or down any column. The expansion across the *i*th row using the cofactors is
$$
\text{det} \mat{A} = a_{i1}C_{i1}+ a_{i2}C_{i2} + \dots + a_{in}C_{in}
$$
The cofactor expansion down the *j*th column is
$$
\text{det} \mat{A} = a_{1j}C_{1j} + a_{2j}C_{2j} + \dots + a_{nj}C_{nj}
$$


##### Theorem 2

If $\mat{A}$ is a triangular matrix, the $\text{det} \mat{A}$ is the product of the entries on the main diagonal of $\mat{A}$



#### 3.2 Properties of Determinants

