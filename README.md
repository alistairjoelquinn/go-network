Social Network Go Server
---

![User profile image](https://res.cloudinary.com/dtirfwiy8/image/upload/v1644506343/go-network/profile_jojn4h.png)

In the final weeks at Spiced Academy students build a full stack social network application using React and Node / Express.js. Here is an alternative solution to the server API routes written in Go instead of Node. API routes have been written using Fiber due to it's similarity to Express. 

![User Login](https://res.cloudinary.com/dtirfwiy8/image/upload/v1644506345/go-network/login_tjqzms.png)

Bcrypt is still being used for handling user password encryption. Once logged in or registered, unlike the Node server, JSON web tokens are being set in place of Node's cookie-session middleware.

![User search](https://res.cloudinary.com/dtirfwiy8/image/upload/v1644506343/go-network/search_odxcel.png)

In terms of functionality, routes are largely similar to their Node counterparts. The go AWS SDK has a very similar API to that in Node when handling image uploads or using the email service. 

![User password reset](https://res.cloudinary.com/dtirfwiy8/image/upload/v1644506343/go-network/reset_vevgrg.png)

TODO: the final part of the project uses web sockets for instant chat. In the existing project, socket.io is being used for handling client - server communication. In the absence of a good server side socket.io v3 implementation, the decision has been made to re-write the existing client side code to utilise the native browser web socket API.

![Chat room](https://res.cloudinary.com/dtirfwiy8/image/upload/v1644506343/go-network/chat_htqawy.png)
