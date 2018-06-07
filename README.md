# uuid-example in Golang

## UUID package Installation

Use the `go` command:

> $ go get github.com/satori/go.uuid
> $ go get github.com/meridiem00/uuid-example

## UUID function
	uuid.NewV1(node=None, clock_seq=None)
		Generate a UUID from a host ID, sequence number, and the current time. If node is not given, getnode() is used to obtain the hardware address. If clock_seq is given, it is used as the sequence number; otherwise a random 14-bit sequence number is chosen.
    uuid.NewV2(domain byte) 
        Generate DCE Security UUID based on POSIX UID/GID.
	uuid.NewV3(namespace, name) 
		Generate a UUID based on the MD5 hash of a namespace identifier (which is a UUID) and a name (which is a string).
	uuid.NewV4()
		Generate a random UUID.
	uuid.NewV5(namespace, name)
		Generate a UUID based on the SHA-1 hash of a namespace identifier (which is a UUID) and a name (which is a string).

## UUID namespace define

    The uuid module defines the following namespace identifiers for use with uuid3() or uuid5().
        uuid.NAMESPACE_DNS
            When this namespace is specified, the name string is a fully-qualified domain name.
        uuid.NAMESPACE_URL
            When this namespace is specified, the name string is a URL.
        uuid.NAMESPACE_OID
            When this namespace is specified, the name string is an ISO OID.
        uuid.NAMESPACE_X500
            When this namespace is specified, the name string is an X.500 DN in DER or a text output format.


<$GOPATH/src/github.com/satori/go.uuid/uuid.go>
``` go
    // Predefined namespace UUIDs.
    var (
        NamespaceDNS  = Must(FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
        NamespaceURL  = Must(FromString("6ba7b811-9dad-11d1-80b4-00c04fd430c8"))
        NamespaceOID  = Must(FromString("6ba7b812-9dad-11d1-80b4-00c04fd430c8"))
        NamespaceX500 = Must(FromString("6ba7b814-9dad-11d1-80b4-00c04fd430c8"))
    )
```

## UUID variant attribute
The uuid module defines the following constants for the possible values of the variant attribute:

    uuid.RESERVED_NCS
        Reserved for NCS compatibility.
    uuid.RFC_4122
        Specifies the UUID layout given in RFC 4122.
    uuid.RESERVED_MICROSOFT
        Reserved for Microsoft compatibility.
    uuid.RESERVED_FUTURE
        Reserved for future definition.

## Result of the output
Run below command in terminal

> $ go run uuid-example.go
```
UUIDv1:         a02906cd-6a9c-11e8-94be-4439c4961e59
UUIDv2:         000003e8-6a9c-21e8-9401-4439c4961e59
UUIDv3-DNS:     9a74c83e-2c09-3513-a74b-91d679be82b8
UUIDv3-URL:     79f1759f-cdf0-3f6e-9f86-2330cb620388
UUIDv3-OID:     d6945a7f-c888-3e32-ab11-6aa72409bc7e
UUIDv3-X500:    ecfa6054-241a-3369-89e5-7be0b84b793c
UUIDv4:         045624d3-0a10-4277-84bc-33cef319a1bb
UUIDv5-DNS:     64ee70a4-8cc1-5d25-abf2-dea6c79a09c8
UUIDv5-URL:     fedb2fa3-8f5c-5189-80e6-f563dd1cb8f9
UUIDv5-OID:     98de35fc-5e4d-5393-87c5-2c717d7c0d9a
UUIDv5-X500:    54bda7a2-a4a9-5b2c-8a5d-e1bf9737c2c2
```
## Verification with python uuid module
```
    $ python
    >>>
    >>> import uuid
    >>> uuid.uuid3(uuid.NAMESPACE_X500, "TMP112")
    UUID('85b5aa25-9de4-3a17-af74-2effc92e9a7a')
    >>> uuid.uuid3(uuid.NAMESPACE_DNS, "google.com")
    UUID('9a74c83e-2c09-3513-a74b-91d679be82b8')
    >>> uuid.uuid3(uuid.NAMESPACE_URL, "google.com")
    UUID('79f1759f-cdf0-3f6e-9f86-2330cb620388')
    >>> uuid.uuid3(uuid.NAMESPACE_X500, "Hello World!")
    UUID('ecfa6054-241a-3369-89e5-7be0b84b793c')
    >>> uuid.uuid5(uuid.NAMESPACE_DNS, "google.com")
    UUID('64ee70a4-8cc1-5d25-abf2-dea6c79a09c8')
    >>> uuid.uuid5(uuid.NAMESPACE_URL, "google.com")
    UUID('fedb2fa3-8f5c-5189-80e6-f563dd1cb8f9')
    >>> uuid.uuid5(uuid.NAMESPACE_X500, "Hello World!")
    UUID('54bda7a2-a4a9-5b2c-8a5d-e1bf9737c2c2')
```

## Copyright

Copyright (C) 2018 by Ryan Kim <email>.

UUID example code released under MIT License.
See [LICENSE](https://github.com/meridiem00/uuid-example/blob/master/LICENSE) for details.
