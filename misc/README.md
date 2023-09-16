```bash
docker run -v $(pwd):$(pwd) -w $(pwd) golang:1.16.15 make

VER=2309.12

docker build -f Dockerfile -t docker.io/scwatts/fusionfs_debug:${VER} .
docker push docker.io/scwatts/fusionfs_debug:${VER}
```
