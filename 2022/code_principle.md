##  日常开发原则

#### 软件设计的原则

- Don’t Repeat Yourself (DRY)
- Keep It Simple, Stupid (KISS 所见即所得, 保持简约)
- Program to an interface, not an implementation （注重接口，而不是实现，依赖接口，而不是实现）
- Command-Query Separation (CQS)  – 命令-查询分离原则
- You Ain’t Gonna Need It (YAGNI 如无必要，勿增复杂性 )
- 面向对象的S.O.L.I.D 原则
  - Single Responsibility Principle (SRP) – 职责单一原则
  - Open/Closed Principle (OCP) – 开闭原则
  - Liskov substitution principle (LSP) – 里氏代换原则
  - Interface Segregation Principle (ISP) – 接口隔离原则
- Dependency Inversion Principle (DIP) – 依赖倒置原则
- Common Closure Principle（CCP）– 共同封闭原则
- High Cohesion & Low/Loose coupling  – 高内聚， 低耦合
- Convention over Configuration（CoC）– 惯例优于配置原则
- Separation of Concerns (SoC) – 关注点分离
- Design by Contract (DbC) – 契约式设计
- Acyclic Dependencies Principle (ADP) – 无环依赖原则



#### 规范

- **服务间调用的协议标准和规范**。这其中包括 Restful API路径, HTTP 方法、状态码、标准头、自定义头等，返回数据 JSon Scheme……等。
- **一些命名的标准和规范**。这其中包括如：用户 ID，服务名、标签名、状态名、错误码、消息、数据库……等等
- **日志和监控的规范**。这其中包括：日志格式，监控数据，采样要求，报警……等等
- **配置上的规范**。这其中包括：操作系统配置、中间件配置，软件包……等等
- **中间件使用的规范**。数据库，缓存、消息队列……等等
- **软件和开发库版本统一**。整个组织架构内，软件或开发库的版本最好每年都升一次级，然后在各团队内统一。



- 尽量使用REST API开定义接口

| 方法    | 描述                                                    | 幂等 |
| ------- | ------------------------------------------------------- | ---- |
| GET     | 用于查询操作，对应于数据库的 `select` 操作              | ✔︎    |
| PUT     | 用于所有的信息更新，对应于数据库的 `update `操作        | ✔︎︎    |
| DELETE  | 用于更新操作，对应于数据库的 `delete` 操作              | ✔︎︎    |
| POST    | 用于新增操作，对应于数据库的 `insert` 操作              | ✘    |
| HEAD    | 用于返回一个资源对象的“元数据”，或是用于探测API是否健康 | ✔︎    |
| PATCH   | 用于局部信息的更新，对应于数据库的 `update` 操作        | ✘    |
| OPTIONS | 获取API的相关的信息。                                   | ✔︎    |

- 使用标准的http状态码返回

  



#### 参考
[左耳朵耗子：“一把梭：REST API 全用 POST”](https://coolshell.cn/articles/22173.html)

[左耳朵耗子：我做系统架构的一些原则](https://coolshell.cn/articles/21672.html)

[左耳朵耗子：一些软件设计的原则](https://coolshell.cn/articles/4535.html)

[左耳朵耗子：优质代码的十诫](https://coolshell.cn/articles/1007.html)

[左耳朵耗子：编程中的命名设计那点事](https://coolshell.cn/articles/990.html)