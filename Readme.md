# 🧩 Task Management Project

A full-stack task management system built with:

- 🦀 **Rust** – API Gateway
- ⚙️ **Go** – Microservices - Backend
- 🖥️ **Vue 3 + TypeScript** – Frontend

## 🚀 How to Run This Project

### 1️⃣ Clone the Repository

```bash
git clone <your-repo-url>
cd <your-repo-folder>
```

### 2️⃣ Set Execution Permissions for Scripts

```bash
chmod +x setup.sh
chmod +x shutdown.sh
```

### 3️⃣ Install Frontend Dependencies

```bash
cd frontend
npm install
cd ..
```

### 4️⃣ Run the Project with Docker

```bash
./setup.sh
```

### 5️⃣ Stop the Project

```bash
./shutdown.sh
```

## 💡 Notes
- Ensure Docker and Docker Compose are installed before running.
- Database is automatically seeded using `db/init.sql`.