# 🚀 Awesome Dockerized To-Do App with NGINX Proxy Manager! 🎯

Welcome to this amazing project that showcases a **powerful Go-based To-Do App** using **Docker** and **NGINX Proxy Manager** for super-simple hosting and reverse proxy management! 🌟

## ✨ Incredible Features

1. 🔥 **Go-Based To-Do App**:
   - ⚡️ Lightning-fast backend powered by Gin!
   - 🎨 Beautiful UI with static HTML, CSS, and images!

2. 🐳 **Dockerized Perfection**:
   - 🏗️ Smart multi-stage Docker build for maximum efficiency!
   - 🚀 Deploy in seconds with containerization!

3. 🛡️ **NGINX Proxy Manager Magic**:
   - 🔒 SSL and reverse proxy management made super easy!
   - 🌐 Public-facing apps with zero hassle!

## 🚀 Let's Get Started!

### 📋 Prerequisites
- 🐳 Docker
- 🔧 Docker Compose

### 🎯 Quick Start Guide

1. 📥 Clone the repository:
   ```bash
   git clone https://github.com/Mark0025/Learn-go
   cd todoapp
   ```

2. 🚀 Launch the containers:
   ```bash
   docker-compose up --build -d
   ```

3. 🔧 Set up Proxy Manager:
   - 🌐 Visit http://<server_ip>:81
   - 🔑 Login credentials:
     - Email: admin@example.com
     - Password: changeme

4. ⚙️ Configure Proxy Host:
   - Domain: todoapp.example.com
   - Forward Hostname/IP: todoapp
   - Forward Port: 8080
   - Scheme: http

5. 🎉 Access Your App:
   - 🌐 Open your browser and go to http://todoapp.example.com
   - ✨ Enjoy your awesome new To-Do App!
