# Overviews

In Kubernetes, when you specify a Pod, you can optionally specify how much of each resource a container needs. The most common resources to specify are CPU and memory (RAM); there are others - in our case, virtualized GPUs by xpuCUBE. When you specify the resource request for containers in a Pod, the kube-scheduler uses this information to decide which node to place the Pod on. When you specify a resource limit for a container, the kubelet enforces those limits so that the running container is not allowed to use more of that resource than the limit you set. The kubelet also reserves at least the request amount of that system resource specifically for that container to use. Extended resources are fully-qualified resource names outside the kubernetes.io domain. They allow cluster operators to advertise and users to consume the non-Kubernetes built-in resources. This is where xpuCUBE’s Extended Scheduler and Device Plugin come into play. No modification of kubelet is necessary. 

xpuCUBE Device Plugin is the Kubernetes custom device interface, it defines the action of heterogeneous resources(vGPUs) to report and allocate.

xpuCUBE Extended Scheduler helps Kubernetes scheduler allocate vGPU resources to Pods based on the configuration.
