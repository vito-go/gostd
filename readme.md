# 构建一个标准的go项目架构
## logic： 业务逻辑层
## internale/data： 数据层
## logic与data层通过接口进行数据传输，logic不允许直接操作db
- 做微服务架构时候，data层可以单独抽出来独立的服务


![avatar](images/gostd.png)
![avatar](images/system.png)
