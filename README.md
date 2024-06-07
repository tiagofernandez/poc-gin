# Quickstart

Requirements:

- [DirEnv](https://direnv.net/)
- [Docker](https://www.docker.com/products/docker-desktop/)
- [GVM](https://github.com/moovweb/gvm#installing)

```sh
export GOVERSION=$(awk '/^go / {print $1$2}' go.mod)

gvm install $GOVERSION
gvm use $GOVERSION --default  # Optional

make install
make run-dev
```
