# UNIX System

## stdin, stdout, stderr
- UNIX system considers everything as a file
- file descriptors (positive int)
## UNIX processes
- process contains instructions, user data and system data parts, and other types of resources
  - running process is uniquely identified by an unsigned int which is called *process ID*
  - User process
  - Daemon process
  - Kernel process
- program is a binary file that contains instructions and data that are used for initializing the instruction and user data parts of a process

## UNIX signals

## io.Reader & io.Writer

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

```go
type Writer interface {
	Write(p []byre) (n int, err error)
}
```

