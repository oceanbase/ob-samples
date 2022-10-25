#echo "export CLASSPATH=/workspace/ob-example/java/mysql-conneect-java-5.1.47.jar" >> ~/.profile
#source ~/.profile
javapath=$(cd `dirname $0`;pwd)
wget -P $javapath https://downloads.mysql.com/archives/get/p/3/file/mysql-connector-java-5.1.47.zip
unzip $javapath/mysql-connector-java-5.1.47.zip
mv $javapath/mysql-connector-java-5.1.47/mysql-connector-java-5.1.47.jar $javapath

