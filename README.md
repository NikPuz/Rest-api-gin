При разработке использовалось:
- Среда разработки IntelliJ IDEA
- Фреймворки/библиотеки Gin, jwt-go, go-sql-driver, bcrypt.
- Система управления базами данных MySQL
- Локальный веб-сервер Open Server Panel
- Система контроля версий GitHub
- Система контейнеризации Docker

Сервис поддерживает следующие запросы:
- GET /v1/albums/
- GET /v1/albums/page/:page
- GET /v1/albums/:id (требуется наличие JWT)
- DELETE /v1/albums/:id (требуется наличие JWT)
- POST /v1/albums (требуется наличие JWT)
- PATCH /v1/albums (требуется наличие JWT)
- POST /v1/signin
- POST /v1/login (отвечает парой токенов RefreshToken и JWToken)
- POST /v1/login/refresh (отвечает парой токенов RefreshToken и JWToken)

Столбцы базы данных:

![Столбцы](https://i.ibb.co/zGsGmGT/image.png)

Запуск:
- Установленный Docker https://www.docker.com/products/docker-desktop/
- `$ docker compose --env-file ./config/app.env up` в папке проекта
