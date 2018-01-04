# timeutil
[![CircleCI](https://circleci.com/gh/azihsoyn/timeutil.svg?style=shield)](https://circleci.com/gh/azihsoyn/timeutil)

golang time util library

# Why? timeutil

## time comparison in go is complicated

check now between t1 and t2 in go standard time (t1 <= now <= t2)

```go
now := time.Now()
if !now.Before(t1) && !t2.Before(now) { // ????
  return true
}
return false
```

in timeutil

```go
now := time.Now()
if (timeutil.Time{now}).Between(t1, t2) {
  return true
}
return false
```

## ParseDuration cannot parse day or more
[https://github.com/golang/go/issues/17767]

time util can parse duration day and week

```
timeutil.ParseDuration("1d") // time.Hour * 24
timeutil.ParseDuration("1w") // time.Hour * 24 * 7
```

# LICENSE
MIT
