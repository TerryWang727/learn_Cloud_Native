```
[
    {
        "Id": "645de3cb865ce53e893dca122ec2803bf418a591968497f3d87d6ca73cadcf6c",
        "Created": "2022-01-16T14:28:33.233154421Z",
        "Path": "./httpserver",
        "Args": [],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 37216,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2022-01-16T14:28:33.45531404Z",
            "FinishedAt": "0001-01-01T00:00:00Z"
        },
        "Image": "sha256:316a530d716c0fb11ea8586554e6529422bb4311691ae77c99c9f4c6917881a0",
        "ResolvConfPath": "/var/lib/docker/containers/645de3cb865ce53e893dca122ec2803bf418a591968497f3d87d6ca73cadcf6c/resolv.conf",
        "HostnamePath": "/var/lib/docker/containers/645de3cb865ce53e893dca122ec2803bf418a591968497f3d87d6ca73cadcf6c/hostname",
        "HostsPath": "/var/lib/docker/containers/645de3cb865ce53e893dca122ec2803bf418a591968497f3d87d6ca73cadcf6c/hosts",
        "LogPath": "/var/lib/docker/containers/645de3cb865ce53e893dca122ec2803bf418a591968497f3d87d6ca73cadcf6c/645de3cb865ce53e893dca122ec2803bf418a591968497f3d87d6ca73cadcf6c-json.log",
        "Name": "/sweet_dubinsky",
        "RestartCount": 0,
        "Driver": "overlay2",
        "Platform": "linux",
        "MountLabel": "",
        "ProcessLabel": "",
        "AppArmorProfile": "docker-default",
        "ExecIDs": null,
        "HostConfig": {
            "Binds": null,
            "ContainerIDFile": "",
            "LogConfig": {
                "Type": "json-file",
                "Config": {}
            },
            "NetworkMode": "default",
            "PortBindings": {},
            "RestartPolicy": {
                "Name": "no",
                "MaximumRetryCount": 0
            },
            "AutoRemove": false,
            "VolumeDriver": "",
            "VolumesFrom": null,
            "CapAdd": null,
            "CapDrop": null,
            "CgroupnsMode": "host",
            "Dns": [],
            "DnsOptions": [],
            "DnsSearch": [],
            "ExtraHosts": null,
            "GroupAdd": null,
            "IpcMode": "private",
            "Cgroup": "",
            "Links": null,
            "OomScoreAdj": 0,
            "PidMode": "",
            "Privileged": false,
            "PublishAllPorts": false,
            "ReadonlyRootfs": false,
            "SecurityOpt": null,
            "UTSMode": "",
            "UsernsMode": "",
            "ShmSize": 67108864,
            "Runtime": "runc",
            "ConsoleSize": [
                0,
                0
            ],
            "Isolation": "",
            "CpuShares": 0,
            "Memory": 0,
            "NanoCpus": 0,
            "CgroupParent": "",
            "BlkioWeight": 0,
            "BlkioWeightDevice": [],
            "BlkioDeviceReadBps": null,
            "BlkioDeviceWriteBps": null,
            "BlkioDeviceReadIOps": null,
            "BlkioDeviceWriteIOps": null,
            "CpuPeriod": 0,
            "CpuQuota": 0,
            "CpuRealtimePeriod": 0,
            "CpuRealtimeRuntime": 0,
            "CpusetCpus": "",
            "CpusetMems": "",
            "Devices": [],
            "DeviceCgroupRules": null,
            "DeviceRequests": null,
            "KernelMemory": 0,
            "KernelMemoryTCP": 0,
            "MemoryReservation": 0,
            "MemorySwap": 0,
            "MemorySwappiness": null,
            "OomKillDisable": false,
            "PidsLimit": null,
            "Ulimits": null,
            "CpuCount": 0,
            "CpuPercent": 0,
            "IOMaximumIOps": 0,
            "IOMaximumBandwidth": 0,
            "MaskedPaths": [
                "/proc/asound",
                "/proc/acpi",
                "/proc/kcore",
                "/proc/keys",
                "/proc/latency_stats",
                "/proc/timer_list",
                "/proc/timer_stats",
                "/proc/sched_debug",
                "/proc/scsi",
                "/sys/firmware"
            ],
            "ReadonlyPaths": [
                "/proc/bus",
                "/proc/fs",
                "/proc/irq",
                "/proc/sys",
                "/proc/sysrq-trigger"
            ]
        },
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/e1f6998a350c561733098a88fcaf024ec015fcf3d1e2073c69b5ec8227dcf364-init/diff:/var/lib/docker/overlay2/40c61b43a702120951a70d79a2d1a3c7830acdc1057b3477c06a1661489b0e42/diff:/var/lib/docker/overlay2/0156ed88692af59f6b50ca406b1cdf59543728d51c05220289b99b725ebf9c1f/diff",
                "MergedDir": "/var/lib/docker/overlay2/e1f6998a350c561733098a88fcaf024ec015fcf3d1e2073c69b5ec8227dcf364/merged",
                "UpperDir": "/var/lib/docker/overlay2/e1f6998a350c561733098a88fcaf024ec015fcf3d1e2073c69b5ec8227dcf364/diff",
                "WorkDir": "/var/lib/docker/overlay2/e1f6998a350c561733098a88fcaf024ec015fcf3d1e2073c69b5ec8227dcf364/work"
            },
            "Name": "overlay2"
        },
        "Mounts": [],
        "Config": {
            "Hostname": "645de3cb865c",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "ExposedPorts": {
                "2022/tcp": {}
            },
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
                "ENV=local"
            ],
            "Cmd": null,
            "Image": "httpserver:0.0.1",
            "Volumes": null,
            "WorkingDir": "/httpserver",
            "Entrypoint": [
                "./httpserver"
            ],
            "OnBuild": null,
            "Labels": {}
        },
        "NetworkSettings": {
            "Bridge": "",
            "SandboxID": "a593082b3258aec3781c7bf3ada7a1a999dfc9f35bf8189bc34e7a6ebbf8aca3",
            "HairpinMode": false,
            "LinkLocalIPv6Address": "",
            "LinkLocalIPv6PrefixLen": 0,
            "Ports": {
                "2022/tcp": null
            },
            "SandboxKey": "/var/run/docker/netns/a593082b3258",
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "EndpointID": "26116e146874a6f9c4bbe068b24dba50bdfeaffd40f7aa40ec678540107c217f",
            "Gateway": "172.17.0.1",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "IPAddress": "172.17.0.3",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
            "MacAddress": "02:42:ac:11:00:03",
            "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "NetworkID": "6209c61d40f483476c5282e814746a521a1e66e38ca823040a31b9ffd2181094",
                    "EndpointID": "26116e146874a6f9c4bbe068b24dba50bdfeaffd40f7aa40ec678540107c217f",
                    "Gateway": "172.17.0.1",
                    "IPAddress": "172.17.0.3",
                    "IPPrefixLen": 16,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "MacAddress": "02:42:ac:11:00:03",
                    "DriverOpts": null
                }
            }
        }
    }
]
```
