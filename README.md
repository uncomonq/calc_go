# CalcGo - Распределенная система вычислений арифметических выражений

CalcGo - это система распределенных вычислений для арифметических выражений, разработанная на языке Golang. Проект включает оркестратор (сервер) и вычислительные агенты, взаимодействующие для эффективного выполнения вычислений.

## Функциональность

- **Оркестратор**

  - Принимает арифметические выражения через API.
  - Декомпозирует выражение на отдельные вычислительные задачи.
  - Распределяет задачи между вычислительными агентами.
  - Собирает и возвращает результаты вычислений.

- **Агенты**

  - Запрашивают задачи у оркестратора.
  - Выполняют вычисления с заданными задержками.
  - Отправляют результаты обратно оркестратору.

## Установка и запуск

### Требования

- Golang >= 1.20
- Docker (опционально, для контейнеризированного развертывания)

### Запуск вручную

1. Клонируйте репозиторий:

   ```sh
   git clone https://github.com/uncomonq/calc_go.git
   cd calc_go
   ```

2. Запустите оркестратор:

   ```sh
   go run cmd/orchestrator/main.go
   ```

3. Запустите вычислительный агент:

   ```sh
   go run cmd/agent/main.go
   ```

### Запуск с Docker

1. Соберите образы:

   ```sh
   docker-compose build
   ```

2. Запустите контейнеры:

   ```sh
   docker-compose up
   ```

## Использование API

Оркестратор предоставляет REST API для управления вычислениями.

### Добавление выражения

```sh
curl --location 'http://localhost:8080/api/v1/calculate' \
--data '{"expression": "3 + 5 * (2 - 8)"}'
```

### Запрос статуса выражения

```sh
curl -X GET "http://localhost:8080/expressions/{id}" -H "Content-Type: application/json"
```

### Получение задач агентами

```sh
curl -X GET "http://localhost:8080/tasks"
```

## Примеры использования

### Пример 1: Простое выражение

Отправляем запрос на вычисление:

```sh
curl --location 'http://localhost:8080/api/v1/calculate' \
--data '{"expression": "10 + 20"}'
```

Ответ:

```json
{
  "id": "1234",
  "status": "processing"
}
```

Через некоторое время запрашиваем результат:

```sh
curl -X GET "http://localhost:8080/expressions/1234"
```

Ответ:

```json
{
  "id": "1234",
  "status": "completed",
  "result": 30
}
```

### Пример 2: Выражение со скобками и операциями

Запрос:

```sh
curl --location 'http://localhost:8080/api/v1/calculate' \
--data '{"expression": "(5 + 3) * 2"}'
```

Ответ после обработки:

```json
{
  "id": "5678",
  "status": "completed",
  "result": 16
}
```

### Пример 3: Ошибочное выражение

Отправляем запрос с некорректным выражением:

```sh
curl --location 'http://localhost:8080/api/v1/calculate' \
--data '{"expression": "3 + * 5"}'
```

Ответ:

```json
{
  {"error":"expected number at position 2"}
}
```

## Переменные среды

- `COMPUTING_POWER` - Количество параллельных вычислений, выполняемых агентом.
- `TASK_DELAY` - Задержка выполнения операций (имитация нагрузки).

## Лицензия

Этот проект распространяется под лицензией MIT. Подробнее см. в файле [LICENSE](LICENSE).

## Автор

Разработано [uncomonq](https://github.com/uncomonq).

