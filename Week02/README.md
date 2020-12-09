学习笔记
## 作业
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答：在底层的基础架构我们应该尽可能的少打些日志，也就是少些wrap，
但是在dao层我们的业务更多的是要处理业务逻辑，需要很好的定位问题，所以当遇到一个sql.ErrNoRows错误时候我们应该尽可能的wrap，能更好的定位到问题。

```
dao: 
 return errors.Wrapf(code.NotFound, fmt.Sprintf("sql: %s error: %v", sql, err))
biz:
if errors.Is(err, code.NotFound} {
}
```
##笔记📒
1、panic 要慎用！！！！
2、