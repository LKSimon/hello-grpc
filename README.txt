测试 grpc

编译 docker images:
sudo docker build -t hello-grpc-server:v0.0.1 .

运行容器：
sudo docker run -p 50051:50051   hello-grpc-server:v0.0.1

进入容器：
 sudo docker exec -it f1d9a3422d22 /bin/bash

执行：
./bin/grpcurl -plaintext  localhost:50051 list