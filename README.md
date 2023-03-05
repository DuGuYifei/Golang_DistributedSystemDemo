# README

This is a distributed system demo using golang. It contains service register, user portal, log service, transaction service.

It use **Service Discovery** to finish service register, add/remove patched services. For example, here it use register service to add a log service to a service which record grade of student.

It use **Heartbeat** as method of **Health Check**.

code in branch [go_src](https://github.com/DuGuYifei/Golang_DistributedSystemDemo/tree/go_src)

![](2023-03-06-02-50-42.png)

![](2023-03-06-02-35-30.png)

![](2023-03-06-02-36-02.png)

Learn from tutorial [source](https://www.bilibili.com/video/BV1ZU4y1577q)

Part about front-end(portal) is download from resource directly.

## main component

- service register
  - service register
  - health check
- user portal
  - web app
  - api gateway
- log service
  - centralized log
- business service
  - business logic
  - data persistence

![](2023-03-04-18-03-07.png)

### register service e.g. log

log service can be add as a patch to other business services.

![](2023-03-05-17-32-13.png)