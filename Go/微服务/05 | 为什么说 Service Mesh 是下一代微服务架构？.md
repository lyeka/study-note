# 05 | 为什么说 Service Mesh 是下一代微服务架构？



## 什么是 Service Mesh 

Service Mesh 是用于处理服务到服务通信的专用基础架构层。云原生有着复杂的服务拓扑，它负责可靠的传递请求。实际上，Service Mesh 通常是作为一组轻量级网络代理实现，这些代理与应用程序代码部署在一起，应用程序无感知。



![Drawing 2.png](https://s0.lgstatic.com/i/image/M00/32/92/Ciqc1F8OnsOAa3MVAABY1memBaA509.png)





## 为什么需要Service Mesh

Service mesh 的出现源于微服务的发展。

微服务需要的基础功能

![Drawing 0.png](https://s0.lgstatic.com/i/image/M00/32/9C/CgqCHl8OniqAaOpTAABLWy0eR68344.png)

背后诉求

- 功能需求（如上图）
- 解耦业务代码与基础组件代码





## Service Mesh 开源组件

### Istio

### Linkerd

### Envoy



### 疑问

1. Istio与Envoy的关系？
2. Service Mesh 只服务k8s？
3. 什么是Api网关？其与Service Mesh的区别？