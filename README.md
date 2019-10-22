# kdump
Dumps all kubernetes api resources (pods, deployments, namespaces, etc..) to files in yaml form

Dumps everything from all configured contexts, and I mean everything.

* Calls `kubectl api-resources` to figure out what it has access to, then starts downloading all of it :).

WARNING: also dumps secrets. If you use this in for example a backup script and then commit to git (my use case), be sure to put secret in your .gitignore (or encrypt them)

### Example

`cd /tmp/kubedump`

`kdump`

`tree` -->


```.
├── dump.txt
├── prod
│   ├── apiservices.yml
│   ├── bgpconfigurations.yml
│   ├── cattle-system
│   │   ├── alertmanagers.yml
│   │   ├── clusterauthtokens.yml
│   │   ├── clusteruserattributes.yml
│   │   ├── configmaps.yml
│   │   ├── controllerrevisions.yml
│   │   ├── cronjobs.yml
│   │   ├── daemonsets.yml
│   │   ├── deployments.yml
│   │   ├── endpoints.yml
│   │   ├── events.yml
│   │   ├── horizontalpodautoscalers.yml
│   │   ├── ingresses.yml
│   │   ├── jobs.yml
│   │   ├── leases.yml
│   │   ├── limitranges.yml
│   │   ├── networkpolicies.yml
│   │   ├── persistentvolumeclaims.yml
│   │   ├── poddisruptionbudgets.yml
│   │   ├── pods.yml
│   │   ├── podtemplates.yml
│   │   ├── prometheuses.yml
│   │   ├── prometheusrules.yml
│   │   ├── replicasets.yml
│   │   ├── replicationcontrollers.yml
│   │   ├── resourcequotas.yml
│   │   ├── rolebindings.yml
│   │   ├── roles.yml
│   │   ├── secrets.yml
│   │   ├── serviceaccounts.yml
│   │   ├── servicemonitors.yml
│   │   ├── services.yml
│   │   └── statefulsets.yml
│   ├── certificatesigningrequests.yml
│   ├── clusterinformations.yml
│   ├── clusterrolebindings.yml
│   ├── clusterroles.yml
│   ├── componentstatuses.yml
│   ├── csidrivers.yml
│   ├── csinodes.yml
│   ├── customresourcedefinitions.yml
│   ├── default
│   │   ├── alertmanagers.yml
│   │   ├── clusterauthtokens.yml
│   │   ├── clusteruserattributes.yml
│   │   ├── configmaps.yml
│   │   ├── controllerrevisions.yml
│   │   ├── cronjobs.yml
│   │   ├── daemonsets.yml
│   │   ├── deployments.yml
│   │   ├── endpoints.yml
│   │   ├── events.yml
│   │   ├── horizontalpodautoscalers.yml
│   │   ├── ingresses.yml
│   │   ├── jobs.yml
│   │   ├── leases.yml
│   │   ├── limitranges.yml
│   │   ├── networkpolicies.yml
│   │   ├── persistentvolumeclaims.yml
│   │   ├── poddisruptionbudgets.yml
│   │   ├── pods.yml
│   │   ├── podtemplates.yml
│   │   ├── prometheuses.yml
│   │   ├── prometheusrules.yml
│   │   ├── replicasets.yml
│   │   ├── replicationcontrollers.yml
│   │   ├── resourcequotas.yml
│   │   ├── rolebindings.yml
│   │   ├── roles.yml
│   │   ├── secrets.yml
│   │   ├── serviceaccounts.yml
│   │   ├── servicemonitors.yml
│   │   ├── services.yml
│   │   └── statefulsets.yml
│   ├── felixconfigurations.yml
│   ├── globalnetworkpolicies.yml
│   ├── globalnetworksets.yml
│   ├── hostendpoints.yml
│   ├── ingress-nginx
│   │   ├── alertmanagers.yml
│   │   ├── clusterauthtokens.yml
│   │   ├── clusteruserattributes.yml
│   │   ├── configmaps.yml
│   │   ├── controllerrevisions.yml
│   │   ├── cronjobs.yml
│   │   ├── daemonsets.yml
│   │   ├── deployments.yml
│   │   ├── endpoints.yml
│   │   ├── events.yml
│   │   ├── horizontalpodautoscalers.yml
│   │   ├── ingresses.yml
│   │   ├── jobs.yml
│   │   ├── leases.yml
│   │   ├── limitranges.yml
│   │   ├── networkpolicies.yml
│   │   ├── persistentvolumeclaims.yml
│   │   ├── poddisruptionbudgets.yml
│   │   ├── pods.yml
│   │   ├── podtemplates.yml
│   │   ├── prometheuses.yml
│   │   ├── prometheusrules.yml
│   │   ├── replicasets.yml
│   │   ├── replicationcontrollers.yml
│   │   ├── resourcequotas.yml
│   │   ├── rolebindings.yml
│   │   ├── roles.yml
│   │   ├── secrets.yml
│   │   ├── serviceaccounts.yml
│   │   ├── servicemonitors.yml
│   │   ├── services.yml
│   │   └── statefulsets.yml
│   ├── ippools.yml
│   ├── kube-node-lease
│   │   ├── alertmanagers.yml
│   │   ├── clusterauthtokens.yml
│   │   ├── clusteruserattributes.yml
│   │   ├── configmaps.yml
│   │   ├── controllerrevisions.yml
│   │   ├── cronjobs.yml
│   │   ├── daemonsets.yml
│   │   ├── deployments.yml
│   │   ├── endpoints.yml
│   │   ├── events.yml
│   │   ├── horizontalpodautoscalers.yml
│   │   ├── ingresses.yml
│   │   ├── jobs.yml
│   │   ├── leases.yml
│   │   ├── limitranges.yml
│   │   ├── networkpolicies.yml
│   │   ├── persistentvolumeclaims.yml
│   │   ├── poddisruptionbudgets.yml
│   │   ├── pods.yml
│   │   ├── podtemplates.yml
│   │   ├── prometheuses.yml
│   │   ├── prometheusrules.yml
│   │   ├── replicasets.yml
│   │   ├── replicationcontrollers.yml
│   │   ├── resourcequotas.yml
│   │   ├── rolebindings.yml
│   │   ├── roles.yml
│   │   ├── secrets.yml
│   │   ├── serviceaccounts.yml
│   │   ├── servicemonitors.yml
│   │   ├── services.yml
│   │   └── statefulsets.yml
│   ├── kube-public
│   │   ├── alertmanagers.yml
│   │   ├── clusterauthtokens.yml
│   │   ├── clusteruserattributes.yml
│   │   ├── configmaps.yml
│   │   ├── controllerrevisions.yml
│   │   ├── cronjobs.yml
│   │   ├── daemonsets.yml
│   │   ├── deployments.yml
│   │   ├── endpoints.yml
│   │   ├── events.yml
│   │   ├── horizontalpodautoscalers.yml
│   │   ├── ingresses.yml
│   │   ├── jobs.yml
│   │   ├── leases.yml
│   │   ├── limitranges.yml
│   │   ├── networkpolicies.yml
│   │   ├── persistentvolumeclaims.yml
│   │   ├── poddisruptionbudgets.yml
│   │   ├── pods.yml
│   │   ├── podtemplates.yml
│   │   ├── prometheuses.yml
│   │   ├── prometheusrules.yml
│   │   ├── replicasets.yml
│   │   ├── replicationcontrollers.yml
│   │   ├── resourcequotas.yml
│   │   ├── rolebindings.yml
│   │   ├── roles.yml
│   │   ├── secrets.yml
│   │   ├── serviceaccounts.yml
│   │   ├── servicemonitors.yml
│   │   ├── services.yml
│   │   └── statefulsets.yml
│   ├── kube-system
│   │   ├── alertmanagers.yml
│   │   ├── clusterauthtokens.yml
│   │   ├── clusteruserattributes.yml
│   │   ├── configmaps.yml
│   │   ├── controllerrevisions.yml
│   │   ├── cronjobs.yml
│   │   ├── daemonsets.yml
│   │   ├── deployments.yml
│   │   ├── endpoints.yml
│   │   ├── events.yml
│   │   ├── horizontalpodautoscalers.yml
│   │   ├── ingresses.yml
│   │   ├── jobs.yml
│   │   ├── leases.yml
│   │   ├── limitranges.yml
│   │   ├── networkpolicies.yml
│   │   ├── persistentvolumeclaims.yml
│   │   ├── poddisruptionbudgets.yml
│   │   ├── pods.yml
│   │   ├── podtemplates.yml
│   │   ├── prometheuses.yml
│   │   ├── prometheusrules.yml
│   │   ├── replicasets.yml
│   │   ├── replicationcontrollers.yml
│   │   ├── resourcequotas.yml
│   │   ├── rolebindings.yml
│   │   ├── roles.yml
│   │   ├── secrets.yml
│   │   ├── serviceaccounts.yml
│   │   ├── servicemonitors.yml
│   │   ├── services.yml
│   │   └── statefulsets.yml
│   ├── mutatingwebhookconfigurations.yml
│   ├── namespaces.yml
│   ├── nodes.yml
│   ├── persistentvolumes.yml
│   ├── podsecuritypolicies.yml
│   ├── priorityclasses.yml
│   ├── runtimeclasses.yml
│   ├── storageclasses.yml
│   ├── validatingwebhookconfigurations.yml
│   └── volumeattachments.yml
└── test
    ├── apiservices.yml
    ├── bgpconfigurations.yml
    ├── cattle-system
    │   ├── alertmanagers.yml
    │   ├── clusterauthtokens.yml
    │   ├── clusteruserattributes.yml
    │   ├── configmaps.yml
    │   ├── controllerrevisions.yml
    │   ├── cronjobs.yml
    │   ├── daemonsets.yml
    │   ├── deployments.yml
    │   ├── endpoints.yml
    │   ├── events.yml
    │   ├── horizontalpodautoscalers.yml
    │   ├── ingresses.yml
    │   ├── jobs.yml
    │   ├── leases.yml
    │   ├── limitranges.yml
    │   ├── networkpolicies.yml
    │   ├── persistentvolumeclaims.yml
    │   ├── poddisruptionbudgets.yml
    │   ├── pods.yml
    │   ├── podtemplates.yml
    │   ├── prometheuses.yml
    │   ├── prometheusrules.yml
    │   ├── replicasets.yml
    │   ├── replicationcontrollers.yml
    │   ├── resourcequotas.yml
    │   ├── rolebindings.yml
    │   ├── roles.yml
    │   ├── secrets.yml
    │   ├── serviceaccounts.yml
    │   ├── servicemonitors.yml
    │   ├── services.yml
    │   └── statefulsets.yml
    ├── certificatesigningrequests.yml
    ├── clusterinformations.yml
    ├── clusterrolebindings.yml
    ├── clusterroles.yml
    ├── componentstatuses.yml
    ├── csidrivers.yml
    ├── csinodes.yml
    ├── customresourcedefinitions.yml
    ├── default
    │   ├── alertmanagers.yml
    │   ├── clusterauthtokens.yml
    │   ├── clusteruserattributes.yml
    │   ├── configmaps.yml
    │   ├── controllerrevisions.yml
    │   ├── cronjobs.yml
    │   ├── daemonsets.yml
    │   ├── deployments.yml
    │   ├── endpoints.yml
    │   ├── events.yml
    │   ├── horizontalpodautoscalers.yml
    │   ├── ingresses.yml
    │   ├── jobs.yml
    │   ├── leases.yml
    │   ├── limitranges.yml
    │   ├── networkpolicies.yml
    │   ├── persistentvolumeclaims.yml
    │   ├── poddisruptionbudgets.yml
    │   ├── pods.yml
    │   ├── podtemplates.yml
    │   ├── prometheuses.yml
    │   ├── prometheusrules.yml
    │   ├── replicasets.yml
    │   ├── replicationcontrollers.yml
    │   ├── resourcequotas.yml
    │   ├── rolebindings.yml
    │   ├── roles.yml
    │   ├── secrets.yml
    │   ├── serviceaccounts.yml
    │   ├── servicemonitors.yml
    │   ├── services.yml
    │   └── statefulsets.yml
    ├── felixconfigurations.yml
    ├── globalnetworkpolicies.yml
    ├── globalnetworksets.yml
    ├── hostendpoints.yml
    ├── ingress-nginx
    │   ├── alertmanagers.yml
    │   ├── clusterauthtokens.yml
    │   ├── clusteruserattributes.yml
    │   ├── configmaps.yml
    │   ├── controllerrevisions.yml
    │   ├── cronjobs.yml
    │   ├── daemonsets.yml
    │   ├── deployments.yml
    │   ├── endpoints.yml
    │   ├── events.yml
    │   ├── horizontalpodautoscalers.yml
    │   ├── ingresses.yml
    │   ├── jobs.yml
    │   ├── leases.yml
    │   ├── limitranges.yml
    │   ├── networkpolicies.yml
    │   ├── persistentvolumeclaims.yml
    │   ├── poddisruptionbudgets.yml
    │   ├── pods.yml
    │   ├── podtemplates.yml
    │   ├── prometheuses.yml
    │   ├── prometheusrules.yml
    │   ├── replicasets.yml
    │   ├── replicationcontrollers.yml
    │   ├── resourcequotas.yml
    │   ├── rolebindings.yml
    │   ├── roles.yml
    │   ├── secrets.yml
    │   ├── serviceaccounts.yml
    │   ├── servicemonitors.yml
    │   ├── services.yml
    │   └── statefulsets.yml
    ├── ippools.yml
    ├── kube-node-lease
    │   ├── alertmanagers.yml
    │   ├── clusterauthtokens.yml
    │   ├── clusteruserattributes.yml
    │   ├── configmaps.yml
    │   ├── controllerrevisions.yml
    │   ├── cronjobs.yml
    │   ├── daemonsets.yml
    │   ├── deployments.yml
    │   ├── endpoints.yml
    │   ├── events.yml
    │   ├── horizontalpodautoscalers.yml
    │   ├── ingresses.yml
    │   ├── jobs.yml
    │   ├── leases.yml
    │   ├── limitranges.yml
    │   ├── networkpolicies.yml
    │   ├── persistentvolumeclaims.yml
    │   ├── poddisruptionbudgets.yml
    │   ├── pods.yml
    │   ├── podtemplates.yml
    │   ├── prometheuses.yml
    │   ├── prometheusrules.yml
    │   ├── replicasets.yml
    │   ├── replicationcontrollers.yml
    │   ├── resourcequotas.yml
    │   ├── rolebindings.yml
    │   ├── roles.yml
    │   ├── secrets.yml
    │   ├── serviceaccounts.yml
    │   ├── servicemonitors.yml
    │   ├── services.yml
    │   └── statefulsets.yml
    ├── kube-public
    │   ├── alertmanagers.yml
    │   ├── clusterauthtokens.yml
    │   ├── clusteruserattributes.yml
    │   ├── configmaps.yml
    │   ├── controllerrevisions.yml
    │   ├── cronjobs.yml
    │   ├── daemonsets.yml
    │   ├── deployments.yml
    │   ├── endpoints.yml
    │   ├── events.yml
    │   ├── horizontalpodautoscalers.yml
    │   ├── ingresses.yml
    │   ├── jobs.yml
    │   ├── leases.yml
    │   ├── limitranges.yml
    │   ├── networkpolicies.yml
    │   ├── persistentvolumeclaims.yml
    │   ├── poddisruptionbudgets.yml
    │   ├── pods.yml
    │   ├── podtemplates.yml
    │   ├── prometheuses.yml
    │   ├── prometheusrules.yml
    │   ├── replicasets.yml
    │   ├── replicationcontrollers.yml
    │   ├── resourcequotas.yml
    │   ├── rolebindings.yml
    │   ├── roles.yml
    │   ├── secrets.yml
    │   ├── serviceaccounts.yml
    │   ├── servicemonitors.yml
    │   ├── services.yml
    │   └── statefulsets.yml
    ├── kube-system
    │   ├── alertmanagers.yml
    │   ├── clusterauthtokens.yml
    │   ├── clusteruserattributes.yml
    │   ├── configmaps.yml
    │   ├── controllerrevisions.yml
    │   ├── cronjobs.yml
    │   ├── daemonsets.yml
    │   ├── deployments.yml
    │   ├── endpoints.yml
    │   ├── events.yml
    │   ├── horizontalpodautoscalers.yml
    │   ├── ingresses.yml
    │   ├── jobs.yml
    │   ├── leases.yml
    │   ├── limitranges.yml
    │   ├── networkpolicies.yml
    │   ├── persistentvolumeclaims.yml
    │   ├── poddisruptionbudgets.yml
    │   ├── pods.yml
    │   ├── podtemplates.yml
    │   ├── prometheuses.yml
    │   ├── prometheusrules.yml
    │   ├── replicasets.yml
    │   ├── replicationcontrollers.yml
    │   ├── resourcequotas.yml
    │   ├── rolebindings.yml
    │   ├── roles.yml
    │   ├── secrets.yml
    │   ├── serviceaccounts.yml
    │   ├── servicemonitors.yml
    │   ├── services.yml
    │   └── statefulsets.yml
    ├── mutatingwebhookconfigurations.yml
    ├── namespaces.yml
    ├── nodes.yml
    ├── persistentvolumes.yml
    ├── podsecuritypolicies.yml
    ├── priorityclasses.yml
    ├── runtimeclasses.yml
    ├── storageclasses.yml
    ├── validatingwebhookconfigurations.yml
    └── volumeattachments.yml

14 directories, 435 files```
