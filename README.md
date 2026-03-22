##go web application

Простое веб-приложение на go которое выводит локальные переменные и переменные окружения, упакованное в Docker контейнер.
# Возможности
- Вывод локальной переменной
- Вывод переменной окружения `MY_VAR`
- Настройка порта через переменную окружения `PORT`
- Docker контейнер с минимальным образом

Для запуска приложения вам понадобится:
- Go 1.21 или выше
- Git
- Docker 20.10 или выше
# Клонирование репозитория
-git clone https://github.com/alextrojnovski/docker_go
-cd docker_go
# Запуск в докер
-docker build -t go-web-app .
-docker run -d -p 8080:8080 -e MY_VAR="Hello from Docker!" --name myapp go-web-app
# Проверка работы
-curl http://localhost:8080
