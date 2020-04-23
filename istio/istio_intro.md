# Istio Introduction

Display Istio Version
```
istio version
```

Verify prerequisites
```
istio verify-install
```

Demo Install
```
istoctl manifest apply --set profile=demo
```

Check Components
```
kubectl get pod -n istio-system
```

Verify Istio Installation
```
istioctl verify-install -f \
<(istioctl manifest generate --set profile=demo)
```
