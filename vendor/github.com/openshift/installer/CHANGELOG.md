# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## 0.3.0 - 2018-10-22

### Added

- Asset state is loaded from the install directory, allowing for a [staged
  install](docs/user/overview.md#multiple-invocations).
- A new `openshift-install destroy bootstrap` command destroys the
  bootstrap resources.  Ideally, this would be safe to run after the
  new `bootstrap-complete` event is pushed to the `kube-system`
  namespace, but there is currently a bug causing that event to be
  pushed too early.  For now, you're on your own figuring out when to
  call this command.

    For consistency, the old `destroy-cluster` has been deprecated in
    favor of `openshift-install destroy cluster`.

- The installer creates worker `MachineSet`s, instead of leaving that to
  [the machine-API operator][machine-api-operator].
- Creates master `Machine`s and tags masters to be picked up by the
  [AWS cluster-API provider][cluster-api-provider-aws].

### Changed

- The installer now respects the `AWS_PROFILE` environment variable
  when launching AWS clusters.
- Worker subnets are now created in the appropriate availability zone
  for AWS clusters.
- Use the released hyperkube and hypershift instead of hard-coded
  images.
- Lots of changes to keep up with the advancing release image, as
  OpenShift operators are added to control various cluster components.
- Lots of internal cleanup and minor fixes.

### Removed

- The Tectonic kube-core operator, which has been replaced by
  OpenShift operators.

## 0.2.0 - 2018-10-12

### Added

- Asset state is preserved between invocations, allowing for a staged
    install like:

    ```console
    $ openshift-install --dir=example install-config
    $ openshift-install --dir=example cluster
    ```

    which creates a cluster using the same data given in the
    install-config (including the same random cluster ID, etc.).
- [The kube-apiserver][kube-apiserver-operator] and
  [kube-controller-manager][kube-controller-manager-operator]
  operators are called to render additional cluster manifests.
- etcd is now available as a service in the `kube-system` namespace,
  and the new service is labeled so [Prometheus][] will scrape it.
- The `service-serving-cert-signer-signing-key` secret is now
  available in the `openshift-service-cert-signer` namespace, which
  gives [the service-serving cert signer][] the keys it needs to mint
  and manage certificates for Kubernetes services.
- The etcd-serving certificate is now passed through to [the
  kube-controller-manager operator][kube-controller-manager-operator].
- We disable some components which [the cluster-version
  operator][cluster-version-operator] would otherwise install but
  which conflict with the legacy tectonic-operators.
- The new `openshift-install graph` outputs the asset graph in [the
  DOT language][dot].
- `openshift-install version` now outputs the Terraform version as
  well as the installer version.

### Changed

- The [cluster-version operator][cluster-version-operator] is no
  longer run as a static pod.  Instead, we just wait until the control
  plane comes up and run it them.
- Terraform errors are logged to standard error even when
  `--log-level` is less than `debug`.
- Terraform is now invoked with `-no-color` and `-input=false`.
- The `cluster` target now includes both launching the cluster and
  populating `metadata.json`, regardless of whether the `terraform`
  invocation succeeds.  This allows `destroy-cluster` to cleanup
  cluster resources even when the `terraform` invocation fails.
- Reported errors now include more context, making them less
  enigmatic.
- Libvirt image caching is more efficient, caching unzipped images
  with a cache that grows by one unzipped image per RHCOS release in
  `$XDG_CACHE_HOME/openshift-install/libvirt/image`.  The previous
  implementation unzipped, when necessary, for every launched cluster,
  which was slow.  And the previous implementation added one unzipped
  image to `/tmp` per cluster launch, which consumed more disk space.
- Work continues on the OpenStack platform.
- Lots of internal cleanup, especially around asset generation.

### Removed

- The operatorstatus CRD.  Now [the cluster-version
  operator][cluster-version-operator] creates this on its own.
- The `machine-config-operator-images` config-map.  Now [the
  cluster-version operator][cluster-version-operator] pulls these from
  [the machine-config images][machine-config-operator].
- The `machine-api` app-version from the `tectonic-system` namespace.

## 0.1.0 - 2018-10-02

### Added

The `openshift-install` command.  This moves us to the new
install-config approach with [asset
generation](docs/design/assetgeneration.md) in Go instead of in
Terraform.  Terraform is still used to push the assets out to
resources on the backing platform (AWS, libvirt, or OpenStack), but
that push happens in a single Terraform invocation instead of in
multiple steps.  This makes installation faster, because more
resources can be created in parallel.  `openshift-install` also
dispenses with the distribution tarball; all required assets except
for a `terraform` binary are distributed in the `openshift-install`
binary.

The configuration and command-line interface are quite different, so
previous `tectonic` users are encouraged to start from scratch when
getting acquainted with `openshift-install`.  AWS users should look
[here](README.md#quick-start).  Libvirt users should look
[here](docs/dev/libvirt-howto.md).  The new `openshift-install` also
includes an interactive configuration generator, so you can launch the
installer and follow along as it guides you through the process.

### Removed

The `tectonic` command and tarball distribution are gone.  Please use
the new `openshift-install` command instead.

[cluster-api-provider-aws]: https://github.com/openshift/cluster-api-provider-aws
[cluster-version-operator]: https://github.com/openshift/cluster-version-operator
[dot]: https://www.graphviz.org/doc/info/lang.html
[kube-apiserver-operator]: https://github.com/openshift/cluster-kube-apiserver-operator
[kube-controller-manager-operator]: https://github.com/openshift/cluster-kube-controller-manager-operator
[machine-api-operator]: https://github.com/openshift/machine-api-operator
[machine-config-operator]: https://github.com/openshift/machine-config-operator
[Prometheus]: https://github.com/prometheus/prometheus
[service-serving-cert-signer]: https://github.com/openshift/service-serving-cert-signer
