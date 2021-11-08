# Start [ent][]

[ent]: https://github.com/ent/ent

See [Quick Introduction](https://entgo.io/docs/getting-started/) for details. Offical Code: [Getting Started Example](https://github.com/ent/ent/tree/master/examples/start).

<!--
cd start-golang/_/start-ent
go mod init github.com/ikuokuo/start-golang/_/start-ent
go get -u entgo.io/ent/cmd/ent
go get -u github.com/mattn/go-sqlite3

go mod edit -replace github.com/ikuokuo/start-golang/_/start-ent/ent=./ent
-->

```zsh
git clone --depth 1 https://github.com/ikuokuo/start-golang.git
cd start-golang/_/start-ent
go mod tidy
```

```zsh
❯ go run start.go
2021/11/08 10:50:00 user was created:  User(id=1, age=30, name=a8m)
2021/11/08 10:50:00 user returned:  User(id=1, age=30, name=a8m)
2021/11/08 10:50:00 car was created:  Car(id=1, model=Tesla, registered_at=Mon Nov  8 10:50:00 2021)
2021/11/08 10:50:00 car was created:  Car(id=2, model=Ford, registered_at=Mon Nov  8 10:50:00 2021)
2021/11/08 10:50:00 user was created:  User(id=2, age=30, name=a8m)
2021/11/08 10:50:00 returned cars: [Car(id=1, model=Tesla, registered_at=Mon Nov  8 10:50:00 2021) Car(id=2, model=Ford, registered_at=Mon Nov  8 10:50:00 2021)]
2021/11/08 10:50:00 Car(id=2, model=Ford, registered_at=Mon Nov  8 10:50:00 2021)
2021/11/08 10:50:00 car "Tesla" owner: "a8m"
2021/11/08 10:50:00 car "Ford" owner: "a8m"
2021/11/08 10:50:00 The graph was created successfully
2021/11/08 10:50:00 cars returned: [Car(id=3, model=Tesla, registered_at=Mon Nov  8 10:50:00 2021) Car(id=4, model=Mazda, registered_at=Mon Nov  8 10:50:00 2021)]
2021/11/08 10:50:00 cars returned: [Car(id=3, model=Tesla, registered_at=Mon Nov  8 10:50:00 2021) Car(id=5, model=Ford, registered_at=Mon Nov  8 10:50:00 2021)]
2021/11/08 10:50:00 groups returned: [Group(id=1, name=GitLab) Group(id=2, name=GitHub)]
```

## See also

[Blog](https://entgo.io/blog/):

- [Generating OpenAPI Specification with Ent](https://entgo.io/blog/2021/09/10/openapi-generator)
  - [使用 Ent 生成 OpenAPI 规范](https://entgo.io/zh/blog/2021/09/10/openapi-generator)
