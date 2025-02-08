# Product Context: k8s-ns-purger

## Problem Space

### Core Problems
1. **Namespace Sprawl**
   - Kubernetes clusters accumulate unused namespaces over time
   - Manual cleanup is time-consuming and error-prone
   - Stale namespaces waste cluster resources
   - Cluttered namespace list makes cluster management difficult

2. **Resource Management**
   - Orphaned resources in forgotten namespaces
   - Unnecessary resource consumption
   - Potential cost implications in cloud environments

3. **Operational Overhead**
   - Manual tracking of namespace lifetimes
   - Manual cleanup processes
   - Risk of accidentally deleting active namespaces

## Solution

### How It Works
1. **Automated Monitoring**
   - Continuously monitors namespaces in the cluster
   - Tracks namespace age from creation timestamp
   - Identifies namespaces with specific labels

2. **Intelligent Purging**
   - Evaluates namespaces against configured criteria
   - Deletes namespaces that:
     * Match the specified label selector (default: k8s-ns-purger=true)
     * Exceed the age threshold (default: 24 hours)
   - Logs all deletion operations for audit purposes

3. **Configurable Behavior**
   - Adjustable purge intervals
   - Customizable age thresholds
   - Flexible label selection
   - Environment variable overrides

### User Experience Goals

1. **Simplicity**
   - Easy installation through Helm
   - Minimal configuration required
   - Clear documentation
   - Intuitive label-based selection

2. **Reliability**
   - Consistent operation
   - Predictable behavior
   - Graceful error handling
   - No unexpected deletions

3. **Transparency**
   - Clear logging of all actions
   - Traceable deletion decisions
   - Visible operational status

4. **Control**
   - Opt-in through labels
   - Configurable thresholds
   - Adjustable intervals
   - Override capabilities

## Use Cases

### Primary Use Cases
1. **Development Environments**
   - Cleanup of temporary testing namespaces
   - Management of development sandboxes
   - CI/CD pipeline cleanup

2. **Cluster Maintenance**
   - Automated housekeeping
   - Resource optimization
   - Namespace lifecycle management

3. **Multi-tenant Clusters**
   - Temporary workspace cleanup
   - Resource quota management
   - Tenant namespace lifecycle

### User Personas

1. **Cluster Administrators**
   - Need: Automated cluster maintenance
   - Goal: Keep cluster clean and efficient
   - Value: Reduced operational overhead

2. **DevOps Engineers**
   - Need: Automated cleanup for CI/CD
   - Goal: Maintain clean test environments
   - Value: Automated lifecycle management

3. **Development Teams**
   - Need: Self-cleaning development spaces
   - Goal: Focus on development, not cleanup
   - Value: Automated workspace management

## Success Metrics

1. **Operational**
   - Successful namespace deletions
   - Error-free operation
   - Resource usage within limits

2. **User Experience**
   - Minimal configuration needs
   - Clear operational feedback
   - Predictable behavior

3. **System Health**
   - Stable memory usage
   - Consistent performance
   - Reliable operation
