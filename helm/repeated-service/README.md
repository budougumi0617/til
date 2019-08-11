# helm results

```yaml
$ helm tiller run -- helm install --debug --dry-run --name sample .
Installed Helm version v2.14.3
Installed Tiller version v2.14.3
Helm and Tiller are the same version!
Starting Tiller...
Tiller namespace: kube-system
Running: helm install --dry-run --name sample .

[debug] SERVER: "127.0.0.1:44134"

[debug] Original chart version: ""
[debug] CHART PATH: /Users/budougumi0617/go/src/github.com/budougumi0617/til/helm/repeated-service

NAME:   sample
REVISION: 1
RELEASED: Sun Aug 11 19:28:37 2019
CHART: repeated-service-0.1.0
USER-SUPPLIED VALUES:
{}

COMPUTED VALUES:
affinity: {}
fullnameOverride: ""
image:
  pullPolicy: IfNotPresent
  repository: nginx
  tag: stable
imagePullSecrets: []
ingress:
  annotations: {}
  enabled: false
  hosts:
  - host: chart-example.local
    paths: []
  tls: []
nameOverride: ""
nodeSelector: {}
replicaCount: 1
resources: {}
servers:
- alice
- bob
- charlie
- dave
service:
  port: 80
  type: ClusterIP
tolerations: []

HOOKS:
---
# sample-repeated-service-test-connection
apiVersion: v1
kind: Pod
metadata:
  name: "sample-repeated-service-test-connection"
  labels:
    app.kubernetes.io/name: repeated-service
    helm.sh/chart: repeated-service-0.1.0
    app.kubernetes.io/instance: sample
    app.kubernetes.io/version: "1.0"
    app.kubernetes.io/managed-by: Tiller
  annotations:
    "helm.sh/hook": test-success
spec:
  containers:
    - name: wget
      image: busybox
      command: ['wget']
      args:  ['sample-repeated-service:80']
  restartPolicy: Never
MANIFEST:

---
# Source: repeated-service/templates/service.yaml
apiVersion: v1
kind: Service
metadata:
  name: = dave-sample-repeated-service
  labels:
    app.kubernetes.io/name: repeated-service
    helm.sh/chart: repeated-service-0.1.0
    app.kubernetes.io/instance: sample
    app.kubernetes.io/version: "1.0"
    app.kubernetes.io/managed-by: Tiller
spec:
  type: ClusterIP
  ports:
    - port: 80
      targetPort: http
      protocol: TCP
      name: http
  selector:
    app.kubernetes.io/name: dave-repeated-service
    app.kubernetes.io/instance: dave-sample
---
# Source: repeated-service/templates/service.yaml
apiVersion: v1
kind: Service
metadata:
  name: = alice-sample-repeated-service
  labels:
    app.kubernetes.io/name: repeated-service
    helm.sh/chart: repeated-service-0.1.0
    app.kubernetes.io/instance: sample
    app.kubernetes.io/version: "1.0"
    app.kubernetes.io/managed-by: Tiller
spec:
  type: ClusterIP
  ports:
    - port: 80
      targetPort: http
      protocol: TCP
      name: http
  selector:
    app.kubernetes.io/name: alice-repeated-service
    app.kubernetes.io/instance: alice-sample
---
# Source: repeated-service/templates/service.yaml
apiVersion: v1
kind: Service
metadata:
  name: = bob-sample-repeated-service
  labels:
    app.kubernetes.io/name: repeated-service
    helm.sh/chart: repeated-service-0.1.0
    app.kubernetes.io/instance: sample
    app.kubernetes.io/version: "1.0"
    app.kubernetes.io/managed-by: Tiller
spec:
  type: ClusterIP
  ports:
    - port: 80
      targetPort: http
      protocol: TCP
      name: http
  selector:
    app.kubernetes.io/name: bob-repeated-service
    app.kubernetes.io/instance: bob-sample
---
# Source: repeated-service/templates/service.yaml
apiVersion: v1
kind: Service
metadata:
  name: = charlie-sample-repeated-service
  labels:
    app.kubernetes.io/name: repeated-service
    helm.sh/chart: repeated-service-0.1.0
    app.kubernetes.io/instance: sample
    app.kubernetes.io/version: "1.0"
    app.kubernetes.io/managed-by: Tiller
spec:
  type: ClusterIP
  ports:
    - port: 80
      targetPort: http
      protocol: TCP
      name: http
  selector:
    app.kubernetes.io/name: charlie-repeated-service
    app.kubernetes.io/instance: charlie-sample
---
# Source: repeated-service/templates/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: sample-repeated-service
  labels:
    app.kubernetes.io/name: repeated-service
    helm.sh/chart: repeated-service-0.1.0
    app.kubernetes.io/instance: sample
    app.kubernetes.io/version: "1.0"
    app.kubernetes.io/managed-by: Tiller
spec:
  replicas: 1
  selector:
    matchLabels:
      app.kubernetes.io/name: repeated-service
      app.kubernetes.io/instance: sample
  template:
    metadata:
      labels:
        app.kubernetes.io/name: repeated-service
        app.kubernetes.io/instance: sample
    spec:
      containers:
        - name: repeated-service
          image: "nginx:stable"
          imagePullPolicy: IfNotPresent
          ports:
            - name: http
              containerPort: 80
              protocol: TCP
          livenessProbe:
            httpGet:
              path: /
              port: http
          readinessProbe:
            httpGet:
              path: /
              port: http
          resources:
            {}
Stopping Tiller...
```
