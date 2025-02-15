CREATE TABLE IF NOT EXISTS students (
    id INT PRIMARY KEY COMMENT '学号',
                                        name VARCHAR(50) COMMENT '姓名',
    gender VARCHAR(10) COMMENT '性别',
    major VARCHAR(50) COMMENT '专业'
    )CHARSET=utf8mb4;
