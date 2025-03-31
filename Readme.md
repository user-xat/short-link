# ShortLink
This repository contains a pet project in the Golang language. A web server and a backend service are implemented, which are connected via gRPC. The backend uses the Redis NoSQL database as storage. The frontend part uses the Memcached database as a cache. The image below shows the architecture of the application.

![shortlink architecture](https://downloader.disk.yandex.ru/preview/097753c43b7794a407438a358a2d1c6a678d6dc1b05ad55a5ad6cd222e5b4806/67eaf43d/VVz-FK0GKMvH50lt3LVT1DAZcwTxPUthKDLa6arWoZqf0Lf9UAvkxQnu5_V9090LoqmmpQTvZUNH-GuRgi_YTg%3D%3D?uid=0&filename=shortlink-main-main%20%281%29.png&disposition=inline&hash=&limit=0&content_type=image%2Fpng&owner_uid=0&tknv=v2&size=2048x2048)

Docker files have also been written to build docker images, as well as a compose file.yml for deploying the entire application using Docker.
```sh
docker compose -f compose.prod.yml up
```

Deployment and service files have been written for the Kubernetes service.