# Recov Goroutines 

## about

This repo is a learning purposed examples to captured goroutine panic.

In Go programing language, goroutines has no relationships like parent-child. Each goroutine is running independency of other goroutines. If one goroutines doesn't handle the occuring panic, it will cause the entire process to end.

So, the panic needs to be captured within the goroutines to ensure the service is available.

## use 
It can be a little component of workpool(goroutine pools) to keep goroutine safe.

Init recover component with Retry times, panic callback function, workerfunc and then let it run.

When pannic happend, it will stdout the error, and recover. You can use panic callback func to capture the pannic and send to Send to your application monitor.

```go
    r := rv.RecovRoutine{
		RetryCnt: 1,
		PCallback: func(err error) {
			if err != nil {
				println(err.Error())
                //.... handle it 
                //....
			}
		},
	}
	r.Worker = func() {
		var slice []int
		slice[2] = 1
	}
	r.Run()
	// Waiting for retry.
	time.Sleep(10 * time.Second)
```
