# What is this ? 
This is a collection of good to have Makefiles, used in real-life.

<img style="display: inline-block;" src="{{ GITLAB_LINK }}/badges/master/pipeline.svg"/>
<img alt="coverage report" src="{{ GITLAB_LINK }}/badges/master/coverage.svg">



# What is it ?



# How to build ?



# Infrastructure
You will see a CloudFormation file `cloudformation.yml` this is for the application, the `builder.yml` is the file that
you run manually to setup 

To use the default file, as stated in Cloudformation.mk, you will be executing on Cloudformation.yml using following
commands:

- make cf-validate
- make cf-package
- make cf-deploy

If you for some reason will use the builder file, simple append -f Builder.mk in between like so:

- make -f Builder.mk cf-validate
- make -f Builder.mk cf package
- make -f Builder.mk cf-deploy

You can always chain the targets in a make command, like so:

- make [-f Filename.mk] cf-validate cf-package cf-deploy


# Deploy
Deployment will be done from the .gitlab-ci.yml, located in the <root> of your project.

Gitlab CI file will contain various stages [test, build, docker, deploy]

and each of those stages have a declarative template for what that stage should do, like so

```

```

# Run locally



