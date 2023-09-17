# restfull_go

Цель проекта:

1. Создать сервис, кторый будет записывать, изменять, удалять и получать данные из БД
2. Запросом GET будем получать данные из таблицы
3. Запросом POST будем добавлять данные в таблицу
4. Запросом PUT будем редактировать данные в таблице
5. Запросом DELETE будем удалять данные из таблицы

Если не будет заводиться, то надо поставить пакет командой:
go get github.com/iamacarpet/go-sqlite3-dynamic


TODO:
1. Реализовать авторизацию            - готово
2. Проверить ошибки                   - готово
3. Сделать разные коды ответа         - готово
4. Убирать ; из входящих запросов     - готово
5. Добавить шифрование                - готово
6. Сделать записи уникальными         - по уму это надо делать в SQL через constraint, поэтому я решил от этого отказаться
