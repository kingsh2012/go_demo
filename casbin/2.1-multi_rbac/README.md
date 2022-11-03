# 多个RBAC
* casbin支持同时存在多个RBAC系统，即用户和资源都有角色：
```azure
[role_definition]
g = _,_
g2 = _,_

[matchers]
m = g(r.sub, p.sub) && g2(r.obj, p.obj) && r.act == p.act
```
上面的模型文件定义了两个RBAC系统g和g2，我们在匹配器中使用g(r.sub, p.sub)判断请求主体属于特定组，g2(r.obj, p.obj)判断请求资源属于特定组，且操作一致即可放行。

