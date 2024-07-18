# Moneyger-Backend

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/zVitorSantos/Moneyger-Backend/blob/main/LICENSE)
[![Build and Test](https://github.com/zVitorSantos/Moneyger-Backend/actions/workflows/ci.yml/badge.svg)](https://github.com/zVitorSantos/Moneyger-Backend/actions)
[![codecov](https://codecov.io/gh/zVitorSantos/Moneyger-Backend/branch/main/graph/badge.svg?token=220E44857K)](https://codecov.io/gh/zVitorSantos/Moneyger-Backend)
[![Go Report Card](https://goreportcard.com/badge/github.com/zVitorSantos/Moneyger-Backend)](https://goreportcard.com/report/github.com/zVitorSantos/Moneyger-Backend)
[![Swagger Validator](https://img.shields.io/swagger/valid/3.0?specUrl=https://raw.githubusercontent.com/zVitorSantos/Moneyger-Backend/main/docs/swagger.yaml)](https://validator.swagger.io/validator/debug?url=https://raw.githubusercontent.com/zVitorSantos/Moneyger-Backend/main/docs/swagger.yaml)


# Backend for a minimalist Money Management API designed to support a Vue interface.

## Table of Contents

1. [Initial Setup](#initial-setup)
2. [Database](#database)
3. [Project Structure](#project-structure)
4. [API Documentation](#api-documentation)
5. [Testing and CI/CD](#testing-and-cicd)
6. [Containerization](#containerization)
7. [Continuous Development](#continuous-development)
8. [Technologies Used](#technologies-used)
9. [Future Steps](#future-steps)

## Checklist

### Implementation Stages

1. ## Initial Setup
   - [x] Set up development environment
   - [x] Initialize Git repository and configure GitHub
   - [x] Configure Go Modules

2. ## Database
   - [x] Choose and configure database (PostgreSQL)
   - [x] Define data model
   - [x] Create migrations for initial tables

3. ## Project Structure
   - [x] Create basic project structure
   - [x] Implement basic CRUD for main entities
   - [ ] Implement authentication and authorization

4. ## API Documentation
   - [x] Set up Swagger for API documentation
   - [ ] Validate Swagger documentation
   - [ ] Add missing documentation:
     - [ ] Endpoints for user management
     - [ ] Endpoints for transaction management
     - [ ] Authentication flow

5. ## Testing and CI/CD
   - [x] Set up unit tests
   - [ ] Implement integration tests
   - [x] Configure GitHub Actions for CI/CD
   - [x] Integrate Codecov for test coverage
   - [ ] Set up staging environment
   - [ ] Implement automated deployment to staging
   - [ ] Implement automated deployment to production

6. ## Containerization
   - [x] Create Dockerfile for the application
   - [x] Set up Docker Compose for development environment
   - [ ] Configure Docker for production environment
   - [ ] Set up Kubernetes for container orchestration
   - [ ] Create Helm charts for Kubernetes deployments

7. ## Continuous Development
   - [ ] Configure complete CI/CD pipeline
   - [ ] Set up automated tests in CI/CD pipeline
   - [ ] Set up monitoring and alerting with Prometheus & Grafana
   - [ ] Configure centralized logging with ELK Stack
   - [ ] Set up Redis for caching
   - [ ] Implement additional backend functionalities
   - [ ] Develop frontend interface in Vue.js
   - [ ] Test and validate full integration between backend and frontend

### Technologies Used

- [x] Go
- [x] PostgreSQL
- [x] Docker
- [x] Docker Compose
- [ ] JWT for authentication
- [ ] Redis for caching
- [ ] Prometheus & Grafana for monitoring
- [ ] ELK Stack for logging
- [x] Swagger for documentation
- [x] GitHub Actions for CI/CD
- [ ] Codecov for test coverage
- [ ] Kubernetes for container orchestration

### Future Steps

- [ ] Implement more functionalities in the backend
- [ ] Develop frontend interface in Vue.js
- [ ] Test and validate complete integration between backend and frontend
