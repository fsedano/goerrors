# Example on how to handle Custom Golang errors

```
❯ go run .
2023/09/29 12:49:52 Opening db nonexistingdb
2023/09/29 12:49:52 It didnt exist: things failed: db does not exist : We failed because db did not exist
2023/09/29 12:49:52 Opening db notpermissionsdb
2023/09/29 12:49:52 We didnt have permissions: things failed: db does not exist : We failed because permissions were not there
2023/09/29 12:49:52 Opening db otherbrokenthings
2023/09/29 12:49:52 It was a db err!: things failed: db error: We failed in a non specific way
```