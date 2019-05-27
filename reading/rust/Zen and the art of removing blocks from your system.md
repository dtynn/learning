原文 [Programming Servo: Zen and the art of removing blocks from your system](https://medium.com/@polyglot_factotum/programming-servo-zen-and-the-art-of-removing-blocks-from-your-system-51c1b7d404e3)



- Servo 是干嘛的以及为什么要参与 Servo

- 群体智慧的优劣

- 正题来了, 如何从并行系统中移除阻塞逻辑

  - 阻塞逻辑的来源
  - 问题
  - 解决方案
    - 把需要等待的响应变成 event-loop 中的一份子.

- 总结

  - 阻塞逻辑会隐式地影响到整个并行系统

  - 由于惯性, 这段阻塞逻辑会变成其他逻辑"依赖"的"功能"

    