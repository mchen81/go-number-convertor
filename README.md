# Number Convertor

# Goal

This project is to convert different bases(binary, digital, hex, etc) of number (in int32) to a format number information.

# Int32

This project is only for int32, don't put the number out range of int32

# Build file

After build, you will get **converters**

```
go build *.go
```

# Example IO

### Example 1.

```
./converters 0x10

(signed) 16
(unsigned) 16
(binary) 0b10000
(binary) 0000 0000 0000 0000 0000 0000 0001 0000
(unsigned) 0x10
```

### Example 2.

```
./converters 4294967293

(signed) -3
(unsigned) 4294967293
(binary) 0b11111111111111111111111111111101
(binary) 1111 1111 1111 1111 1111 1111 1111 1101
(unsigned) 0xfffffffd

```

### Example 3.

```
./converters 0b101010100010

(signed) 2722
(unsigned) 2722
(binary) 0b101010100010
(binary) 0000 0000 0000 0000 0000 1010 1010 0010
(unsigned) 0xaa2
```

### Example 4.

```
./converters 429496729332

429496729332 is too big or too small(for int32).
```
