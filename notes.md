# TODOs

- Reconsider what should to be exported
    - Customers may need access to internals to handle e.g. usertypes or
      to work around issues with older versions
- Maybe split into multiple packages:
    tds/
        bind subpackages together to act as an interface for database/sql
    tds/transport/
        connection/interface handling and abstraction?
        conn/message/package/packet could reside here
        security mechanisms could also be handled here
    tds/datatypes/
        data type abstraction?
        only when generating code or producing a large number of structs
        with lots of code
- DONE implement loginack
- DONE implement msg
    - comb through the notes on msg; there's a lot of auxiliary
      information on tds behaviour
- REWORK implement paramfmt
- REWORK implement params
- DONE implement done
- DONE implement datatypes
- implement `TDS_CAPABILITY`
    - request `TDS_DATA_COLUMNSTATUS` capability
- update ReadFrom methods - thanks to the changes to the stream
    generating/acceptance these are run in a separate blocking
    goroutine now they don't need to provide .Error and .Finish anymore
    and instead can return an error akin to WriteTo

# Memory handling

## Image

Images can have up to 2 147 483 647 bytes (2,147 TB) of data

## Flow

conn.Receive
    Message.readFrom
        - start goroutine reading from conn and writing packets into
          channel
        loop:
            inspeckt packet for token
            -> new token: add old package if exists to message
                create new package based on token
                read packet into package
            -> no token: read packet into package
            exit if packet.Header.Status has TDS_BUFSTAT_EOM

## handling token & tokenless

Tokenless messages can only occur if the message has no packages and can
simply read into a packageTokenless until EOM.
May not be needed server->client

# endianness

- Headers are always in big endian
- ASE sends big endian by default
- client can request little endian
    - big endian will be used until the login package has been sent
      successfully (incl. cipher communication)
 -> check if client-library can be set to use little endian
    that would allow to use big endian everywhere instead of requiring
    a flag whose pointer must be passed around
    => cannot be changed - little endian is fixed.
- when exactly is the byte order being changed?
    -> Maybe fix byte order for cgo/go to little/big endian respectively
        and pass byte order along to library functions
        Less work, would be easy to test

# security

TODO const security bits

## protocols

1. encrypted password
    only encrypts password
    send encryption key using TDS_VARBINARY, then encrypted password
2. challenge response

3. encrypted password & trusted user
    encrypts password and security labels
    TODO: What are security labels?
4. trusted user
    security labels
5. extended encrypted password
    only encrypts password
    send cipher suite ID (TDS_INT4) and encryption key (TDS_VARBINARY),
        then proceed
6. extended plus encrypted password
    only encrypts password
    send cipher suite ID (TDS_INT4), encryption key (TDS_VARBINARY) and
        32 byte nonce (TDS_LONGBINARY), then proceed
7. on demand encryption

-> implement {extended,}{,plus} encrypted password

## encrypted password

```
-> LOGIN TDS_SEC_LOG_ENCRYPT
        password = ""
        lpwdlen = 0
        rempwd = ""
        lrempwdlen = 0

<- loginack TDS_LOG_NEGOTIATE
    MSG TDS_MSG_SEC_ENCRYPT
    PARAMFMT TDS_VARBINARY
    PARAMS password key
    DONE FINAL

-> MSG TDS_MSG_SEC_LOGPWD
    PARAMFMT TDS_VARBINARY
    PARAMS encrypted password
    MSG TDS_MSG_SEC_REMPWD
    PARAMFMT TDS_VARCHAR, TDS_VARBINARY, ... (repeat)
    PARAMS remote server name, remote server password, ... (repeat)

<- loginack SUCCEED
    DONE FINAL
```

## extended encrypted password

```
-> LOGIN TDS_SEC_LOG_ENCRYPT2
        password = ""
        lpwdlen = 0
        rempwd = ""
        lrempwdlen = 0

<- loginack TDS_LOG_NEGOTIATE
    MSG TDS_MSG_SEC_ENCRYPT2
    PARAMFMT TDS_INT4, TDS_VARBINARY
    PARAMS password cipher suite, key
    DONE FINAL

-> MSG TDS_MSG_SEC_LOGPWD2
    PARAMFMT TDS_VARBINARY
    PARAMS encrypted password
    MSG TDS_MSG_SEC_REMPWD2
    PARAMFMT TDS_VARCHAR, TDS_VARBINARY, ... (repeat)
    PARAMS remote server name, remote server password, ... (repeat)

<- loginack SUCCEED
    DONE FINAL
```

## extended plus encrypted password

```
-> LOGIN TDS_SEC_LOG_ENCRYPT3
        password = ""
        lpwdlen = 0
        rempwd = ""
        lrempwdlen = 0

<- loginack TDS_LOG_NEGOTIATE
    MSG TDS_MSG_SEC_ENCRYPT3
    PARAMFMT TDS_INT4, TDS_LONGBINARY, TDS_LONGBINARY
    PARAMS password cipher suite, public key (PKCS#1), nonce
    DONE FINAL

-> MSG TDS_MSG_SEC_LOGPWD3
    PARAMFMT TDS_LONGBINARY
    PARAMS encrypted message (== RSA(pubkey, nonce+password)
    MSG TDS_MSG_SEC_REMPWD3
    PARAMFMT TDS_VARCHAR, TDS_LONGBINARY, ... (repeat)
    PARAMS remote server name, encryped message (see above), ... (repeat)

<- loginack SUCCEED
    DONE FINAL
```

## data types

data types either via switch/case in ReadFrom function managing _or_
functions attached to datatype struct to consolidate functions (e.g. for
other similar packages)

### groups

DONE
format: length scale
data: status? length data
members:
- bigdatetimen
- bigtimen

DONE
format: length{1}
data: status? length{1} data
members:
- binary
- boundary
- char
- daten
- datetimen
- fltn
- intn
- uintn
- longbinary
- longchar
- moneyn
- sensitivity
- timen
- varbinary
- varchar
notes:
- if length == 0 -> null / no difference between zero-length and null

DONE
format:
data: status? data
members:
- bit{1}
- datetime{8}
- date{4}
- shortdate{4}
- flt4{4}
- flt8{8}
- int1{1}
- int2{2}
- int4{4}
- int8{8}
- interval{8}
- sint1{1}
- uint2{2}
- uint4{4}
- uint8{8}
- money{8}
- shortmoney{4}
- time{4}

DONE
format: length, precision, scale
data: status?, length, data
members:
- decn
- numn

DONE
format: blobtype, classidlength, classid
data: status, serialization type, classid length, classid, LOB Length?,
        LocatorLength?, Locator?, {datalen, data}...
members:
- blob
TODOs:
- ClassIDs as consts
- Serialization Type as consts

DONE
format: length{4}, NameLength{2}, name
data: status?, TxtPtrLen{1}, txtpr, TimeStamp{8}, DataLen{4}, data
members:
- image
- text
- unitext
- xml
notes:
- if txtprlen == 0 -> null, no other fields

format:
data:
members:
-

# capabilities

This is a !FUN! one.

TDS has three different types of capabilities: Request, response and
security.

Each of these capbility types have a number of possible capabilities
associated with them.

When passing these through `TDS_CAPABILITY` they're passed as
`ValueMask`s, which are essentially bitmasks.

Example:

Capability type `ACap` has 15 capabilities - `cap1` through `cap12`:

```go
type ACap uint8

const (
    cap1 ACap = 1
    cap2
    cap3
    cap4
    cap5
    cap6
    cap7
    cap8
    cap9
    cap10
    cap11
    cap12
)
```

This capability can be represented in a 2 byte ValueMask (hence
a `ValueMask` is a multi-byte bitmask):

```
0000 0000  0000 0000
|||| ||||  |||| |||+- cap1
|||| ||||  |||| ||+- cap2
|||| ||||  |||| |+- cap3
|||| ||||  |||| +- cap4
|||| ||||  |||+- cap5
|||| ||||  ||+ cap6
|||| ||||  |+- cap7
|||| ||||  +- cap8
|||| |||+- cap9
|||| ||+- cap10
|||| |+- cap11
|||| +- cap12
|||+- unused
||+- unused
|+- unused
+- unused

```

If cap1, cap8 and cap11 are supported the ValueMask^(tm) looks as
follows:

```
0000 0100  1000 0001
```

# ParamFmt and Params

## First approach

1. Use `FieldData` for both ParamFmt and Params
2. ParamFmt creates one `FieldData` for each column
3. Params copies each columns' `FieldData` for each field

### Problems

1. Would require each typedef to have a copy method attached
    -> either through generation or manually
2. Lots of memory wasted on large return sets

## Second approach

FieldData retains slices of data and stati

### Problems

1. Confusing in the code
2. `append()` is not performant (especially with large data sets)
3. can't preemtpively create fitting slices since TDS server does not
   respond with result set size

## Third approach

- Separate FieldData into Fmt and Data structs
- Data retains a pointer to Fmt
- Fmt structs have a method to return a referring data struct?

### Discussion

- Possible problem: Large amount for typedefs (each data type needs a fmt struct and
   a datastruct)
- Maybe one struct per combination type via embedding (as is now)
  without type-specific typedefs

### Layout

ParamFmt
    -> creates columnFmts, read format information
Params
    -> retains pointer to ParamFmt
    -> calls NewData on each columnFmt, read data
        -> NewData could access a similar function to (current) LookupFieldData to retrieve a new fieldData based on the token

# Caching for large datasets, low-overhead byte storage

## Reasons

1. `append()` is inefficient, especially with large datsets
    Every time the capacity of the underlying array is not big enough to
    handle an append request, `append()` initializes a new array and
    copies the data over - which results in considerable overhead
    Example:
        Assuming 1GB of data sent by TDS server for a single field with
        an average of 500 bytes per payload
        1048576 bytes sent in 2097 payloads, resulting in one initial
        allocation and 20196 new allocations with copy instructions
