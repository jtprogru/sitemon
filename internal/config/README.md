### RUN

Create file `main.go`

```golang
package main

import (
    "fmt"
    "log"

    "github.com/jtprogru/sitemon/internal/config"
)

func main() {
    cfg, err := config.BuildConfig("./configs/config.yml")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%+v", cfg)
}

```

```shell
[welcome sitemon]$ SITEMON_LOG_LEVEL=INFO SITEMON_TGTOKEN=qwe SITEMON_SENTRYDSN=sentry:dsn SITEMON_TGCHAT=-123 go run
main.go
```