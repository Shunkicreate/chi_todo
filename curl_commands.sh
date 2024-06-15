#!/bin/bash

# すべてのToDoを取得 (GET /todos)
curl -X GET http://localhost:8080/todos

# 特定のToDoを取得 (GET /todos/{id})
curl -X GET http://localhost:8080/todos/1

# 新しいToDoを作成 (POST /todos)
curl -X POST http://localhost:8080/todos -H "Content-Type: application/json" -d '{"title": "New Todo", "completed": false}'
curl -X POST http://localhost:8080/todos -H "Content-Type: application/json" -d '{"title": "Another New Todo", "completed": false}'


# 特定のToDoを更新 (PUT /todos/{id})
curl -X PUT http://localhost:8080/todos/1 -H "Content-Type: application/json" -d '{"id": 1, "title": "Updated Todo", "completed": true}'


# 特定のToDoを削除 (DELETE /todos/{id})
curl -X DELETE http://localhost:8080/todos/?id=1
