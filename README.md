#Beellama REST API
**Beellama – это REST API проект, разработанный на языке Go для управления пользователями и  запросами в языковой модель tinyllama через фреймворк ollama.

##Запуск проекта
###1. Настройте config.yaml файл в папке app/config
###2. Поднимите докер с помощью команды:
```shell
make build
```
###3. Запустите миграцию этой командой:
```shell
make migrate-up
```

##Документация API
##Основные эндпоинты:
###1. POST /api/process
Обрабатывает запрос через tinyllama.
Пример запроса:
```json
{
  "text": "Привет, как дела?"
}
```
Пример ответа:
```json
{
  "response": "Привет! У меня всё хорошо, спасибо за вопрос!"
}
```
###2. GET /api/history
Возвращает историю обработанных запросов.
Пример ответа:
```json
{
  "history": [
    {
      "text": "Привет, как дела?",
      "response": "Привет! У меня всё хорошо, спасибо за вопрос!",
      "created_at": "2025-01-24T12:00:00Z"
    }
  ]
}
```
###3. POST /api/register
Регистрирует нового пользователя.
Пример запроса:
```json
{
  "username": "test_user",
  "password": "secure_password"
}
```
Пример ответа:
```json
{
  "message": "User registered successfully",
  "user_id": 1
}
```