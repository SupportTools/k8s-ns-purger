# Project Brief: k8s-ns-purger

## Overview
k8s-ns-purger is a Kubernetes utility designed to automatically manage namespace lifecycle by purging namespaces based on age and label criteria. It serves as an automated maintenance tool to keep Kubernetes clusters clean and efficient.

## Core Requirements

### Functional Requirements
1. Automatically delete Kubernetes namespaces that:
   - Match specified label selectors
   - Exceed a configurable age threshold
2. Run continuously as a Kubernetes deployment
3. Support configurable purge intervals
4. Provide flexible configuration through:
   - Helm values
   - Environment variables

### Technical Requirements
1. Kubernetes Integration
   - Use in-cluster configuration
   - Implement proper RBAC permissions
   - Handle Kubernetes API interactions safely

2. Security Requirements
   - Run as non-root user
   - Use read-only root filesystem
   - Implement minimal RBAC permissions
   - Prevent privilege escalation

3. Resource Efficiency
   - Minimal resource footprint
   - Default resource limits:
     - CPU: 250m (limit), 5m (request)
     - Memory: 256Mi (limit), 32Mi (request)

## Project Goals
1. Maintain cluster cleanliness by automatically removing stale namespaces
2. Provide a reliable, secure, and resource-efficient solution
3. Offer easy deployment and configuration through Helm
4. Enable flexible namespace selection through label-based targeting
5. Ensure proper logging and monitoring capabilities

## Success Criteria
1. Successfully identifies and deletes namespaces matching criteria
2. Operates continuously without memory leaks or resource issues
3. Provides clear logging of all operations
4. Maintains security best practices
5. Can be easily deployed and configured in any Kubernetes cluster

## Constraints
1. Must operate within specified resource limits
2. Must maintain backward compatibility with existing label selectors
3. Must handle Kubernetes API errors gracefully
4. Must respect Kubernetes RBAC and security model

## Scope
### In Scope
- Namespace deletion based on age and labels
- Configuration through Helm and environment variables
- Logging and error handling
- Kubernetes RBAC setup
- Security hardening

### Out of Scope
- Namespace creation
- Backup/restore of deleted namespaces
- UI/Dashboard
- Direct cluster administration
- Resource management within namespaces
