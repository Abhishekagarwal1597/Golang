HTTP Protocol:
	The HTTP (Hypertext Transfer Protocol) is an application-level protocol widely used for communication between web browsers and servers. It defines how messages are formatted and transmitted, as well as the actions that the client and server 
should take in response to these messages

	HTTP operates in a client-server model, where the client (typically a web browser) sends HTTP requests to the server, and the server responds with HTTP responses. These requests and responses are exchanged using a standardized
format called HTTP messages.


REST (Representational State Transfer) is an architectural style for designing networked applications. It defines a set of principles and constraints for creating web services that are scalable, stateless, and interoperable. 
	REST is often implemented over the HTTP protocol due to its alignment with the principles of REST.

Context:
  The context package is used for managing and propagating cancellation signals, deadlines, and other request-scoped values across Goroutines (concurrently executing functions).
  The context package provides the Context type, which represents the execution context of a request or operation. It allows you to control the lifecycle and behavior of concurrent operations, enabling the propagation of cancellation signals and 
  deadlines to Goroutines involved in the operation.The context package provides the Context type, which represents the execution context of a request or operation. It allows you to control the lifecycle and behavior of concurrent operations, 
  enabling the propagation of cancellation signals and deadlines to Goroutines involved in the operation
  
 Some common use cases for the context package in Go include:

1- Cancellation: You can cancel a context to propagate a cancellation signal to Goroutines, allowing them to terminate gracefully.
2- Deadlines: You can set deadlines for operations, ensuring that they complete within a certain time frame. If the deadline is exceeded, the context is canceled.
3- Value propagation: You can pass request-scoped values through the context, allowing Goroutines to access them without having to pass them explicitly through function parameters.v
