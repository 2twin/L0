
## Начальная страница
```http request
http://localhost:8080/order
```

![order](https://imgur.com/PgmAWuD)

## Отображение данных заказа по order_uid
```http request
http://localhost:8080/order?order_uid=<order_uid>
```

![order_uid](https://imgur.com/0JA4JAq)

## Результаты тесты с помощью go-wrk

![go-wrk](https://imgur.com/3IA9DCJ)

---

## Развертывание приложения

```shell
docker-compose build
```

```shell
docker-compose up
```

---