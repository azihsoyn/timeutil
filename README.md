# timex
[![CircleCI](https://circleci.com/gh/azihsoyn/timex.svg?style=shield)](https://circleci.com/gh/azihsoyn/timex)

golang time util library

# Why? timex

## time comparison in go is too simple

### Between

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

or more graceful

```go
now := time.Now()
if timex.X(now).Between().Y(t1).And().Z(t2) {
  return true
}
return false
```

### [Less|Greater]Than

```
t1 < t2
``

go standard time
```go
if t1.Before(t2) {
  return true
}
return false

if t2.After(t1) {
  return true
}
return false
```

in timex

```go
if (timex.Time{t1}).LessThan(t2) {
  return true
}
return false

if (timex.Time{t2}).GreaterThan(t1) {
  return true
}
return false
```

or more graceful

```go
if timex.A(t1).LessThan().B(t2) {
  return true
}
return false

if timex.A(t2).GreaterThan().B(t1) {
  return true
}
return false
```

### [Less|Greater]ThanEqual

```
t1 <= t2
``

go standard time
```go
if !t1.After(t2) { // or if t1.Before(t2) || t1.Equal(t2) {
  return true
}
return false

if(!t2.Before(t1)){ // or if t2.After(t1) || t2.Equal(t1) {
  return true
}
return false
```

in timex

```go
if (timex.Time{t1}).LessEqual(t2) {
  return true
}
return false

if (timex.Time{t2}).GreaterEqual(t1) {
  return true
}
return false
```

or more graceful

```go
if timex.A(t1).LessEqual().B(t2) {
  return true
}
return false

if timex.A(t2).GreaterEqual().B(t1) {
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
