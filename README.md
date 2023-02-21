# gointerval

execute a function in an interval, cancel anytime. (if you just want to call soemthing once after a delay, use [time.AfterFunc](https://pkg.go.dev/time#AfterFunc) )

## versions

* v0.1.0 initial release

## example

```go
func main() {
  var runImmediately bool = true
  var intervalDuration time.Duration = time.Second*1
  // the function's return value has to be a bool and returning anything but true will stop the interval
  var myInterval *GoInterval = NewInterval(yourFunction, intervalDuration, runImmediately)
  time.Sleep(time.Second*3)
  myInterval.Stop()
  willBeFalse := myInterval.State()
  time.Sleep(time.Second*3)
  myInterval.Start()
  willBeTrue := myInterval.State()
}

func YourFunction() bool {
  fmt.Println("tick tick tick")
  // if you return false, the interval is stopped!
  return true
}
```
