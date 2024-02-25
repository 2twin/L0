
## Начальная страница
```
http://localhost:8080/order
```

![order](https://imgur.com/PgmAWuD.png)

## Отображение данных заказа по order_uid
```
http://localhost:8080/order?order_uid=<order_uid>
```

![order_uid](https://imgur.com/0JA4JAq.png)

## Результаты теста с помощью go-wrk

![go-wrk](https://imgur.com/3IA9DCJ.png)

---

## Развертывание приложения

```shell
docker-compose build
```

```shell
docker-compose up
```

---
