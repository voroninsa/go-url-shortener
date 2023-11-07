# go-url-shortener
## Тестовое задание для OZON

Запуск: 
```
docker-compose up --build
```

Сервис бежит на tcp порту 8080, работает по HTTP

Эндпоинты:
- `/shorten-url` - поддерживает _POST_, _Content-Type: application/json_.
  _body: {«url»: string}_.
  Возвращает сокращенный URL
- `/{short-url}` - поддерживает _GET_, где _short-url_ - сокращенный URL,
  полученный от сервиса после _POST_ запроса. Возвращает полный URL

Для смены хранилища заменить в Dockerfile 
в последней строке `"SQL"` на `"InMemory"`

Для запуска нужного теста перейдите в 
директорию с ним и выполните 
```
go test -v
```
