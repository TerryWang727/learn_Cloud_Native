[
    {
        "Id": "sha256:87c084835d1b044cbf925155002435860344177e3960aa896021427473966a46",
        "RepoTags": [
            "httpserver:0.0.1"
        ],
        "RepoDigests": [],
        "Parent": "sha256:914fd8ed9061b88fcfd153ee312674fa678e35076b0fcb4fe432f911c1027301",
        "Comment": "",
        "Created": "2022-01-16T14:17:11.507153092Z",
        "Container": "0a91fc257062d0e2f4aaa92bf290abe522dd9a18bf4ba2e59572290f6b4d4d05",
        "ContainerConfig": {
            "Hostname": "0a91fc257062",
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
            "Cmd": [
                "/bin/sh",
                "-c",
                "#(nop) ",
                "ENTRYPOINT [\"./httpserver\"]"
            ],
            "Image": "sha256:914fd8ed9061b88fcfd153ee312674fa678e35076b0fcb4fe432f911c1027301",
            "Volumes": null,
            "WorkingDir": "/httpserver",
            "Entrypoint": [
                "./httpserver"
            ],
            "OnBuild": null,
            "Labels": {}
        },
        "DockerVersion": "20.10.7",
        "Author": "",
        "Config": {
            "Hostname": "",
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
            "Image": "sha256:914fd8ed9061b88fcfd153ee312674fa678e35076b0fcb4fe432f911c1027301",
            "Volumes": null,
            "WorkingDir": "/httpserver/",
            "Entrypoint": [
                "./httpserver"
            ],
            "OnBuild": null,
            "Labels": null
        },
        "Architecture": "amd64",
        "Os": "linux",
        "Size": 7333416,
        "VirtualSize": 7333416,
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/0156ed88692af59f6b50ca406b1cdf59543728d51c05220289b99b725ebf9c1f/diff",
                "MergedDir": "/var/lib/docker/overlay2/0bbd4cee6ec4349218fae03997a83b5c3b8e89ad0bfd1af1f98e33a1ceaa1600/merged",
                "UpperDir": "/var/lib/docker/overlay2/0bbd4cee6ec4349218fae03997a83b5c3b8e89ad0bfd1af1f98e33a1ceaa1600/diff",
                "WorkDir": "/var/lib/docker/overlay2/0bbd4cee6ec4349218fae03997a83b5c3b8e89ad0bfd1af1f98e33a1ceaa1600/work"
            },
            "Name": "overlay2"
        },
        "RootFS": {
            "Type": "layers",
            "Layers": [
                "sha256:01fd6df81c8ec7dd24bbbd72342671f41813f992999a3471b9d9cbc44ad88374",
                "sha256:3b65fe7915a539fb8a70d158c2fe26510fb7a233908e0582d0d4f5390fb4f4b8"
            ]
        },
        "Metadata": {
            "LastTagTime": "2022-01-16T22:17:11.520077739+08:00"
        }
    }
]
