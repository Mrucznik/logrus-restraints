# logrus restraints

Log with [logrus](https://github.com/sirupsen/logrus) with restraints on frequency of logging.

## Examples

```go
logrus_restraints.WithTTL(200 * time.Millisecond).
	Infoln("this message will be logged not more often than every 200ms")
```

```go
logrus_restraints.EveryN(5).Infoln("this message will be logged every 5 calls")
```