# skydive-visualizer

A visualization tool for [skydive](https://github.com/skydive-project/skydive) data

## How-to run

```
skydive-visualizer -config config.yml
```

## Example configuration

```
---
skydive:
  url: http://skydive-analyzer.preprod.crto.in:8082/ # skydive analyzer URL

server:
  listen: ':8000' # listen address

# optional: mysql IPAM
ipam:
  host: my-db.domain
  port: 3306
  user: user
  password: pass
  database: phpipam
  mapping:
  - match: "^.+:PROD$" # match subnet description
    value: "$1" # value to display
    id: 200 # dimension id
    name: Datacenter # dimension name

# optional: Chef
chef:
  user: chef-user
  key: chef-user-private-key
  servers:
  - https://chef.domain

  attrs_mapping:
  - key: my_attribute
    id: 450 # dimension id
    name: TLA # dimension name

```
