# 构建一个标准的go项目架构
## logic： 业务逻辑层
## internale/data： 数据层
## logic与data层通过接口进行数据传输，logic不允许直接操作db

![avatar](images/gostd.png)

./gostd/
├── **cmd**         
│   └── server
│       ├── app.yaml
│       └── main.go
├── **internal**
│   └── **data**
│       ├── api
│       │   ├── student
│       │   │   ├── grades
│       │   │   │   ├── grades.go
│       │   │   │   └── interface.go
│       │   │   ├── student_api.go
│       │   │   └── user-info
│       │   │       ├── interface.go
│       │   │       └── user_info.go
│       │   └── teacher
│       │       └── teacher_api.go
│       ├── database
│       │   └── database.go
│       ├── data.go
│       └── readme.md
├── **logic**
│   ├── api
│   │   ├── api.go
│   │   └── handler
│   │       ├── file
│   │       │   └── file.go
│   │       ├── handler.go
│   │       └── index.go
│   ├── conf
│   │   └── conf.go
│   ├── logic.go
│   └── readme.md
└── readme.md
