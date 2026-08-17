# gocline

### Nx run many projects

```bash
nx run-many --target=serve --projects=user-svc,bff-svc --parallel=2
```

### Listing the methods

```bash
grpcurl -plaintext localhost:50051 list user.v1.UserService
```

### Testing the endpoint using grpcurl

```bash
grpcurl -plaintext -d '{"name":"Arsalan"}' localhost:50051 user.v1.UserService.CreateUser
```
