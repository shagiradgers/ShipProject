***
Таблица cars
```sqlite
create table if not exists cars(
    id      int     primary key,
    num     text    not null,
    model   text    not null,
    mileage int     not null
)

```
***
Таблица drivers
```sqlite
create table if not exists drivers(
    id              int                 primary key,
    name            text    not null,
    second_name     text    not null,
    age             int     not null,
    work_experience int     not null,
    citizenship     text    not null,
    address         text    not null,
    mobile_phone    text    not null,
    email           text    not null,
    carId           int,
    CONSTRAINT fk_car FOREIGN KEY (carId)
        REFERENCES cars(id)
)
```
***