## Concurrency 

- Being able to run code concurrently is becoming a larger part of programming as computers move from running a single stream of code faster to running more streams of code simultaneously. 
- To run programs faster, a programmer needs to design their programs to run concurrently, so that each concurrent part of the program can be run independently of the others.
- Two features in Go, `goroutines` and `channels`, make concurrency easier when used together. 
- `Goroutines` solve the difficulty of setting up and running concurrent code in a program.
- `Channels` solve the difficulty of safely communicating between the code running concurrently.

Intro 
  - In a modern computer, the processor, or CPU, is designed to run as many streams of code as possible at the same time.
  - These processors have one or more “cores,” each capable of running one stream of code at the same time.
  - So, the more cores a program can use simultaneously, the faster the program will run. 
  - However, in order for programs to take advantage of the speed increase that multiple cores provide, programs need to be able to be split into multiple streams of code.
  - A `goroutine` is a special type of function that can run while other `goroutines` are also running.
  - The power goroutines provide is that each goroutine can run on a processor core at the same time.
  - If your computer has four processor cores and your program has four goroutines, all four goroutines can run simultaneously. 
  - When multiple streams of code are running at the same time on different cores like this, it’s called running in parallel.



## Context package in GOLANG

- When developing a large application, especially in server software, sometimes it’s helpful for a function to know more about the environment it’s being executed in aside from the information needed for a function to work on its own
- For example, if a web server function is handling an HTTP request for a specific client, the function may only need to know which URL the client is requesting to serve the response. The function might only take that URL as a parameter. 
- However, things can always happen when serving a response, such as the client disconnecting before receiving the response. If the function serving the response doesn’t know the client disconnected, the server software may end up spending more computing time than it needs calculating a response that will never be used.
- In this case, being aware of the context of the request, such as the client’s connection status, allows the server to stop processing the request once the client disconnects.
- This saves valuable compute resources on a busy server and frees them up to handle another client’s request.
- This type of information can also be helpful in other contexts where functions take time to execute, such as making database calls. To enable ubiquitous access to this type of information, Go has included a context package in its standard library.

- A way to think about context package in go is that it allows you to pass in a “context” to your program. Context like a timeout or deadline or a channel to indicate stop working and return. For instance, if you are doing a web request or running a system command, it is usually a good idea to have a timeout for production-grade systems. Because, if an API you depend on is running slow, you would not want to back up requests on your system, because, it may end up increasing the load and degrading the performance of all the requests you serve. Resulting in a cascading effect. This is where a timeout or deadline context can come in handy.




### context.Background() ctx Context
- This function returns an empty context. This should be only used at a high level (in main or the top level request handler). This can be used to derive other contexts that we discuss later.

- `ctx, cancel := context.Background()`

### context.TODO() ctx Context
- This function also creates an empty context. This should also be only used at a high level or when you are not sure what context to use or if the function has not been updated to receive a context. Which means you (or the maintainer) plans to add context to the function in future.

- `ctx, cancel := context.TODO()`


### context.WithValue(parent Context, key, val interface{}) (ctx Context, cancel CancelFunc)
- This function takes in a context and returns a derived context where value val is associated with key and flows through the context tree with the context. This means that once you get a context with value, any context that derives from this gets this value. It is not recommended to pass in critical parameters using context value, instead, functions should accept those values in the signature making it explicit.

- `ctx := context.WithValue(context.Background(), key, "test")`


### context.WithCancel(parent Context) (ctx Context, cancel CancelFunc)

- This is where it starts to get a little interesting. This function creates a new context derived from the parent context that is passed in. The parent can be a background context or a context that was passed into the function.

- This returns a derived context and the cancel function. Only the function that creates this should call the cancel function to cancel this context. You can pass around the cancel function if you wanted to, but, that is highly not recommended. This can lead to the invoker of cancel not realizing what the downstream impact of canceling the context may be. There may be other contexts that are derived from this which may cause the program to behave in an unexpected fashion. In short, NEVER pass around the cancel function.

- `ctx, cancel := context.WithCancel(context.Background())`



### context.WithDeadline(parent Context, d time.Time) (ctx Context, cancel CancelFunc)
- This function returns a derived context from its parent that gets cancelled when the deadline exceeds or cancel function is called. For example, you can create a context that will automatically get canceled at a certain time in future and pass that around in child functions. When that context gets canceled because of deadline running out, all the functions that got the context get notified to stop work and return.

- `ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2 * time.Second))`




### context.WithTimeout(parent Context, timeout time.Duration) (ctx Context, cancel CancelFunc)
- This function is similar to context.WithDeadline. The difference is that it takes in time duration as an input instead of the time object. This function returns a derived context that gets canceled if the cancel function is called or the timeout duration is exceeded.

- `ctx, cancel := context.WithTimeout(context.Background(), time.Duration(150)*time.Millisecond)`

