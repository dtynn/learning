## Tower of Hanoi

original recursive algorithm:

```
Hanoi(n, peg_src, peg_dst, peg_tmp):
	if n > 0:
		Hanoi(n-1, peg_src, peg_tmp, peg_dst);
		move n to peg_dst;
		Hanoi(n-1, peg_tmp, peg_dst, peg_src);
```





