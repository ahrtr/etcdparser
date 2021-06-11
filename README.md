etcdparser
======
etcdparser is a tool to parse etcd's data, including snapshot files and WAL files. 

# Build
```
$ go build -ldflags "-X github.com/ahrtr/etcdparser/cmd.Version=0.1.0" -o ~/go/bin/ep main.go
```

# Usage
```
Etcd parser is used to parse etcd's data, including WAL and snapshot

Usage:
  ep [command]

Available Commands:
  help        Help about any command
  snap        Parse snap files
  wal         Parse wal files

Flags:
  -d, --data-dir string   Etcd data directory
  -h, --help              help for ep
  -r, --raw               Whether to print the data in raw format
  -s, --show-details      Whether to show the details: entries or snapshot data
  -v, --version           version for ep

Use "ep [command] --help" for more information about a command.
```
# Examples
## Example 1: parse the newest snapshot info
```
$ ep -d /tmp/etcd snap
```

The output is something like below,
```
Snapshot Metadata: 
{
    "conf_state": {
        "voters": [
            1725449293188291250
        ],
        "auto_leave": false
    },
    "index": 200002,
    "term": 2
}
```

## Example 2: show all the data of the newest snapshot
```
$ ep -d /tmp/etcd snap -s
```
The output is something like below,
```
Snapshot Metadata: 
{
    "conf_state": {
        "voters": [
            1725449293188291250
        ],
        "auto_leave": false
    },
    "index": 200002,
    "term": 2
}

-----------------------------------------------------

Snapshot Data: 
{
    "Root": {
        "Path": "/",
        "CreatedIndex": 0,
        "ModifiedIndex": 0,
        "ExpireTime": "0001-01-01T00:00:00Z",
        "Value": "",
        "Children": {
            "0": {
                "Path": "/0",
                "CreatedIndex": 0,
                "ModifiedIndex": 0,
                "ExpireTime": "0001-01-01T00:00:00Z",
    ......
```
Click **[examples/snapshot.log](examples/snapshot.log)** to get a complete example.

## Example 3: parse WAL file after the newest snapshot
```
$ ep -d /tmp/etcd wal
```
The output is something like below,
```
Snapshot Metadata: 
{
    "conf_state": {
        "voters": [
            1725449293188291250
        ],
        "auto_leave": false
    },
    "index": 200002,
    "term": 2
}

-----------------------------------------------------

Cluster Metadata: 
{
    "NodeID": 1725449293188291250,
    "ClusterID": 7895810959607866176
}

-----------------------------------------------------

HardState: 
{
    "term": 3,
    "vote": 1725449293188291250,
    "commit": 240576
}

-----------------------------------------------------

Entry: 
Entry number: 40574
```

## Example 4: show all the entries after the newest snapshot
```
$ ep -d /tmp/etcd wal -s
```
The output is something like below,
```
Snapshot Metadata:
{
    "conf_state": {
        "voters": [
            1725449293188291250
        ],
        "auto_leave": false
    },
    "index": 200002,
    "term": 2
}

-----------------------------------------------------

Metadata:
{
    "NodeID": 1725449293188291250,
    "ClusterID": 7895810959607866176
}

-----------------------------------------------------

HardState:
{
    "term": 3,
    "vote": 1725449293188291250,
    "commit": 240576
}

-----------------------------------------------------

Entry:
Entry number: 40574

{
    "Term": 2,
    "Index": 200003,
    "Type": 0,
    "Data": "CJCW+p6crZ7Z2gESBFNZTkMaACIAKAAyADgASABQAFgAYABoAHAAeM7x07viz8K/FoABAA=="
}
{
    "Term": 2,
    "Index": 200004,
    "Type": 0,
    "Data": "SgoIopX6npytntlaogYTCJGW+p6crZ7Z2gESBHJvb3QYBA=="
}
......
```
Click **[examples/wal.log](examples/wal.log)** to get a complete example.

# Contribute to this repo
Anyone is welcome to contribute to this repo. Please raise an issue firstly, then fork this repo and submit a pull request.

Currently this repo is under heavily development, any helps are appreciated!

# Support
If you need any support, please raise issues.

If you have any suggestions or proposals, please also raise issues. Thanks!