![Go version](https://img.shields.io/github/go-mod/go-version/appuio/control-api)
![Kubernetes version](https://img.shields.io/badge/k8s-v1.24-blue)
[![GitHub downloads](https://img.shields.io/github/downloads/vshn/appcat/total)](https://github.com/appuio/control-api/releases)

# Schedar Task

This repository is meant for VSHN candidates that want to join VSHN Schedar Team

## Requirements
* go (latest)
* kubectl 
* helm3
* make
* docker

## Context
The customer wants an easy way to deploy PostgreSQL instances using a single CustomResourceDefinition (CRD) in Kubernetes.
There are a couple of Kubernetes Operators that can help us achieve this goal such as [Stackgres Operator](https://stackgres.io/).
Once the Operator is installed in a cluster a PostgreSQL instance can be easily created with [SGCluster](https://stackgres.io/doc/1.1/reference/crd/sgcluster/#postgres) (CRD).
There are other CRDs such as [SGInstanceProfile](https://stackgres.io/doc/1.1/reference/crd/sginstanceprofile/) or [SGPostgresConfig](https://stackgres.io/doc/1.1/reference/crd/sgpgconfig/) 
which can further configure a PostgreSQL instance. Our goal is to abstract some of these CRDs into one single CRD (ex. API) of our own thus we need to create an Operator for that.
The bootstrap of the Operator is already done, moreover the bare-bones of Operator itself can be already deployed in a Kind cluster using the following commands:

* make install - installs a kind cluster and deploys the Stackgres Operator.
* make generate - generates CRDs.
* make deploy - builds and deploys the docker image in the cluster.

## The Task

1) Update the API (CRD) of this project so that the customer can issue a single resource to create a PostgreSQL instance. 
The API should expose at least the attributes found in [Custom Resource](postgres.yaml) under _spec_. 
2) Implement the [Reconciliation Loop](pkg/reconciler.go) so that the PostgreSQL instance is created (more details on [operator reconciliation](https://kubebyexample.com/learning-paths/operator-framework/operator-sdk-go/controller-reconcile-function)).
The update operation is allowed but should not perform any changes.
3) Resolve any issue that may arise during deployment.
4) If you successfully deployed and implemented all the above steps, try to implement the delete operation.

## Important links
* https://stackgres.io/doc/latest/reference/crd/
* https://kubebyexample.com/learning-paths/operator-framework/operator-sdk-go/controller-reconcile-function

## Notes
* After installation use the following command to connect to the cluster:
`kind get kubeconfig --name schedar-task > ~/.kube/config`
* Whenever a make deploy is issued, delete the current pod running in the cluster so that the latest changes apply.
`kubectl -n schedar-task delete pod <pod name>`
* Creating an SGCluster object requires:
  * Number of instances `spec.instances`
  * The storage capacity for the persistent volume `spec.pods.persistentvolume.size`
  * The version of the postgres `spec.postgres.version`
* Use the existing [CR](postgres.yaml) to try your operator.

