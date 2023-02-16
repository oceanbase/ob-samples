# ob-example
Oceanbase build its online platform for fast use based on gitpod, Clike the button to use

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/oceanbase/ob-example)

OceanBase will auto deploy a local observer. Please wait until see "Oceanbase boot Success" on the left terminal and try to connect in the right.
![示意图](./tools/scripts/gitpod1.png)

Examples for different language to connect Oceanbase are presented in this repository and those with 'run.sh' can run online. Three steps help to have a fast experience. 
```
//1. open 
cd xxxx
//2. prepare relative env
sh env.sh
//3. run to query
sh run.sh
```
Here we use python3-pymysql as example and others are the same:
```
cd python3-pymysql
sh env.sh
sh run.sh
```
Simple only-read sysbench for fun is also supported in gitpod!
```
docker exec -it obstandalone obd test sysbench obcluster
```
Refer to the [OceanBase](https://open.oceanbase.com) for more details about OceanBase Database.

# Contributors
- go_sql_driver [@akaError]
- jdbc [@akaError]
- pymysql [@akaError]
- Hibernate
- MyBatis
- SprintJDBC

to be add...


