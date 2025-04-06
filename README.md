# 📝 REST API на Go + Fiber

## 📌 Описание проекта

Необходимо разработать REST API для управления задачами (TODO-лист).  
API должно позволять:
- Создавать задачу.
- Читать список задач.
- Обновлять задачу.
- Удалять задачу.

---

## ⚙️ Стек технологий

- Язык программирования: **Go**
- Веб-фреймворк: **Fiber**
- База данных: **PostgreSQL** (через `pgx`)
- Среда выполнения: **локальная** (использование Docker не обязательно, но приветствуется)

---

## 📄 Запуск проекта

- Скопировать `.env.example` в файл `.env`
- В корневой папке проекта прописать команду:

```bash
    make start
```

## 🗃️ Структура базы данных

Создаётся таблица `tasks` со следующими полями:

```sql
CREATE TABLE tasks (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  description TEXT,
  status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new',
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);
```

---

## 🚀 API Эндпоинты

- `POST /tasks` — создание новой задачи
- `GET /tasks` — получение списка всех задач
- `PUT /tasks/:id` — обновление задачи по ID
- `DELETE /tasks/:id` — удаление задачи по ID

---

## 📌 Требования к реализации

- CRUD-операции должны быть реализованы полностью.
- Ошибки обрабатываются корректно.
- Код должен быть читаемым и логически структурированным.
- Работа с базой данных через `pgx` должна быть надёжной и устойчивой.

---

## 📄 Пример запроса

```bash
curl -X POST http://localhost:3000/tasks \
  -H "Content-Type: application/json" \
  -d '{
        "title": "Протестировать API",
        "description": "Проверить создание задачи",
        "status": "new"
      }'
```



