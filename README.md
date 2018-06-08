# logging with context
  An implementation on top of golang log framework to add context(requestID, clientID, userID) in each log statement with log level (Info, Error, Debug).
  
#  Problem:
  The logger that golang provides (https://golang.org/pkg/log/) has nothing that a logger in real life services should have. This is an attempt at making logging in golang useful for better debugging.
  
  To give you a brief insight on what basic necessity I'm trying to solve here, consider following code with go logger-
  
    package main

    import (
      "log"
    )

    func main() {
      log.Printf("test info log formatted with var %v, var %v ", 1, "two")
      doSomethingError()
      doSomethingDebug()
    }

    func doSomethingError() {
      log.Printf("test error log formatted with var %v, var %v ", 1, "two")
    }

    func doSomethingDebug() {
      log.Printf("test debug log formatted with var %v, var %v ", 1, "two")
    }
    
  This prints-
  
    2018/06/08 18:28:49 test info log formatted with var 1, var two 
    2018/06/08 18:28:49 test error log formatted with var 2, var three 
    2018/06/08 18:28:49 test debug log formatted with var one, var 2 
    
  But then you deploy your code to production, you look at your logs and see this sort of thing:
  
    2018/06/08 18:28:49 test info log formatted with var 1, var two 
    2018/06/08 18:28:49 test info log formatted with var 1, var two 
    2018/06/08 18:28:50 test error log formatted with var 2, var three 
    2018/06/08 18:28:50 test debug log formatted with var one, var 2 
    2018/06/08 18:28:50 test error log formatted with var 2, var three 
    2018/06/08 18:28:51 test debug log formatted with var one, var 2 

    
  Mr. Akhil reports something went wrong. But because your program is a high performance parallelized wonder, you can’t be sure which line relates to his request. There is just no link between logs, neither is there any INFO/ERROR/DEBUG tag to know something went wrong in first look.

# Solution
  Use some simple implementation like in this project. Just copy directories apicontext and logger in your project, and you can write the same code above like-
  
    package main

    import (
      "context"
      "log-context/apicontext"
      "log-context/logger"
    )

    func main() {
      ctx := apicontext.WithReqID(context.Background())
      logger.Info(ctx, "test info log formatted with var %v, var %v ", 1, "two")
      doSomethingError(ctx)
      doSomethingDebug(ctx)
    }

    func doSomethingError(ctx context.Context) {
      logger.Error(ctx, "test error log formatted with var %v, var %v ", 1, "two")
    }

    func doSomethingDebug(ctx context.Context) {
      logger.Debug(ctx, "test debug log formatted with var %v, var %v ", 1, "two")
    }

  To give you output like-
  
    2018/06/08 18:32:19 [88153e20-efe3-40cb-ad68-4ba8f6a1e19a:0:0] main.go:11  INFO: test info log formatted with var 1, var two 
    2018/06/08 18:32:19 [b000a173-713b-430b-8815-6ac5cfec50ba:0:0] main.go:11  INFO: test info log formatted with var 1, var two 
    2018/06/08 18:32:19 [88153e20-efe3-40cb-ad68-4ba8f6a1e19a:0:0] main.go:17  ERROR: test error log formatted with var 2, var three 
    2018/06/08 18:32:19 [88153e20-efe3-40cb-ad68-4ba8f6a1e19a:0:0] main.go:21  DEBUG: test debug log formatted with var one, var 2 
    2018/06/08 18:32:20 [b000a173-713b-430b-8815-6ac5cfec50ba:0:0] main.go:17  ERROR: test error log formatted with var 2, var three 
    2018/06/08 18:32:20 [b000a173-713b-430b-8815-6ac5cfec50ba:0:0] main.go:21  DEBUG: test debug log formatted with var one, var 2 
    
  As you can see, just by first look, you are able to get the link between different logs and are able to figure out if it's INFO/ERROR/DEBUG. Simple.
