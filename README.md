# tcphub

tcphub is program that connects to a TCP server and serve the same stream to its own TCP server.

Nothing more, nothing less.

It's just used by me for tests.

# Example

docker-compose.yml : 

```
version: '3.2'

services:
    tcphub:
        image: sgaunet/tcphub:latest
        environment: 
            - SERVER=example.test:65432
            - PORT=65433
        ports:
            - 65433:65433

    tcphub2:
        image: sgaunet/tcphub:latest
        environment: 
            - SERVER=tcphub:65433
            - PORT=65433
        ports:
            - 65434:65433
```