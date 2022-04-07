
### requirement

1. redis via etc/atest.yaml
2. mysql db via etc/atest.yaml

### steps

1. import atest.sql to db
2. [already done] generate or regenerate model and proto pb
3. run rpc server
 ` go run atest.go -f etc/atest.yaml `
4. run rpc client call main logic once 
    `go run cli/main.go`

### result

In the rpc server logs, 
we can see these prefix of logs step by step, 
[direct], [query-1],[update-1],[query-3], [query-4]

from [update 1], the value of type must be changed to a new value(random value)

but the result is not,  [query-3] log show the type is old, because it is select by id

#### eg

except:

```
[direct] type 50
[update-1] type 30
[query-3] type 30
[query-4] type 30
```

result:

```
[direct] type 50
[update-1] type 30
[query-3] type 50
[query-4] type 30
```




