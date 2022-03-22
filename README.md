# TestsGenerics

## Build binary
```sh
go build -o 40chairs
```

## Start Root node
Root functionality should pass in ring.
```sh
./40chairs # default ID is 7
```

## Join Root node
```sh
./40chairs -id=8 7
```

