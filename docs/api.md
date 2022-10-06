# Настройка
в configs/config.yml

port - порт сервера

dbFilePath - название базы данных 

---

# GET запросы

> Получение списка всех водителей
> > curl /get/all-drivers
>  
_Пример ответа_
```json
{
  "res": [
    {
      "Id": 0,
      "Name": "qwe",
      "SecondName": "rty",
      "Age": 12,
      "WorkExperience": 1,
      "Citizenship": "РФ",
      "Address": "Пушкина",
      "MobilePhone": "89999999999",
      "Email": "email@email.ru",
      "Car": null
    }
  ],
  "status": "ok"
}
```
---
> Получение списка всех машин
> > curl /get/all-cars

_Пример ответа_
```json
{
  "res": [
    {
      "Id": 1,
      "Num": "123321",
      "Model": "Lada granta",
      "Mileage": 123
    }
  ],
  "status": "ok"
}
```
---
> Получение машины
> 
> num - номер машины (string)
> > curl /get/car?num
>

_Пример ответа_
```json
{
  "res": [
    {
      "Id": 1,
      "Num": "123321",
      "Model": "Lada granta",
      "Mileage": 123
    }
  ],
  "status": "ok"
}
```
---
> получение водителя
> 
> name - имя водителя (string)
> 
> secondName - фамилия водителя (string)
> > curl get/driver?name&secondName
> 
_Пример ответа_
```json
{
  "res": [
    {
      "Id": 0,
      "Name": "qwe",
      "SecondName": "rty",
      "Age": 12,
      "WorkExperience": 1,
      "Citizenship": "РФ",
      "Address": "Пушкина",
      "MobilePhone": "89999999999",
      "Email": "email@email.ru",
      "Car": null
    }
  ],
  "status": "ok"
}
```
---
# POST запросы
>Добавить водителя
> 
> name - имя водителя (string)
> 
> secondName - фамилия водителя (string)
> 
> age - возраст водителя (int8)
> 
> workExperience - опыт работы (int8)
> 
> citizenship - гражданство (string)
> 
> address - адрес проживания (string)
> 
> mobilePhone - номер мобильного телефона (string)
> 
> email - почта (string)
> 
> carId - id машины водителя, если машины нет, то -1 (int64)
> >curl /add/driver?name&secondName&age&workExperience&citizenship&address&mobilePhone&email&carId
> 
_Пример ответа_

```json
{
  "status": "ok"
}
```
---
> Добавить машину
> num - номер машины (string)
> model - моедель машины (string)
> mileage - пробег (int8)
> > curl /add/car?num&model&mileage

```json
{
  "status": "ok"
}
```
---