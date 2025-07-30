-- ✅ สร้างตาราง users (ใน task_management)
CREATE TABLE
    IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        password TEXT NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        first_name VARCHAR(100) NOT NULL,
        last_name VARCHAR(100) NOT NULL,
        created_at TIMESTAMP DEFAULT NOW (),
        updated_at TIMESTAMP DEFAULT NOW (),
        deleted_at TIMESTAMP DEFAULT NULL
    );

-- ✅ สร้างตาราง tasks
CREATE TABLE
    IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255),
        description TEXT,
        priority INT,
        critical_level VARCHAR(50),
        status VARCHAR(50),
        deadline_at TIMESTAMP,
        started_at TIMESTAMP,
        completed_at TIMESTAMP,
        user_id INT,
        created_at TIMESTAMP DEFAULT NOW (),
        updated_at TIMESTAMP DEFAULT NOW (),
        deleted_at TIMESTAMP DEFAULT NULL
    );