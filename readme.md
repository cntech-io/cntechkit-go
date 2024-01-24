### Common methods for projects

| Method                  | Description                                                                                                                          |
| ----------------------- | ------------------------------------------------------------------------------------------------------------------------------------ |
| utils.GetBooleanEnv(key string, panicFlag bool)     | Gets environment variable key value. if it exists returns value as boolean if not gives a pretty explanation                         |
| utils.GetNumberEnv(key string, panicFlag bool)      | Gets environment variable key value. if it exists returns value as int if not gives a pretty explanation                             |
| utils.GetStringEnv(key string, panicFlag bool)      | Gets environment variable key value. if it exists returns value as string if not gives a pretty explanation                          |
| utils.GetStringArrayEnv(key string, seperator string, panicFlag bool) | Gets environment variable key value. if it exists returns value as string array by separator if not exits gives a pretty explanation |
| NewLogger(config *LoggerConfig) | Creates logger with config (logrus) |
| logger.Warn(msg string) | Warn log |
| logger.Info(msg string) | Info log |
| logger.Error(msg string) | Error log |
| NewServerEnv() | Loads predefined environment variables |
