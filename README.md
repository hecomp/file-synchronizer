# file-synchronizer
A Cliente needs to ensure it's global BIN file is accurately managed, updated timely and proactively monitored for changes, which will help ensure compliance, avoid fines and accept all applicable customer cards. Global BIN file is created by processing various BIN files.
Below are different types of BIN files which are need from different vendors to create a Global BIN file:
### First Data BIN File
This is the "master" bin file. Edits are made to this file to create the Global BIN File.
# Build batch application from Local -
1. Create a binary image by running below command (for mac use sudo in front of the command) -
```shell script
$ GOARCH=amd64 GOOS=linux go build -i -v -o bin/filesynchronizer_batch cmd/main.go
```
2. Build docker image. Also tag version for every new build -
```shell script
$ docker build -t default-docker-virtual/filesynchronizer_batch:1.0.0 .
```
3. Build docker tag to push it to hecomp.jfrog.io -
```shell script
$ docker tag default-docker-virtual/filesynchronizer_batch:1.0.0 hecomp.jfrog.io/default-docker-virtual/filesynchronizer_batch:1.0.0
```
4. Login into docker.artifactory.homedepot.com -
```shell script
$ docker login hecomp.jfrog.io --username=hebercomp@yahoo.com --password=H@ecomp21
```
5. Push the tagged image into docker artifactory -
```shell script
$ docker push hecomp.jfrog.io/default-docker-virtual/filesynchronizer_batch:1.0.0
```
6. Pull the tag image from docker artifactory
```shell script
$ docker pull hecomp.jfrog.io/default-docker-virtual/filesynchronizer_batch:1.0.0
```
7. 
```shell script
docker run -v /first_data_bin/current:/data/first_data_bin/current -v /first_data_bin/archive:/data/first_data_bin/archive -e ENV=docker hecomp.jfrog.io/default-docker-virtual/filesynchronizer_batch
```
# Run the batch application on local -
1. Create first_data_bin folder in the root directory
```shell script
$ mkdir first_data_bin
$ cd first_data_bin
$ mkdir current
$ chmod 777 current
$ mkdir archive
$ chmod 777 archive
```
3. Place a sample FIRST_DATA_BIN file in first_data_bin/current directory
4. Run cmd/main.go
# Run the docker image of batch application on local -
3. mount the above folders on docker and run the application -
```shell script
$ docker run -v /tmp/first_data_bin/current:/home/first_data_bin/current -v /tmp/first_data_bin/archive:/home/first_data_bin/archive  -e ENV=docker hecomp.jfrog.io/default-docker-virtual/filesynchronizer_batch
```
# Run the docker image of batch application on Server
1. Login into GCP VM
```shell script
$ ssh root@10.142.0.3
```
2. Pull the docker image from docker artifactory -
```shell script
$ docker pull hecomp.jfrog.io/default-docker-virtual/filesynchronizer_batch:1.0.0
```
6. Copy sample BIN file from local to dev server (in-case BizLink is not yet setup)
```shell script
$ scp /tmp/first_data_bin/current/FIRST_DATA_BIN root@10.142.0.3:/first_data_bin/current
```
7. mount the above folders on docker and run the application -
```shell script
$ docker run -v /tmp/first_data_bin/current:/home/first_data_bin/current -v /tmp/first_data_bin/archive:/home/first_data_bin/archive  -e ENV=docker hecomp.jfrog.io/default-docker-virtual/filesynchronizer_batch
```