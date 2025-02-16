
HTTP（HyperText Transfer Protocol）方法是客户端与服务器通信时使用的操作类型，用于指示客户端希望对资源执行的操作。以下是常见的 HTTP 方法及其用途：

---

### 1. **GET**
   - **用途**：请求指定的资源。通常用于获取数据，而不对服务器状态进行修改。
   - **特点**：
     - 请求参数附加在 URL 中（例如 `/users?id=123`）。
     - 可以被缓存。
     - 不应修改服务器状态。
   - **示例**：
     ```bash
     GET /users/123 HTTP/1.1
     Host: example.com
     ```

---

### 2. **POST**
   - **用途**：向服务器提交数据，通常用于创建新资源或触发某些操作。
   - **特点**：
     - 请求参数包含在请求体中。
     - 不可缓存。
     - 可能会修改服务器状态。
   - **示例**：
     ```bash
     POST /users HTTP/1.1
     Host: example.com
     Content-Type: application/json

     {"name": "John", "email": "john@example.com"}
     ```

---

### 3. **PUT**
   - **用途**：更新指定资源的所有内容。如果资源不存在，则创建新资源。
   - **特点**：
     - 请求参数包含在请求体中。
     - 幂等性（多次执行结果相同）。
   - **示例**：
     ```bash
     PUT /users/123 HTTP/1.1
     Host: example.com
     Content-Type: application/json

     {"name": "John", "email": "john@example.com"}
     ```

---

### 4. **PATCH**
   - **用途**：部分更新指定资源的内容。
   - **特点**：
     - 请求参数包含在请求体中。
     - 非幂等性（多次执行结果可能不同）。
   - **示例**：
     ```bash
     PATCH /users/123 HTTP/1.1
     Host: example.com
     Content-Type: application/json

     {"email": "john.doe@example.com"}
     ```

---

### 5. **DELETE**
   - **用途**：删除指定的资源。
   - **特点**：
     - 幂等性（多次执行结果相同）。
   - **示例**：
     ```bash
     DELETE /users/123 HTTP/1.1
     Host: example.com
     ```

---

### 6. **HEAD**
   - **用途**：与 GET 类似，但只返回响应头，不返回响应体。通常用于检查资源是否存在或获取资源的元数据。
   - **特点**：
     - 不返回响应体。
     - 可以被缓存。
   - **示例**：
     ```bash
     HEAD /users/123 HTTP/1.1
     Host: example.com
     ```

---

### 7. **OPTIONS**
   - **用途**：获取服务器支持的 HTTP 方法或其他选项。
   - **特点**：
     - 通常用于跨域请求（CORS）预检。
   - **示例**：
     ```bash
     OPTIONS /users/123 HTTP/1.1
     Host: example.com
     ```

---

### 8. **TRACE**
   - **用途**：用于诊断，返回客户端发送的请求内容。
   - **特点**：
     - 主要用于调试。
   - **示例**：
     ```bash
     TRACE /users/123 HTTP/1.1
     Host: example.com
     ```

---

### 9. **CONNECT**
   - **用途**：用于建立与目标资源的隧道连接（通常用于 HTTPS 代理）。
   - **特点**：
     - 主要用于代理服务器。
   - **示例**：
     ```bash
     CONNECT example.com:443 HTTP/1.1
     Host: example.com
     ```

---

### 总结

| 方法      | 用途                           | 幂等性 | 缓存性 | 请求体 |
|-----------|--------------------------------|--------|--------|--------|
| GET       | 获取资源                       | 是     | 是     | 无     |
| POST      | 提交数据或创建资源             | 否     | 否     | 有     |
| PUT       | 更新或创建资源                 | 是     | 否     | 有     |
| PATCH     | 部分更新资源                   | 否     | 否     | 有     |
| DELETE    | 删除资源                       | 是     | 否     | 无     |
| HEAD      | 获取响应头                     | 是     | 是     | 无     |
| OPTIONS   | 获取服务器支持的选项           | 是     | 是     | 无     |
| TRACE     | 诊断请求                       | 是     | 是     | 无     |
| CONNECT   | 建立隧道连接（用于代理）       | 是     | 否     | 无     |

---

### 实际应用场景

1. **RESTful API 设计**：
   - `GET /users`：获取用户列表。
   - `POST /users`：创建新用户。
   - `GET /users/123`：获取 ID 为 123 的用户。
   - `PUT /users/123`：更新 ID 为 123 的用户。
   - `DELETE /users/123`：删除 ID 为 123 的用户。

2. **表单提交**：
   - 使用 `POST` 提交表单数据。

3. **文件上传**：
   - 使用 `POST` 或 `PUT` 上传文件。

4. **跨域请求**：
   - 使用 `OPTIONS` 进行预检请求。

---

-- 在windows powershell 使用scoop(网址：https://scoop.sh/)下载
   与更新migrate
   下载命令：scoop install migrate
   更新命令：scoop update migrate

-- 使用migrate更新数据库表结构（包括添加新表，添加外键）

migrate create -ext sql -dir db/migration -seq add_users

-- 升级migrate

brew upgrade golang-migrate

