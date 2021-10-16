# 构建一个标准的go项目架构


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
