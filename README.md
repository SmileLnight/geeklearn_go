# geeklearn_go
## go异常处理作业
参考了go1.13的新特性，利用**errors.Is**进行判断，将dao层产生的错误包装并往上抛，当遇到一个 **sql.ErrNoRows** 的时候应返回nil而不是空数组
