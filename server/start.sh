docker build --build-arg PORT=8081 -t test1 .
docker build --build-arg PORT=8082 -t test2 .
docker build --build-arg PORT=8083 -t test3 .

docker network create mynet

docker run --rm --net mynet -d -p 8081:8081 -it test1
docker run --rm --net mynet -d -p 8082:8082 -it test2
docker run --rm --net mynet -d -p 8083:8083 -it test3

# or do the following in terminal
# docker run --net mynet -p 8081:8081 -it test1
# docker run --net mynet -p 8082:8082 -it test2
# docker run --net mynet -p 8083:8083 -it test3