while (!(docker logs obstandalone | grep -c "boot success!" 1>/dev/null 2>/dev/null))
do
    cur_dateTime=`date +"%Y-%m-%d %H:%M:%S"`
    echo $cur_dateTime" ob is booting..."
    sleep 5s
done
echo -e "\033[32m\nob boot success! use following to connect:\033[0m"
echo -e "\033[32m connect with user root@sys: docker exec -it obstandalone ob-mysql sys\033[0m"
echo -e "\033[32m connect with user root@test: docker exec -it obstandalone ob-mysql root\033[0m"
echo -e "\033[32m connect test user test@test: docker exec -it obstandalone ob-mysql test\033[0m"

