while (!(docker logs obstandalone | grep -c "boot success!" 1>/dev/null 2>/dev/null))
do
    cur_dateTime=`date +"%Y-%m-%d %H:%M:%S"`
    echo $cur_dateTime" ob is booting..."
    sleep 5s
done

clear
echo -e "$(< ./tools/gitpod/banner.txt)"

echo -e "\033[32m\nOceanBase server boot success!\nNow you can use the following commands to connect to the database:\033[0m"
echo -e "\033[32m  [1] Connect with user root@sys:\n      docker exec -it obstandalone ob-mysql sys\033[0m"
echo -e "\033[32m  [2] Connect with user root@test:\n      docker exec -it obstandalone ob-mysql root\033[0m"
echo -e "\033[32m  [3] Connect with user test@test:\n      docker exec -it obstandalone ob-mysql test\033[0m"

