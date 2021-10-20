# 构建一个标准的go项目架构
## logic： 业务逻辑层
## internale/data： 数据层
## logic与data层通过接口进行数据传输，logic不允许直接操作db


![avatar](images/gostd.png)
![avatar](images/system.png)
