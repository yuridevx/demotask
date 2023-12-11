# Structure
Domain - package to store golang code generated from proto files
Game - Service that generates random numbers
Numbers - The game itself, including frontend and golang static file serving
Proto - protocol files for grpc
Integrations - integration tests for the game
Frontend - self-explanatory

# How to scale it for 10000+

1. Scale random number generator service, use load balancer on client side.
2. Move frontend to CDN + Storage
3. Write RedisGameManager that will store game state in Redis
4. Deploy multiple instances of Numbers service, use load balancer on client side.
5. Actually 10000 will work with just one instance. It should be way more than 10000 to overload that.

# Design approach used principles

1. Decoupling
2. Inversion of control
3. Graceful shutdown
4. High cohesion, loose coupling + Single source of truth (Meaningful class structure)
5. MVC partially for numbers service.