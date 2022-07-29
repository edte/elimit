令牌桶：自适应限流，请求多时限流多，请求少时限流少，也没有毛刺现象

溜桶：速率永远相同，不符合实际请求速率不均衡的情况

计数器：毛刺想象，二倍上限现象，实现有同步异步之分（睡眠还是等待？还是返回 bool），同时计算器的清空也有同步异步之分，究竟是每次请求取时间比较，还是单独开携程更新（这里还涉及到同一个函数部分代码并发调用只调用一次的情况，是另外抽出来开携程？还是 once.do 函数？同事并发更改计数器也是个问题），这里和一般的缓存更新设计方法差不多。

滑动窗口，基本上去掉计数器那种毛刺现象，把请求均分到周期的不同部分，窗口越多越均衡，不过同样不算绝对均衡，也存在速率请求分布不合理的问题

而综上几个，显然令牌桶是包含了所有优点，而又克服了其它算法缺点的存在，所以令牌桶的使用比较高频。

但是同时，上面这些算法都有两个无法解决的问题：
1. 速率限制无法更改，像多级速率的场景无法实现，比如1分钟5，1小时20这样递增，可以使用日志滑动窗口算法实现，是滑动窗口算法的改进，增加了不同的策略，在不同窗口选择不同限制

2. 上面几个都是单机的算法，不适用于分布式限流的情况，而分布式限流则可以考虑使用 redis lua，以及 kafka 解藕的方式

