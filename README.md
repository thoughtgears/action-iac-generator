# IAC generator action

This an example of how you can use a templating engine to generate terraform in a GHA allowing you to use an opinionated
YAML definition to generate terraform code. This is a very simple example but it can be extended to do more complex, you 
eiter extend with more templates, or build modules that is used in the templates. This is not production ready code but
a way to show how you can use a templating engine to generate terraform code and how to create a simple interface for 
developers to define resources in a YAML file rather than writing terraform code.

You should then be able to use this action in your workflow to generate terraform code and then chain this with tools 
like [conftest](https://www.conftest.dev/) and [digger](https://www.digger.dev/) to validate the generated terraform code
and deploy it to your cloud provider.

## Usage


### Infrastructure Yaml file

```yaml
name: service-x
project_id: project-y-1234
region: europe-west2
modules:
  - type: pubsub
    resource_name: service-x-pubsub
  - type: pubsub
    resource_name: service-x-pubsub-2
  - type: cloud-run
    resource_name: service-x-cloud-run
    cloud_run:
      image: us-docker.pkg.dev/cloudrun/container/hello
      limits:
        cpu: 1
        memory: 512Mi
```