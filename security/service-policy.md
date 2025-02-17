---
title: Use service rules in policy
description: Use Kubernetes Service names in policy rules.
---

### Big picture

Use {{site.prodname}} network policy to allow/deny traffic for Kubernetes services.

### Value

Using {{site.prodname}} network policy, you can leverage Kubernetes Service names to easily define access to Kubernetes services. Using service names in policy enables you to:

- Allow or deny access to the Kubernetes API service.
- Reference port information already declared by the application, making it easier to keep policy up-to-date as application requirements change.

### Limitations

Service rules in network policy are not implemented on Windows nodes. This will be included in a future release.

### Features

This how-to guide uses the following {{site.prodname}} features:

**NetworkPolicy** or **GlobalNetworkPolicy** with a service match criteria.

### How to

- [Allow access to the Kubernetes API for a specific namespace](#allow-access-to-the-kubernetes-api-for-a-specific-namespace)
- [Allow access to Kubernetes DNS for the entire cluster](#allow-access-to-kubernetes-dns-for-the-entire-cluster)

#### Allow access to the Kubernetes API for a specific namespace

In the following example, egress traffic is allowed to the `kubernetes` service in the `default` namespace for all pods in the namespace `my-app`. This service is the typical 
access point for the Kubernetes API server.

```yaml
apiVersion: projectcalico.org/v3
kind: NetworkPolicy
metadata:
  name: allow-api-access
  namespace: my-app
spec:
  selector: all()
  egress:
    - action: Allow
      destination:
        services:
          name: kubernetes
          namespace: default
```

Endpoint addresses and ports to allow will be automatically detected from the service.

#### Allow access to Kubernetes DNS for the entire cluster

In the following example, a GlobalNetworkPolicy is used to select all pods in the cluster to apply a rule which ensures 
all pods can access the Kubernetes DNS service.

```yaml
apiVersion: projectcalico.org/v3
kind: GlobalNetworkPolicy
metadata:
  name: allow-kube-dns
spec:
  selector: all()
  egress:
    - action: Allow
      destination:
        services:
          name: kube-dns
          namespace: kube-system
```

> **Note**: This policy also enacts a default-deny behavior for all pods, so make sure any other required application traffic is allowed by a policy.
{: .alert .alert-info}

### Above and beyond

- [Network policy]({{ site.baseurl }}/reference/resources/networkpolicy)
- [Global network policy]({{ site.baseurl }}/reference/resources/globalnetworkpolicy)
