# k8s-practices

### links

* [apiserver-builder-alpha](https://github.com/kubernetes-incubator/apiserver-builder-alpha/blob/master/README.md)
* [k8s api](https://kubernetes.io/docs/concepts/overview/kubernetes-api/#v1beta1-v1beta2-and-v1beta3-are-deprecated-please-move-to-v1-asap)
* [k8s gc](https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection)
* [k8s extensible-admission-controllers](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/)
* [bazel](https://bazel.build/)
* [bazel for go](https://github.com/bazelbuild/bazel-gazelle)

### steps

* [api structure](https://github.com/kubernetes-incubator/apiserver-builder-alpha/blob/master/docs/concepts/api_building_overview.md#api-structure)
* [build apiserver from scratch](https://github.com/kubernetes-incubator/apiserver-builder-alpha/blob/master/docs/tools_user_guide.md)
* [run in minikube](https://github.com/kubernetes-incubator/apiserver-builder-alpha/blob/master/docs/running_in_minikube.md)
* [example](https://github.com/kubernetes-incubator/apiserver-builder-alpha/commits/example-simple)
* [library user guide](https://github.com/kubernetes-incubator/apiserver-builder-alpha/blob/master/docs/libraries_user_guide.md)

```
> cd $GOPATH/src/k8s-practices
> touch boilerplate.go.txt
> apiserver-boot init repo --domain jw.io

> apiserver-boot create group version resource --group insect --version v1beta1 --kind Bee

> apiserver-boot build generated # run the code generators and build the executables
> apiserver-boot build executables
> apiserver-boot build docs # generate doc

> apiserver-boot run local # run locally
> kubectl --kubeconfig kubeconfig api-versions
> kubectl --kubeconfig kubeconfig create -f sample/bee.yaml
> go test ./pkg/... # run test case

> # run in minikube
> kubectl create ns insect
> apiserver-boot build config --local-minikube --name insect --namespace insect
> kubectl create -f config/apiserver.yaml
> apiserver-boot run local-minikube # aggregated with the minikube cluster

> apiserver-boot build container # cross-compile the go binaries into a container image
> apiserver-boot build config # emit yaml configuration for running the apiserver, controller-manager and etcd in a cluster
```

![store_reconcile](https://github.com/kubernetes-incubator/apiserver-builder-alpha/raw/master/docs/concepts/store_reconcile.jpg)
![storage](https://github.com/kubernetes-incubator/apiserver-builder-alpha/blob/master/docs/concepts/storage.jpg)
![reconciliation](https://github.com/kubernetes-incubator/apiserver-builder-alpha/raw/master/docs/concepts/reconciliation.jpg)
![End to end Deployment example](https://github.com/kubernetes-incubator/apiserver-builder-alpha/raw/master/docs/concepts/store_reconcile_example.jpg)
![Deployment](https://github.com/kubernetes-incubator/apiserver-builder-alpha/raw/master/docs/concepts/extensionserver.jpg)

