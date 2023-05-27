HTTP Protocol:
	The HTTP (Hypertext Transfer Protocol) is an application-level protocol widely used for communication between web browsers and servers. It defines how messages are formatted and transmitted, as well as the actions that the client and server 
should take in response to these messages

HTTP operates in a client-server model, where the client (typically a web browser) sends HTTP requests to the server, and the server responds with HTTP responses. These requests and responses are exchanged using a standardized
format called HTTP messages.


HTTP messages consist of a start line, headers, and an optional message body. The start line includes the HTTP method (such as GET, POST, PUT, DELETE), the target resource (URL or URI), and the HTTP version.

Commonly used HTTP methods include:
	GET: Retrieves a representation of the specified resource.
	POST: Submits data to be processed to the specified resource.
	PUT: Updates the specified resource with the provided data.
	DELETE: Deletes the specified resource.
	PATCH: Partially updates the specified resource.
	
HTTP headers provide additional information about the request or response, such as content type, content length, caching directives, and more.

HTTP supports various features, including:
	Stateless: Each HTTP request is independent and does not require the server to maintain any client-specific information between requests. This allows for scalability and simplicity.
	Caching: HTTP supports caching mechanisms that allow responses to be stored and reused to improve performance.
	Authentication and Security: HTTP provides mechanisms for authentication and secure communication, such as HTTPS (HTTP over TLS/SSL).
	Redirects: HTTP supports redirecting clients to different URLs or resources.
	Cookies: HTTP allows for the use of cookies to maintain state between requests.


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
