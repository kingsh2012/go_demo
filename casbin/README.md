# RBAC
* Role-Based Access Control(下简称rbac)以角色为基础的访问控制模型，通过角色连接用户和权限，达到两者解耦的目的
## rbac中有三种类型：用户、角色、权限，分别解释一下各自的作用
* 用户：使用系统的个体，一般是自然人，通过操作系统达到一定目的
* 权限：系统中每一个模块的操作，curd，的使用许可证明
* 角色：用户和权限之间的桥梁，是一个抽象的概念，与用户或权限可以是一对多，一对一，多对一，多对多的关系，当用户操作某一个模块时，会先将用户名下的角色取出来和模块权限进行对比，如果满足条件即可操作，反之则不行

# Casbin官方介绍
* https://casbin.org/docs/zh-CN/tutorials
## Casbin是什么？
### Casbin 可以：
* 1.支持自定义请求的格式，默认的请求格式为{subject, object, action}。
* 2.具有访问控制模型model和策略policy两个核心概念。
* 3.支持RBAC中的多层角色继承，不止主体可以有角色，资源也可以具有角色。
* 4.支持内置的超级用户 例如：root 或 administrator。超级用户可以执行任何操作而无需显式的权限声明。
* 5.支持多种内置的操作符，如 keyMatch，方便对路径式的资源进行管理，如 /foo/bar 可以映射到 /foo*
### Casbin 不能：
* 1.身份认证 authentication（即验证用户的用户名和密码），Casbin 只负责访问控制。应该有其他专门的组件负责身份认证，然后由 Casbin 进行访问控制，二者是相互配合的关系。
* 管理用户列表或角色列表。 Casbin 认为由项目自身来管理用户、角色列表更为合适， 用户通常有他们的密码，但是 Casbin 的设计思想并不是把它作为一个存储密码的容器。 而是存储RBAC方案中用户和角色之间的映射关系。
## 工作原理
* 在 Casbin 中, 访问控制模型被抽象为基于 PERM (Policy, Effect, Request, Matcher) 的一个文件。 因此，切换或升级项目的授权机制与修改配置一样简单。 您可以通过组合可用的模型来定制您自己的访问控制模型。 例如，您可以在一个model中结合RBAC角色和ABAC属性，并共享一组policy规则。
* PERM模式由四个基础（政策、效果、请求、匹配）组成，描述了资源与用户之间的关系。
### 请求
* 定义请求参数。基本请求是一个元组对象，至少需要主题(访问实体)、对象(访问资源) 和动作(访问方式)
* 例如，一个请求可能长这样： r={sub,obj,act}
* 它实际上定义了我们应该提供访问控制匹配功能的参数名称和顺序。
### 策略
* 定义访问策略模式。事实上，它是在政策规则文件中定义字段的名称和顺序。
* 例如： p={sub, obj, act} 或 p={sub, obj, act, eft}
* 注：如果未定义eft (policy result)，则策略文件中的结果字段将不会被读取， 和匹配的策略结果将默认被允许。
### 匹配器
* 匹配请求和政策的规则。
* 例如： m = r.sub == p.sub && r.act == p.act && r.obj == p.obj 这个简单和常见的匹配规则意味着如果请求的参数(访问实体，访问资源和访问方式)匹配， 如果可以在策略中找到资源和方法，那么策略结果（p.eft）便会返回。 策略的结果将保存在 p.eft 中。
### 效果
* 它可以被理解为一种模型，在这种模型中，对匹配结果再次作出逻辑组合判断。
* 例如： e = some (where (p.eft == allow))
* 这句话意味着，如果匹配的策略结果有一些是允许的，那么最终结果为真。
* 让我们看看另一个示例： e = some (where (p.eft == allow)) && !some(where (p.eft == deny) 此示例组合的逻辑含义是：如果有符合允许结果的策略且没有符合拒绝结果的策略， 结果是为真。 换言之，当匹配策略均为允许（没有任何否认）是为真（更简单的是，既允许又同时否认，拒绝就具有优先地位)。
* Casbin中最基本、最简单的model是ACL。ACL中的model 配置为:
```azure
# Request definition
[request_definition]
r = sub, obj, act

# Policy definition
[policy_definition]
p = sub, obj, act

# Policy effect
[policy_effect]
e = some(where (p.eft == allow))

# Matchers
[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
```
* ACL模型的一个示例策略类似：
```azure
p, alice, data1, read
p, bob, data2, write
```
* 它意味着：
  * alice可以读取data1
  * bob可以编写data2
* 我们还支持多行模式，通过在结尾处追加“\”：
```azure
#  匹配器
[matchers]
m = r.sub == p.sub && r.obj == p.obj \ 
  && r.act == p.act
```
* 此外，对于 ABAC，您在可以在 Casbin golang 版本中尝试下面的in (jCasbin 和 Node-Casbin 尚不支持) 操作：
```azure
# Matchers
[matchers]
m = r.obj == p.obj && r.act == p.act || r.obj in ('data2', 'data3')
```
但是你应确保数组的长度大于 1，否则的话将会导致 panic 。

# Casbin使用
