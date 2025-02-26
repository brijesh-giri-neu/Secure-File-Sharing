# Secure File Transfer and Sharing Platform

## Overview
This repository represents  a Secure File Transfer and Sharing Platform built using Go, PostgreSQL, Redis, and AWS S3. It provides a secure and scalable way to upload, store, and share files with role-based access control (RBAC), ensuring privacy and security. The platform enables real-time notifications, efficient file management, and end-to-end encryption for sensitive data.

## Features
- Secure File Uploads with AES-256 Encryption
- Role-Based Access Control (RBAC)
- Real-Time Notifications via WebSockets
- Efficient File Downloads with Pre-Signed URLs
- Redis Queue for Background Processing
- JWT-Based Authentication

## Tech Stack
**Backend**: Golang (`net/http`)
**Database**: PostgreSQL
**Cache & Queue**: Redis
**Storage**: AWS S3 (or local storage for dev mode)
**Security**: AES-256 encryption, JWT authentication
**Deployment**: Docker, Kubernetes

## Setup Instructions
1. Clone this repository:
   ```sh
   git clone https://github.com/yourusername/secure-file-transfer.git
   cd secure-file-transfer
   ```

2. Set up environment variables:
   ```sh
   cp .env.example .env
   ```

3. Run the backend server:
   ```sh
   cd backend
   go run main.go
   ```

4. Start the frontend:
   ```sh
   cd frontend
   npm install
   npm start
   ```

## Deployment
- Docker & Kubernetes configurations included.
- CI/CD with GitHub Actions.
