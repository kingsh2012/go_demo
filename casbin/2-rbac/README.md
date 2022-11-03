# RBAC模型
* ACL模型在用户和资源都比较少的情况下没什么问题，但是用户和资源量一大，ACL就会变得异常繁琐。想象一下，每次新增一个用户，都要把他需要的权限重新设置一遍是多么地痛苦。
* RBAC（role-based-access-control）模型通过引入角色（role）这个中间层来解决这个问题。每个用户都属于一个角色，例如开发者、管理员、运维等，每个角色都有其特定的权限，权限的增加和删除都通过角色来进行。
* 这样新增一个用户时，我们只需要给他指派一个角色，他就能拥有该角色的所有权限。修改角色的权限时，属于这个角色的用户权限就会相应的修改。 

## 在casbin中使用RBAC模型需要在模型文件中添加role_definition模块：
```azure
[role_definition]
g = _, _

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
```
g = _,_定义了用户——角色，角色——角色的映射关系，前者是后者的成员，拥有后者的权限。然后在匹配器中，我们不需要判断r.sub与p.sub完全相等，只需要使用g(r.sub, p.sub)来判断请求主体r.sub是否属于p.sub这个角色即可。
## 最后我们修改策略文件添加用户——角色定义：
```azure
p, admin, data, read
p, admin, data, write
p, developer, data, read
g, dajun, admin
g, lizi, developer
```
上面的policy.csv文件规定了，dajun属于admin管理员，lizi属于developer开发者，使用g来定义这层关系。另外admin对数据data用read和write权限，而developer对数据data只有read权限。
很显然lizi所属角色没有write权限：
```azure
sub:dajun,obj:data,act:read 通过验证
sub:dajun,obj:data,act:write 通过验证 
                                      
sub:lizi,obj:data,act:read 通过验证   
sub:lizi,obj:data,act:write 未通过验证
```