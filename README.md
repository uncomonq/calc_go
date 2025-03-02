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

### Запуск 

## 1. Клонируйте репозиторий:

   ```sh
   git clone https://github.com/uncomonq/calc_go.git
   cd calc_go
   ```

## 2. Запуск орекстратора:
### Linux / macOS (bash)
```sh 
# Установка времени операций (в миллисекундах)
export TIME_ADDITION_MS=200
export TIME_SUBTRACTION_MS=200
export TIME_MULTIPLICATIONS_MS=300
export TIME_DIVISIONS_MS=400

# Запуск оркестратора
go run ./cmd/orchestrator/main.go
```

### Windows (PowerShell)
```sh
# Установка времени операций (в миллисекундах)
$env:TIME_ADDITION_MS = "200"
$env:TIME_SUBTRACTION_MS = "200"
$env:TIME_MULTIPLICATIONS_MS = "300"
$env:TIME_DIVISIONS_MS = "400"

# Запуск оркестратора
go run .\cmd\orchestrator\main.go
```

## 3. Запуск агента:
### Linux / macOS (bash)
```sh
# Указание вычислительной мощности (количество горутин) и URL оркестратора
export COMPUTING_POWER=4
export ORCHESTRATOR_URL=http://localhost:8080

# Запуск агента
go run ./cmd/agent/main.go
```
### Windows (PowerShell)
```sh
# Указание вычислительной мощности (количество горутин) и URL оркестратора
$env:COMPUTING_POWER = "4"
$env:ORCHESTRATOR_URL = "http://localhost:8080"

# Запуск агента
go run .\cmd\agent\main.go
``` 
## Использование

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

### Оркестратор

- `PORT` - порт сервера (по умолчанию 8080)
- `TIME_ADDITION_MS` - время сложения (мс)
- `TIME_SUBTRACTION_MS` - время вычитания (мс)
- `TIME_MULTIPLICATIONS_MS` - время умножения (мс)
- `TIME_DIVISIONS_MS` - время деления (мс)

### Агент

- `ORCHESTRATOR_URL` - URL оркестратора
- `COMPUTING_POWER` - количество параллельных задач


## Автор

Разработано [uncomonq](https://github.com/uncomonq).
