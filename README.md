# Тестовое задание FRESHQIWI 2023

## Задача

Реализовать консольную утилиту, которая выводит курсы валют ЦБ РФ за определенную дату. Для получения курсов необходимо использовать официальный API ЦБ РФ https://www.cbr.ru/development/sxml/.

## Интерфейс взаимодействия

Linux/Mac
```bash
./currency_rates --code=USD --date=2022-10-08
```

## Описание параметров

- code - код валюты в формате ISO 4217

- date - дата в формате YYYY-MM-DD

## Используемые технологии

- [`Go 1.20`](https://go.dev/)
- XML парсер [`encoding/xml`](https://pkg.go.dev/encoding/xmlr)
- Библиотека для работы со слайсами [`slices`](https://pkg.go.dev/golang.org/x/exp/slices)
- Библиотека для работы с http [`net/http`](https://pkg.go.dev/net/http)
