apps:
  - name: adminserver
    port: 60004
    nodePort: 30970
  - name: webserver
    ports:
      - name: http
        port: 8090
        nodePort: 32162
      - name: extend
        port: 8081
        nodePort: 32163
  - name: apiserver
    port: 8080
    nodePort: 31921
  - name: coreservice
    port: 50009
    nodePort: 32001
    cacheAffinity: true
  - name: toposerver
    port: 60002
  - name: hostserver
    port: 60001
    nodePort: 32357
  - name: operationserver
    port: 60011
  - name: cacheservice
    port: 50010
  - name: cloudserver
    port: 60013
  - name: eventserver
    port: 60009
  - name: procserver
    port: 60003
  - name: taskserver
    port: 60012

image:
  repository: harbor.dev.21vianet.com/cmdb/
  tag: {{.version}}
#  tag: "latest"
env:
#  pullPolicy: IfNotPresent
  pullPolicy: Always

# cache coresvc 同一node提高效率
cache:
  enabled: true
  labels:
    - name: app.kubernetes.io/name
      value: redis
    - name: app.kubernetes.io/component
      value: master
#res:
#  req:
#    mem: 1024Mi
#    cpu: 600m
#  req2:
#    mem: 2048Mi
#  req3:
#    mem: 3072Mi
#  ave:
#    mem: 1024Mi
#    cpu: 1000m