-- ✅ users
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

-- ✅ projects
CREATE TABLE
    IF NOT EXISTS projects (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL UNIQUE,
        description TEXT,
        status VARCHAR(50) NOT NULL,
        owner_id INT,
        started_at TIMESTAMP DEFAULT NULL,
        finished_at TIMESTAMP DEFAULT NULL,
        deadline_at TIMESTAMP DEFAULT NULL,
        created_at TIMESTAMP DEFAULT NOW (),
        updated_at TIMESTAMP DEFAULT NOW (),
        deleted_at TIMESTAMP DEFAULT NULL,
        FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE SET NULL
    );

-- ✅ many-to-many: project_users
CREATE TABLE
    IF NOT EXISTS project_members (
        project_id INT NOT NULL,
        user_id INT NOT NULL,
        PRIMARY KEY (project_id, user_id),
        FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE,
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );


-- ✅Tasks
CREATE TABLE
    IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        description TEXT,
        priority INT,
        critical_level VARCHAR(50),
        status VARCHAR(50),
        deadline_at TIMESTAMP,
        started_at TIMESTAMP,
        completed_at TIMESTAMP,
        project_id INT NOT NULL,
        created_at TIMESTAMP DEFAULT NOW (),
        updated_at TIMESTAMP DEFAULT NOW (),
        deleted_at TIMESTAMP DEFAULT NULL,
        FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
    );

-- ✅ Task Assignees (Many-to-Many)
CREATE TABLE
    IF NOT EXISTS task_users (
        task_id INT NOT NULL,
        user_id INT NOT NULL,
        PRIMARY KEY (task_id, user_id),
        FOREIGN KEY (task_id) REFERENCES tasks (id) ON DELETE CASCADE,
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );