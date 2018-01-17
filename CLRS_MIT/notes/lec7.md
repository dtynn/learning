#### Hashing I

##### Symbol-table problem



##### Direct-access tabel

problem: keys range



##### Hash function



##### Resolving collisions by chaining

worst case:

- every key hashes to the same slot
- access time $\Theta(n)$



average case:

- expected unsuccessful serach time: $\Theta(1+\alpha)$

- expected search time: $\Theta(1)$

  if $\alpha = O(1)$ or equivalently, if $n = O(m)$



##### Choosing a hash function

- distribute the keys uniformly into the slots
- regularity in the key distribution not affect this uniformity



##### Division method



##### Multiplication method

$h(k) = (A \cdot k \text{ mod } 2^w)\text{rsh}(w - r)$

$A$ is an odd integer in the range $2^{w-1} < A < 2^w$, DON'T pick it too close to either edges.



##### Resolving collisions by open addressing