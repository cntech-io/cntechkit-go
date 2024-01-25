### install

```bash
go get github.com/cntech-io/cntechkit-go/v2
```

## Methods

| Method                                                              | Description                                                                                                                          |
| ------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------ |
| env.NewServerEnv()                                                  | Loads predefined environment variables                                                                                               |
| env.GetBooleanEnv(key string, panicFlag bool)                       | Gets environment variable key value. if it exists returns value as boolean if not gives an explanation                         |
| env.GetNumberEnv(key string, panicFlag bool)                        | Gets environment variable key value. if it exists returns value as int if not gives an explanation                             |
| env.GetStringEnv(key string, panicFlag bool)                        | Gets environment variable key value. if it exists returns value as string if not gives an explanation                          |
| env.GetStringArrayEnv(key string, seperator string, panicFlag bool) | Gets environment variable key value. if it exists returns value as string array by separator if not exits gives a pretty explanation |
| logger.NewLogger(config \*LoggerConfig)                             | Creates logger with config (logrus)                                                                                                  |
| &nbsp;&nbsp;&nbsp;&nbsp;.Warn(msg string)                           | Warn log                                                                                                                             |
| &nbsp;&nbsp;&nbsp;&nbsp;.Info(msg string)                           | Info log                                                                                                                             |
| &nbsp;&nbsp;&nbsp;&nbsp;.Error(msg string)                          | Error log                                                                                                                            |