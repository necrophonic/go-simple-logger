# go-simple-logger

```
  import "github.com/necrophonic/log"
```

Very simple logger for golang. Wrap default go [log](https://golang.org/pkg/log/) package to add granular functions.

```go

import "github.com/necrophonic/log"

// LevelTrace > LevelDebug > LevelWarn > LevelInfo > LevelError > LevelNone
log.Init(log.LevelInfo)

log.Trace("Trace this out")

log.Info("This is more informational")

```

Functions
---------

  - ``Fatal(format string, v ...interface{})``
  - ``Trace(format string, v ...interface{})``
  - ``Debug(format string, v ...interface{})``
  - ``Info(format string, v ...interface{})``
  - ``Warn(format string, v ...interface{})``
  - ``Error(format string, v ...interface{})``
