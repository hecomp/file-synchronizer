# file-synchronizer
A Cliente needs to ensure it's global BIN file is accurately managed, updated timely and proactively monitored for changes, which will help ensure compliance, avoid fines and accept all applicable customer cards. Global BIN file is created by processing various BIN files.
Below are different types of BIN files which are need from different vendors to create a Global BIN file:
### First Data BIN File
This is the "master" bin file. Edits are made to this file to create the Global BIN File.
# Build batch application from Local -
1. Create a binary image by running below command (for mac use sudo in front of the command) -
```shell script
$ GOARCH=amd64 GOOS=linux go build -i -v -o bin/filesynchronizer_batch/main.go
```
2. Build docker image. Also tag version for every new build -
```shell script
$ docker build -t filesynchronizer_batch:1.0.0 .
```
3. Build docker tag to push it to hecomp.jfrog.io -
```shell script
$ docker tag filesynchronizer_batch:1.0.0 hecomp.jfrog.io/default-docker-virtual/default-docker-virtual:filesynchronizer_batch:1.0.0
```
4. Login into docker.artifactory.homedepot.com -
```shell script
$ docker login hecomp.jfrog.io --username=<LDAP> --password=<Authentication API Key>
```
5. Push the tagged image into docker artifactory -
```shell script
$ docker push hecomp.jfrog.io/default-docker-virtual/default-docker-virtual:filesynchronizer_batch:1.0.0
```
## Starting Cockroach
This application utilizes Docker to run a local instance of Cockroach on your machine.
Docker desktop is used to help make this possible.
You can download Docker Desktop [here](https://www.docker.com/products/docker-desktop).
\***Note for Windows users**\* Ensure that the line endings for the `build/local/cockroach/init_cockroach.sh`
and `build/local/cockroach/init_db.sh` files are set to `LF`.  The default for Windows is `CLRF`.
```shell script
# starts the database
$ docker-compose -f ./build/local/cockroach/docker-compose.yml up --build -d
# stops the database
$ docker-compose -f ./build/local/cockroach/docker-compose.yml down
# connects to SQL client
$ cockroach sql --certs-dir=./build/local/cockroach/volumes/certs --host=localhost:26257
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
$ docker run -v /tmp/secrets:/home/secrets -v /tmp/chase_bin/current:/home/chase_bin/current -v /tmp/chase_bin/archive:/home/chase_bin/archive -v /tmp/first_data_bin/current:/home/first_data_bin/current -v /tmp/first_data_bin/archive:/home/first_data_bin/archive -v /tmp/amex_bin/current:/home/amex_bin/current -v /tmp/amex_bin/archive:/home/amex_bin/archive -v /tmp/global_bin/current:/home/global_bin/current -v /tmp/global_bin/archive:/home/global_bin/archive -e COCKROACH_DB_SSL_ROOT_CERT_PATH=/home/secrets/ca.crt -e COCKROACH_DB_SSL_CLIENT_KEY_PATH=/home/secrets/client.payment_tokenization_application_user.key -e COCKROACH_DB_SSL_CLIENT_CERT_PATH=/home/secrets/payment_tokenization_application_user.cer -e COCKROACH_DB_APP_USER_NAME_TOKENIZATION=payment_tokenization_application_user -e COCKROACH_DB_HOST_PRIMARY=172.18.179.251 -e COCKROACH_DB_HOST_FALLBACK=lnc138fb.homedepot.com -e COCKROACH_DB_PORT_PRIMARY=26257 -e COCKROACH_DB_PORT_FALLBACK=26257 -e ENV=docker payment-tokenization/bin-file-management-batch:1.0.0
```
# Run the docker image of batch application on Server
1. Login into Dev Ephemeral server
```shell script
$ ssh root@ld19706.homedepot.com
```
2. Pull the docker image from docker artifactory -
```shell script
$ docker pull docker.artifactory.homedepot.com/payment-tokenization/bin-file-management-batch:1.0.0
```
3. Execute below steps to create secrets directory in dev server-
```shell script
$ cd /root
$ mkdir bin_managemnet_files
$ chmod 777 bin_managemnet_files
$ cd bin_managemnet_files
$ mkdir secrets
$ chmod 777 secrets
```
4. Execute below steps to create various BIN directories and to place the files from BizLink-
```shell script
$ cd /home/ftfadm/
$ mkdir chase_bin
$ chmod 777 chase_bin
$ cd chase_bin
$ mkdir current
$ mkdir archive
$ chmod 777 current
$ chmod 777 archive
$ cd ..
$ mkdir first_data_bin
$ cd first_data_bin
$ mkdir current
$ chmod 777 current
$ mkdir archive
$ chmod 777 archive
$ cd ..
$ mkdir amex_bin
$ cd amex_bin
$ mkdir current
$ chmod 777 current
$ mkdir archive
$ chmod 777 archive
$ cd ..
$ mkdir global_bin
$ cd global_bin
$ mkdir current
$ chmod 777 current
$ mkdir archive
$ chmod 777 archive
```
5. Copy cockroach certs from local to dev server -
```shell script
$ scp /tmp/secrets/payment_bin_range_application_user.cer root@ld19706.homedepot.com:/root/bin_managemnet_files/secrets
$ scp /tmp/secrets/client.payment_bin_range_application_user.key root@ld19706.homedepot.com:/root/bin_managemnet_files/secrets
$ scp /tmp/secrets/ca.crt root@ld19706.homedepot.com:/root/bin_managemnet_files/secrets
```
6. Copy sample BIN file from local to dev server (in-case BizLink is not yet setup)
```shell script
$ scp /tmp/chase_bin/current/CHASE_GB_BIN root@ld19706.homedepot.com:/home/ftfadm/chase_bin/current
$ scp /tmp/first_data_bin/current/FIRST_DATA_BIN root@ld19706.homedepot.com:/home/ftfadm/first_data_bin/current
$ scp /tmp/amex_bin/current/AMEX_GB_BIN root@ld19706.homedepot.com:/home/ftfadm/amex_bin/current
```
7. mount the above folders on docker and run the application -
```shell script
$ docker run -v /root/bin_management_files/secrets:/data/secrets -v /home/ftfadm/chase_bin/current:/data/chase_bin/current -v /home/ftfadm/chase_bin/archive:/data/chase_bin/archive -v /home/ftfadm/first_data_bin/current:/data/first_data_bin/current -v /home/ftfadm/first_data_bin/archive:/data/first_data_bin/archive -v /home/ftfadm/amex_bin/current:/data/amex_bin/current -v /home/ftfadm/amex_bin/archive:/data/amex_bin/archive -v /home/ftfadm/global_bin/current:/data/global_bin/current -v /home/ftfadm/global_bin/archive:/data/global_bin/archive -e COCKROACH_DB_SSL_ROOT_CERT_PATH=/data/secrets/ca.crt -e COCKROACH_DB_SSL_CLIENT_KEY_PATH=/data/secrets/client.payment_bin_range_application_user.key -e COCKROACH_DB_SSL_CLIENT_CERT_PATH=/data/secrets/payment_bin_range_application_user.cer -e COCKROACH_DB_APP_USER_NAME_TOKENIZATION=payment_bin_range_application_user -e COCKROACH_DB_HOST_PRIMARY=lnc138fb.homedepot.com -e COCKROACH_DB_PORT_PRIMARY=26257 -e COCKROACH_DB_HOST_FALLBACK=lnc138fb.homedepot.com -e COCKROACH_DB_PORT_FALLBACK=26257 -e ENV=docker payment-tokenization/bin-file-management-batch:1.0.0
```