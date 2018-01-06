# timex
[![CircleCI](https://circleci.com/gh/azihsoyn/timex.svg?style=shield)](https://circleci.com/gh/azihsoyn/timex)

golang time util library

# Why? timex

## time comparison in go is complicated

check now between t1 and t2 in go standard time (t1 <= now <= t2)

```go
now := time.Now()
if !now.Before(t1) && !t2.Before(now) { // ????
  return true
}
return false
```

in timex

```go
now := time.Now()
if (timex.Time{now}).Between(t1, t2) {
  return true
}
return false
```

## ParseDuration cannot parse day or more
[https://github.com/golang/go/issues/17767]

time util can parse duration day and week

```
timex.ParseDuration("1d") // time.Hour * 24
timex.ParseDuration("1w") // time.Hour * 24 * 7
```

# LICENSE
MIT
