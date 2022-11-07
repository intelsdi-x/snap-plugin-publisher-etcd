DISCONTINUATION OF PROJECT. 

This project will no longer be maintained by Intel.

This project has been identified as having known security escapes.

Intel has ceased development and contributions including, but not limited to, maintenance, bug fixes, new releases, or updates, to this project.  

Intel no longer accepts patches to this project.

# DISCONTINUATION OF PROJECT 

**This project will no longer be maintained by Intel.  Intel will not provide or guarantee development of or support for this project, including but not limited to, maintenance, bug fixes, new releases or updates.  Patches to this project are no longer accepted by Intel. If you have an ongoing need to use this project, are interested in independently developing it, or would like to maintain patches for the community, please create your own fork of the project.**


[![Build Status](https://travis-ci.org/intelsdi-x/snap-plugin-publisher-etcd.svg?branch=master)](https://travis-ci.org/intelsdi-x/snap-plugin-publisher-etcd)


# snap publisher plugin - etcd

Allows publishing of data to [etcd](https://coreos.com/etcd/) from the [Snap framework](http://github.com:intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](configuration-and-usage)
2. [Documentation](#documentation)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license-and-authors)
6. [Acknowledgements](#acknowledgements)

## Getting Started
### System Requirements
* A running instance of etcd ([release page](https://github.com/coreos/etcd/releases/))
* [golang 1.5+](https://golang.org/dl/) (needed only for building)

### Operating systems
All OSs currently supported by snap:
* Linux/amd64
* Darwin/amd64

### Installation
#### Download etcd plugin binary:
You can get the pre-built binaries for your OS and architecture at snap's [GitHub Releases](https://github.com/intelsdi-x/snap/releases) page.

#### To build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-publisher-etcd
Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-publisher-etcd.git
```

Build the plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `/build/rootfs/`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Set Snap Global Global Config field `etcd_host` to proper etcd host string (e.g. `http://localhost:2379`)

## Documentation

### Examples
Example of using snap-plugin-publisher-etcd to store metrics collected by snap-plugin-colector-psutil.

Run etcd (for more info on how to run etcd see [etcd getting started](https://coreos.com/etcd/docs/latest/getting-started-with-etcd.html) and [etcd GitHub](https://github.com/coreos/etcd)):
```
$ etcd &
```

Ensure [Snap daemon is running](https://github.com/intelsdi-x/snap#running-snap):
```
$ snapteld -l -t 0 &
```
Download and load Snap plugins:
```
$ snaptel plugin load snap-plugin-collector-psutil
$ snaptel plugin load snap-plugin-publisher-etcd
```
Create a task:
```
$ snaptel task create -t psutil-etcd.yml
Using task manifest to create task
Task created
ID: 3b2bc29c-8256-4ca0-a958-05f5d41f6f11
Name: Task-3b2bc29c-8256-4ca0-a958-05f5d41f6f11
State: Running
```
You may view example tasks [here](https://github.com/intelsdi-x/snap-plugin-publisher-etcd/blob/master/examples/tasks/).

Ensure the task is running and collecting metrics:
```
$ snaptel task watch 3b2bc29c-8256-4ca0-a958-05f5d41f6f11
```
Query etcd to see metrics:
```
$ etcdctl ls /intel/psutil --recursive
```


### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-publisher-etcd/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-publisher-etcd/pulls).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[snap](http://github.com:intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: [@kindermoumoute](https://github.com/kindermoumoute/)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.
