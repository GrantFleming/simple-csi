# Architecture

As per the CSI spec, there are three main services distrinbuted across two
plugins. The services are:

- Identity Service
- Controller Service
- Node Service

The two plugins are:

- Controller Plugin
- Node Plugin

The Identity Service must be implemeneted by both plugins.

For this project, we opt for the simplest archiecture where a single binary is
provided that runs on all nodes and implements all the services.

```Figure
+-------------------------------------------+
|                                           |
|  +------------+           +------------+  |
|  |     CO     |   gRPC    | Controller |  |
|  |            +----------->    Node    |  |
|  +------------+           |   Plugin   |  |
|                           +------------+  |
|                                           |
+-------------------------------------------+

Figure 3: Headless Plugin deployment, only the CO Node hosts run
Plugins. A unified Plugin component supplies both the Controller
Service and Node Service.
```

Since we want to support dynamic creation of volumes we opt for the following
interaction pattern:

```Figure
   CreateVolume +------------+ DeleteVolume
 +------------->|  CREATED   +--------------+
 |              +---+----^---+              |
 |       Controller |    | Controller       v
+++         Publish |    | Unpublish       +++
|X|          Volume |    | Volume          | |
+-+             +---v----+---+             +-+
                | NODE_READY |
                +---+----^---+
               Node |    | Node
            Publish |    | Unpublish
             Volume |    | Volume
                +---v----+---+
                | PUBLISHED  |
                +------------+

Figure 5: The lifecycle of a dynamically provisioned volume, from
creation to destruction.
```

Given that we are essentially creating an emptydir in a specific directory, it
is expected that `CreateVolume` and `ControllerPublishVolume` will do very
little, perhaps just logging the that they have been called.
`NodePublishVolume` will probably create the volume as required.
