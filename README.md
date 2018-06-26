# logging with context
  An implementation on top of golang log package to add context(requestID, clientID, userID) in each log statement with log level (Info, Error, Debug).
  
#  Problem
  The logger that golang provides (https://golang.org/pkg/log/) has nothing that a logger in real life services should have. This is an attempt at making logging in golang useful for better debugging.
  
  To give you a brief insight on what basic necessity I'm trying to achieve here, consider following code with golang log package-
  
    package main

    import (
      "log"
    )

    func main() {
      log.Printf("test log with var %v, var %v", 1, "two")
      doSomethingError()
      doSomethingDebug()
    }

    func doSomethingError() {
      log.Printf("test log with var %v, var %v", 2, "three")
    }

    func doSomethingDebug() {
      log.Printf("test log with var %v, var %v", "one", 2)
    }
    
  This prints-
  
    2018/06/08 18:28:49 test log with var 1, var two
    2018/06/08 18:28:49 test log with var 2, var three
    2018/06/08 18:28:49 test log with var one, var 2
    
  But then you deploy your code to production, you look at your logs and see this sort of thing:
  
    2018/06/08 18:28:49 test log with var 1, var two
    2018/06/08 18:28:49 test log with var 1, var two
    2018/06/08 18:28:50 test log with var 2, var three
    2018/06/08 18:28:50 test log with var one, var 2
    2018/06/08 18:28:50 test log with var 2, var three
    2018/06/08 18:28:51 test log with var one, var 2

    
  Mr. Akhil reports something went wrong. But because your program is a high performance parallelized wonder, you can’t be sure which line relates to his request. There is just no link between logs, neither is there any INFO/ERROR/DEBUG tag to know something went wrong.

# Solution
  Use a simple implementation presented in here. Just copy directories apicontext and logger in your project root, and you can write the same code above like-
  
    package main

    import (
      "context"
      "log-context/apicontext"
      "log-context/logger"
    )

    func main() {
      ctx := apicontext.WithReqID(context.Background())
      logger.Info(ctx, "test log with var %v, var %v", 1, "two")
      doSomethingError(ctx)
      doSomethingDebug(ctx)
    }

    func doSomethingError(ctx context.Context) {
      logger.Error(ctx, "test log with var %v, var %v", 1, "two")
    }

    func doSomethingDebug(ctx context.Context) {
      logger.Debug(ctx, "test log with var %v, var %v", 1, "two")
    }

  To give you output like-
  
    2018/06/08 18:32:19 [88153e20-efe3-40cb-ad68-4ba8f6a1e19a:0:0] main.go:11  INFO: test log with var 1, var two
    2018/06/08 18:32:19 [b000a173-713b-430b-8815-6ac5cfec50ba:0:0] main.go:11  INFO: test log with var 1, var two
    2018/06/08 18:32:19 [88153e20-efe3-40cb-ad68-4ba8f6a1e19a:0:0] main.go:17  ERROR: test log with var 2, var three
    2018/06/08 18:32:19 [88153e20-efe3-40cb-ad68-4ba8f6a1e19a:0:0] main.go:21  DEBUG: test log with var one, var 2
    2018/06/08 18:32:20 [b000a173-713b-430b-8815-6ac5cfec50ba:0:0] main.go:17  ERROR: test log with var 2, var three
    2018/06/08 18:32:20 [b000a173-713b-430b-8815-6ac5cfec50ba:0:0] main.go:21  DEBUG: test log with var one, var 2
    
  As you can see, just by one look, you are able to spot the link between different logs and are able to figure out if it's INFO/ERROR/DEBUG. Simple.
  
# P.S.
  1. clientID and userID are printed as 0:0 in above logs as I haven't set them. The methods to set them are provided in `apicontext/context.go` though. If you just want to log requestID, then just edit `ctxString` in `logger/logging.go` to remove them.
  2. With this implementation, you need to pass argument `ctx` to each method being called, so that you can print reqID:clientID:userID. I personally searched a lot over internet, but couldn't find any implementation to print context in log without actually passing it in each method. However, if YOU do find it, poke me :)
  3. The logging format that suited me is `DATE TIME [requestID:clientID:userID] fileName:lineNo logLevel: msgToPrint`. But you can play around in file `logger/logging.go` to get your combination right, the code is pretty straightforward.
