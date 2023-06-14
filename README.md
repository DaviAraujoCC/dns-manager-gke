# About

A project created to manage dns entries in GKE.

## Requirements

* Go 1.17+
* Make

## Authentication

This application supports ~/kube/config authentication or via service account.

To authenticate with GCP, you need to create a service account that have access to <b>dns.resourceRecordSets</b> and <b>dns.change.*</b>.

## Variables:


| Variable | Description |
| --- | --- |
| GOOGLE_APPLICATION_CREDENTIALS | Path to service_account.json used to authenticate to the GCP services|
| PROJECT_ID | GCP Project ID |
| DNS_SUFFIX | Your managedZone suffix, Example: (kube.stg.gce.domain.com.br.) |
| MANAGED_ZONE | Managed zone name in GCP |
| NAMESPACE | Namespace where the script will listen to |

## Usage

### Local

Built executable releases can be found on the homepage and on the github release page, https://github.com/DaviAraujoCC/dns-manager-gke/releases

### Cronjob

To create the image and push it to docker hub:

```
$ make docker-build docker-push IMG={Image name}
```


Create a secret with the service account file:

```
$ kubectl create secret generic dns-manager-gke-secret --from-file=service-account.json=/path/to/service_account.json
```

Edit the manifest inside manifests folder and after that run:

```
$ kubectl apply -k manifests/
```

To undeploy:

```
$ kubectl delete -k manifests/
```


## Step-by-step

 1. The application will get all available services with LoadBalancer feature enabled from the k8s cluster in the specified namespace.
 2. After that it will list for dns entries in the specified managed zone via gcloud API.
 3. It will compare the entries with the current services created in your cluster, if is no dns present and the service is running, it will create dns for you using the template (service_name.DNS_SUFFIX).
 4. if it's present but the IP differs from the current one used in cluster, it will update the respective dns entry.
 5. If the service is not running, it will delete the respective dns entry, with all unused entries.


## TODO

- [ ] Create an kubernetes operator
- [ ] Code cleanup