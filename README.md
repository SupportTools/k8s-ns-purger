# k8s-ns-purger

A Kubernetes utility that automatically purges namespaces based on age and labels. This tool helps maintain cluster cleanliness by removing old or temporary namespaces that are no longer needed.

## Features

- Automatically deletes namespaces based on age
- Label-based namespace selection
- Configurable purge intervals and age thresholds
- Runs as a Kubernetes deployment
- Helm chart for easy installation
- Secure by default with non-root execution

## Installation

### Using Helm

```bash
helm repo add supporttools https://charts.supporttools.dev
helm repo update
helm install k8s-ns-purger supporttools/k8s-ns-purger
```

## Configuration

The following configuration options are available through Helm values:

| Parameter | Description | Default |
|-----------|-------------|---------|
| `settings.debug` | Enable debug logging | `false` |
| `settings.purgeInterval` | Time in seconds between purge runs | `3600` (1 hour) |
| `settings.purgeThreshold` | Time in seconds before a namespace is considered stale | `86400` (24 hours) |
| `settings.purgeLabel` | Label selector for namespaces to purge | `k8s-ns-purger=true` |
| `replicaCount` | Number of replicas to run | `1` |
| `image.repository` | Container image repository | `supporttools/k8s-ns-purger` |
| `image.tag` | Container image tag | `latest` |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` |

## Usage

1. Label the namespaces you want to be automatically purged:

```bash
kubectl label namespace my-temp-namespace k8s-ns-purger=true
```

2. The namespace will be automatically deleted after it reaches the age threshold (default 24 hours).

### Environment Variables

When running outside of Helm, the following environment variables can be used:

| Variable | Description | Default |
|----------|-------------|---------|
| `DEBUG` | Enable debug logging | `false` |
| `LABEL_SELECTOR` | Label selector for namespaces | `k8s-ns-purger=true` |
| `PURGE_AGE` | Age threshold for namespace deletion | `8h` |
| `INTERVAL` | Time between purge checks | `1h` |

## Security

k8s-ns-purger is designed with security in mind:

- Runs as non-root user
- Uses read-only root filesystem
- Requires minimal RBAC permissions
- No privilege escalation allowed

## Resource Requirements

Default resource limits:

```yaml
resources:
  limits:
    cpu: 250m
    memory: 256Mi
  requests:
    cpu: 5m
    memory: 32Mi
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the terms of the LICENSE file included in the repository.
