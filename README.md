# CLI-менеджер задач на Go
Небольшая CLI-утилита для удобного оформления задач
## Установка:
```bash
go install github.com/wasurene/gotodo@latest
```
Можно еще собрать локально:
```bash
git clone https://github.com/wasurene/gotodo.git
cd gotodo
go build
```

## Примеры использования
```bash
gotodo add "Купить хлеб"                    # добавить задачу
gotodo add "Сдать курсовую" --priority high --due 2026-12-25 # добавить задачу с приоритетом и дедлайном
gotodo list                                  # показать все задачи
gotodo list --filter done                    # только выполненные
gotodo list --filter active                  # только активные
gotodo done 1                                # отметить задачу 1 выполненной
gotodo rm 1                                  # удалить задачу 1
gotodo clear                                 # удалить все задачи
```

## Флаги
```bash
--priority - приоритет задачи (по умолчанию низкий)
--due - указать дедлайн
--filter - показать выполненные(done)/активные(active)
```

## Где хранятся данные

Данные хранятся в домашней папке: `~/.gotodo/tasks.json`

## Технологии
- Go 1.27+
- `github.com/fatih/color`

## Лицензия
MIT