
---
- List image/container:

```none
$ docker image/container ls
```

- Delete image/container:

```none
$ docker image/container rm <tên image/container >
```

- Delete all image hiện có:

```none
$ docker image rm $(docker images –a –q)
```

- List all container hiện có:

```none
$ docker ps –a
```

- Stop a container cụ thể:

```none
$ docker stop <tên container>
```

- Run container từ image và thay đổi tên container:

```none
$ docker run –name <tên container> <tên image>
```

- Stop all container:

```none
$ docker stop $(docker ps –a –q)
```

- Delete all container hiện có:

```none
$ docker rm $(docker ps –a –q)
```

- Show log a container:

```none
$ docker logs <tên container>
```

- Build một image từ container:

```none
$ docker build -t <tên container>
```

- Tạo một container chạy ngầm:

```none
$ docker run -d <tên image>
```

- Tải một image trên docker hub:

```none
$ docker pull <tên image>
```

- Start một container:

```none
$ docker start <tên container>
```

