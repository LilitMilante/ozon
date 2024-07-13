# OZON 

1. Создать сервис ондайн магазина.
2. Создать сервис отправки имэйлов.
3. Создать сервис оплаты (балансы, транзакции).

### Сервис онлайн магазина

- [ ] База postgreSQL pgx
- [ ] Логгер slog
- [ ] config yml
- [ ] migrator goose
- [ ] unit tests на config и validation
- [ ] integration tests на репозиторий
- [ ] gitlab config с джобами на линтер и запуск тестов
___________________
- [ ] Все входящие запросы должны логироваться и иметь ID
- [ ] Все ошибки должны логироваться и иметь тот же ID запроса
- [X] Регистрация и вход продавцов, вход черех Cookie
- [ ] Добавление карточек товаров продавца
- [ ] Получение списка товаров продавцом (его товаров)
- [ ] Всего есть 5 категорий товаров (электроника, косметика, авто, товары для дома, товары для детей)
- [ ] Редактирование товаров
- [ ] Удаление товаров (soft)
- [ ] Добавить валидацию на все запросы
- [ ] Фильтрация и сортировка списка товаров (поиск по ID или имени, по категории; сортировка по цене, категории, дате добавления в обе стороны)
- [ ] Загрузка товаров через csv file
- [ ] Регистрация требует верификации через почту
