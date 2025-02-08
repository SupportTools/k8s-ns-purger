# Technical Context: k8s-ns-purger

## Technology Stack

### Core Technologies
1. **Go**
   - Version: 1.23.0
   - Toolchain: go1.23.6
   - Purpose: Main implementation language

2. **Kubernetes**
   - Client Version: v0.32.1
   - API Version: v0.32.1
   - Purpose: Target platform and API integration

### Key Dependencies

1. **Kubernetes Libraries**
   ```go
   k8s.io/apimachinery v0.32.1
   k8s.io/client-go v0.32.1
   ```
   - Core Kubernetes API interaction
   - Client configuration
   - Resource management

2. **Logging**
   ```go
   github.com/sirupsen/logrus v1.9.3
   ```
   - Structured logging
   - Log level management
   - Error reporting

3. **Testing**
   ```go
   github.com/stretchr/testify v1.10.0
   ```
   - Unit testing framework
   - Assertions
   - Mocking support

## Development Setup

### Requirements
1. **Go Development Environment**
   - Go 1.23.0 or higher
   - Go modules enabled
   - GOPATH configured

2. **Kubernetes Development**
   - Access to Kubernetes cluster
   - kubectl configured
   - Helm (for deployment)

3. **Container Tools**
   - Docker (for building images)
   - Container registry access

### Build Process
1. **Go Build**
   ```bash
   go build -o k8s-ns-purger
   ```

2. **Container Build**
   ```bash
   docker build -t supporttools/k8s-ns-purger .
   ```

3. **Helm Package**
   ```bash
   helm package charts/k8s-ns-purger
   ```

## Configuration

### Environment Variables
| Variable | Description | Default |
|----------|-------------|---------|
| `DEBUG` | Enable debug logging | `false` |
| `LABEL_SELECTOR` | Label selector | `k8s-ns-purger=true` |
| `PURGE_AGE` | Age threshold | `8h` |
| `INTERVAL` | Check interval | `1h` |

### Helm Values
| Parameter | Description | Default |
|-----------|-------------|---------|
| `settings.debug` | Debug logging | `false` |
| `settings.purgeInterval` | Purge interval | `3600` |
| `settings.purgeThreshold` | Age threshold | `86400` |
| `settings.purgeLabel` | Label selector | `k8s-ns-purger=true` |

## Technical Constraints

### Resource Limits
1. **Container Resources**
   ```yaml
   resources:
     limits:
       cpu: 250m
       memory: 256Mi
     requests:
       cpu: 5m
       memory: 32Mi
   ```

2. **Operation Limits**
   - Single instance recommended
   - In-cluster deployment only
   - Namespace-level operations

### Security Constraints
1. **Container Security**
   - Non-root user
   - Read-only filesystem
   - No privilege escalation

2. **RBAC Requirements**
   - Namespace list/delete permissions
   - ClusterRole binding
   - Service account

## Deployment Architecture

### Kubernetes Resources
1. **Deployment**
   - Single replica
   - Resource limits
   - Security context
   - Liveness probe

2. **Service Account**
   - RBAC integration
   - Minimal permissions
   - Cluster-wide access

3. **RBAC**
   - ClusterRole
   - ClusterRoleBinding
   - Namespace permissions

### Container Configuration
1. **Image**
   - Base: scratch/minimal
   - Non-root user
   - Read-only filesystem

2. **Security Context**
   ```yaml
   securityContext:
     runAsNonRoot: true
     readOnlyRootFilesystem: true
     allowPrivilegeEscalation: false
   ```

## Development Workflow

### Local Development
1. **Setup**
   ```bash
   go mod download
   go mod verify
   ```

2. **Testing**
   ```bash
   go test ./...
   ```

3. **Building**
   ```bash
   go build
   docker build .
   ```

### Deployment Process
1. **Package**
   ```bash
   helm package charts/k8s-ns-purger
   ```

2. **Deploy**
   ```bash
   helm install k8s-ns-purger supporttools/k8s-ns-purger
   ```

3. **Verify**
   ```bash
   kubectl get pods -l app=k8s-ns-purger
   ```

## Monitoring & Debugging

### Logging
- Structured JSON logging
- Configurable log levels
- Operation tracking
- Error details

### Metrics
- Container resource usage
- Operation counts
- Error rates
- Deletion statistics

### Debugging
- Debug log level
- Container logs
- Event tracking
- Error tracing
