# Проект "llds"

Этот проект представляет собой простой API на языке программирования Go, предоставляющий возможности регистрации, входа и выхода из системы. Проект также взаимодействует с базой данных PostgreSQL для хранения пользовательских данных.

## Начало работы

1. **Склонируйте репозиторий:**
    ```bash
    git clone https://github.com/logikablat/llds.git
    cd llds
    ```

2. **Инициализация модуля Go:**
    ```bash
    go mod init llds
    ```

3. **Установка зависимостей:**
    ```bash
    go get -u github.com/gorilla/mux
    go get -u github.com/sirupsen/logrus
    go get -u github.com/lib/pq
    ```

4. **Запуск проекта:**
    ```bash
    go run main.go
    ```

5. **Использование API:**
    - Регистрация нового пользователя:
      ```http
      POST http://localhost:8080/register
      {"username": "your_username", "password": "your_password"}
      ```
    - Вход в систему:
      ```http
      POST http://localhost:8080/login
      {"username": "your_username", "password": "your_password"}
      ```
    - Выход из системы:
      ```http
      POST http://localhost:8080/logout
      ```

## Зависимости

- [gorilla/mux](https://github.com/gorilla/mux) - Мощный маршрутизатор HTTP для Go.
- [sirupsen/logrus](https://github.com/sirupsen/logrus) - Библиотека логирования для Go.
- [lib/pq](https://github.com/lib/pq) - PostgreSQL драйвер для Go.
