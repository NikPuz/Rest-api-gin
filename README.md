При разработке использовалось:
- Среда разработки IntelliJ IDEA
- Фреймворки/библиотеки Gin, jwt-go, go-sql-driver, bcrypt.
- Система управления базами данных MySQL
- Локальный веб-сервер Open Server Panel
- Система контроля версий GitHub

Сервис поддерживает следующие запросы:

- POST /album/:id (требуется наличие JWT)
- POST /album/delete/:id (требуется наличие JWT)
- POST /album/add (требуется наличие JWT)
- POST /album/update (требуется наличие JWT)
- GET /albums/
- GET /albums/:page (пагинация)
- POST /login (отвечает парой токенов RefreshToken и JWToken)

Столбцы базы данных:

![Столбцы](https://i.ibb.co/zGsGmGT/image.png)
