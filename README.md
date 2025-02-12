# Posts

Сервис создания постов и комментариев с использованием GraphQL

## Запуск

Для запуска используйте команду `make start`.
Это поднимет Docker-контейнер вместе с PostgreSQL.

Для запуска тестов используйте`make test`

Если нужно остановить контейнеры, то можно использовать:
1) `make down` - удалит контейнеры 
2) `make stop` - остановит контейнеры 

### Выбор хранилища

Хранилище можно выбрать в конфиге.
Это делается в файле [stage.yaml](./configure/stage.yaml) (для запуска в Docker).
Чтобы использовать in-memory хранилище, надо в поле IN_MEMORY поставить true.
В противном случае, будет использоваться postgres

## API
API описано в graphqls файлах в папке [graph](./graph).

## Запросы
### Queries

- Список постов с пагинацией БЕЗ комментариев.
- Подробная информация о посте с комментариями первого уровня.
- Получение одного уровня ответов на конкретный комментарий.

Проблема большой вложенности решается тем, что ответы на комментарии подгружаются по 1 уровню.
На мой взгляд, это оптимальнее, чем получать полное дерево комментариев, поскольку оно может быть избыточно.
Так мы даем клиенту возможность решать, сколько уровней он хочет видеть.

### Mutations

- Создание поста
- Создание комментария на конкретный пост ИЛИ комментария-ответа

В запросе создания поста можно указать возможность оставлять комментарии.
При создании комментария учитывается эта проверка.

Комментарий создается для конректного поста. Если у него поле ReplyTo не заполнено, то это комментарий 1 уровня.
В противном случае - это ответ на другой комментарий уровнем ниже.

### Subscriptions

Можно подписаться на комментарии к конкретному посту.

Для управления подписками создан CommentObserver который может создавать слушателя, удалять его,
и оповещать всех подписанных слушателей.

Слушатели - это структура из канала для комментариев и идентификатора слушателя.
При закрытии контекста подключения слушатель удаляется из хранилища.

## Тесты

Есть unit-тесты для сервисного слоя. Моки генерировал с помощью [gomock](https://github.com/golang/mock).

## Примеры запросов

### Создание поста

```graphql
mutation CreatePost {
    CreatePost(
        post: {
            name: "Cool title!!"
            content: "Meow"
            commentsAllowed: true
        }
    ) {
        id
        createdAt
        name
        content
    }
}

```

### Получение списка постов

```graphql
query GetAllPosts {
    GetAllPosts(page: 1, pageSize: 5) {
        id
        createdAt
        name
        content
    }
}

```

### Получение подробной информации о конкретном посте по Id

```graphql
query GetPostById {
    GetPostById(id: 1) {
        id
        createdAt
        name
        content
        commentsAllowed
        comments(page: 1, pageSize: 5) {
            id
            createdAt
            content
            post
        }
    }
}

```

### Создание комментария 1 уровня

```graphql
mutation CreateComment {
    CreateComment(input: { content: "cool", post: "1"}) {
        id
        createdAt
        content
        post
        replyTo
    }
}

```

### Создание ответа на комментарий

```graphql
mutation CreateComment {
    CreateComment(input: { content: "you`re wrong....", post: "1", replyTo: "1"}) {
        id
        createdAt
        content
        post
        replyTo
    }
}
```

### Получение одного уровня ответов на комментарий

```graphql
query GetReplies {
    GetReplies(commentId:"1") {
        id
        createdAt
        content
    }
}
```

### Подписка на комментарии поста

```graphql
subscription CommentsSubscription {
    CommentsSubscription(postId: "1") {
        id
        createdAt
        content
        post
        replyTo
    }
}
```
