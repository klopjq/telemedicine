# Dummy Server 

```
# How to execute 
truncate -s0 /tmp/request.log | go run main.go | tail -f /tmp/request.log 
```