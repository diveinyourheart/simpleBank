-----------------------------------------------------------------------------------------------------------------------

从当前分支切换到另一本地分支
git checkout master
git switch master

删除本地分支
git branch -d <branch_name>

删除远程分支
git push origin --delete <branch_name>

-----------------------------------------------------------------------------------------------------------------------

良好的分支管理对于团队合作和项目维护至关重要。使用 Git 进行分支管理时，以下是一些最佳实践和策略，可以帮助你保持仓库整洁、可维护性高，并使团队协作更加顺畅。

### 1. **定义明确的分支策略**
   在团队中使用 Git 时，应该定义一个明确的分支策略，确保每个成员都遵循相同的规则。这些策略通常包括：
   
   - **主分支**（`master` 或 `main`）：该分支始终保持可发布状态，通常用于存放生产环境的代码。
   - **开发分支**（`develop`）：用于集成所有特性分支，在这个分支上进行集成测试。
   - **功能分支**（`feature` 分支）：每个新功能或任务创建一个新的分支。
   - **修复分支**（`hotfix` 或 `bugfix`）：用于紧急修复生产环境中的 bug。
   - **发布分支**（`release` 分支）：用于准备发布版本的分支。

   **示例分支模型**：
   - `master/main`（生产环境分支）
   - `develop`（开发集成分支）
   - `feature/<feature-name>`（功能分支）
   - `release/<version>`（发布分支）
   - `hotfix/<bug-id>`（修复分支）

   例如，Git Flow 模型就是一种广泛使用的分支策略。

### 2. **使用功能分支**（Feature Branches）
   每次开发一个新特性时，应该从 `develop`（或 `master`）分支切出一个新的功能分支，并在该分支上进行开发。功能分支的命名应简洁明了，描述所做的功能。例如：
   
   ```bash
   git checkout develop
   git checkout -b feature/add-login-page
   ```

   开发完成后，将功能分支合并回 `develop`（或者是 `master`，如果是紧急修复）。这样做有助于保持 `develop` 分支的干净，并且便于进行代码审查。

### 3. **定期合并更新**
   - 在开发过程中，`feature` 分支应该定期合并 `develop`（或 `master`）分支的最新代码，以避免后期合并时发生大规模冲突。例如：
   
     ```bash
     git checkout feature/add-login-page
     git fetch origin
     git merge origin/develop
     ```

   这样可以确保你的功能分支始终保持与主分支一致，并减少最后合并时的冲突。

### 4. **做好合并和拉取请求（Pull Requests）**
   使用 Pull Request（PR）进行代码审查是一个非常好的实践。开发完成后，将功能分支推送到远程仓库，然后通过 Pull Request 将其合并到 `develop` 或 `master` 分支。

   在合并 PR 时，确保：
   - **PR 的描述清晰**：写清楚该功能实现了什么，是否涉及到数据库变化，是否需要特别注意的地方等。
   - **PR 中的变更较小**：尽量保持 PR 的变更集中和简洁，一个 PR 不应该过于庞大。
   - **进行代码审查**：团队成员互相审查代码，检查逻辑、性能、安全性等方面的问题。

### 5. **使用语义化的版本控制（Semantic Versioning）**
   使用语义化版本控制（SemVer）帮助团队管理发布版本。版本号通常是 `MAJOR.MINOR.PATCH`，其规则如下：
   - **MAJOR**（主版本号）：当你做了不兼容的 API 修改，
   - **MINOR**（次版本号）：当你做了向后兼容的功能增强，
   - **PATCH**（修补版本号）：当你做了向后兼容的问题修正。

   **示例**：
   - 发布一个新版本时，使用 `v1.2.0`，表示一个包含新功能的版本。
   - 修复一个 bug 后，发布 `v1.2.1`，表示这是一个补丁版本。

   这样可以帮助团队和用户理解每次发布的内容和重要性。

### 6. **避免过多的临时分支**
   分支的使用应该有目的，不要创建大量的临时分支。临时分支（例如测试、调试分支）在完成后应及时删除，保持仓库的清晰。

   **删除分支**：
   - 删除本地分支：
     ```bash
     git branch -d <branch_name>
     ```
   - 删除远程分支：
     ```bash
     git push origin --delete <branch_name>
     ```

### 7. **做好冲突管理**
   当多个开发者在同一分支上工作时，可能会发生冲突。解决冲突时，遵循以下原则：
   - **理解冲突的根本原因**：不要草率地选择某一方的更改，应该理解两者更改的目的。
   - **合理拆分提交**：如果修改范围太大，尽量拆分成小的提交，这样可以在合并时更容易处理冲突。

   如果有冲突，Git 会标记冲突的文件，并且你可以手动解决冲突，然后使用：
   ```bash
   git add <file>
   git commit
   ```

### 8. **保持提交信息简洁清晰**
   良好的提交信息有助于团队成员理解每次提交的内容。提交信息通常遵循以下结构：
   - **标题**：简短明了，描述提交的内容（< 50 个字符）。
   - **正文**（可选）：详细描述为什么做这个更改，如何做的，以及可能影响的地方。

   示例：
   ```
   Add user authentication to login page

   - Added login form with username and password
   - Implemented validation logic
   - Linked login with backend API
   ```

### 9. **自动化工具与 CI/CD**
   在分支管理过程中，使用持续集成（CI）和持续部署（CD）工具可以帮助你自动化测试、构建和部署。常见的 CI/CD 工具如 Jenkins、GitHub Actions、GitLab CI 等，可以与 Git 分支管理集成，确保每个分支都经过自动化的测试和构建过程，减少人为错误。

### 10. **使用标签进行版本标记**
   每当发布新版本时，可以使用 Git 标签来标记版本，这样可以很方便地追溯到某个特定版本。

   **创建标签**：
   ```bash
   git tag v1.0.0
   ```

   **推送标签**：
   ```bash
   git push origin v1.0.0
   ```

   标签可以帮助你管理发布的版本，并与团队中的其他成员共享版本信息。

### 总结：
良好的分支管理可以显著提高团队开发的效率和代码质量。采取合适的分支策略、定期合并更新、使用 Pull Request 进行代码审查、避免过多的临时分支、保持清晰的提交信息，并结合 CI/CD 工具，能让开发过程更加顺畅和高效。

如果你有更具体的需求或问题，随时可以告诉我，我可以帮你进一步完善分支管理策略。


-----------------------------------------------------------------------------------------------------------------------

-- 在windows powershell 使用scoop(网址：https://scoop.sh/ )下载
   与更新migrate
   下载命令：scoop install migrate
   更新命令：scoop update migrate

-- 使用migrate更新数据库表结构（包括添加新表，添加外键）

migrate create -ext sql -dir db/migration -seq add_users

-----------------------------------------------------------------------------------------------------------------------

一条经验法则：永远不要将更改直接推送到主分支
当研究新功能时（working on new feature），我们应该从master创建一个新的独立分支
并且只有在正确测试和审查新代码后才将其合并回来
要创建一个新分支，我们运行：git checkout -b [分支名]

docker build -t simplebank:latest . 当前文件路径下有Dockerfile，根据Dockerfile
中的信息创建image并命名为simplebank:latest，"."意味着当前目录路径

docker rm [name]删除容器

docker rmi [name]删除image

docker container inspect [容器名] 查看对应容器的网络设置

docker run --name simplebank -p 8080:8080 -e GIN_MODE=release 
-e DB_SOURCE="postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable" 
simplebank:latest 在当前窗口运行image对应的容器，-e为环境变量参数，这里设置gin模式为release，
意味着软件正式对外发布，相对的是debug模式，另外一个环境变量是与数据库进程通信的链接，
simplebank项目通过viper读取.env文件中内容来加载常量，通过添加这个环境变量可以覆盖掉
.env文件中的DB_SOURCE关键字对应的常量

docker network ls 可以看到默认的桥接网络

docker network inspect bridge 可以看到关于这个网络的更多细节

docker create network [name] 创建一个网络

docker network connect bank-network [name] 将容器连接到bank-network网络

docker run --name simplebank --network bank-network -p 8080:8080 
-e GIN_MODE=release 
-e DB_SOURCE="postgresql://root:123456@postgres1:5432/simple_bank?sslmode=disable" 
simplebank:latest 这条命令的--network参数使得该容器链接到bank-work网络
这使得simplebank和postgres1在同一网络下，此时simplebank网络可以通过
容器名找到postgres1容器，需要将链接中的ip地址改为容器名

