/* Standard C++ includes */
#include <iostream>
#include <memory>

#include <mysql/jdbc.h>

#define DEFAULT_URI "tcp://127.0.0.1:2881"
#define EXAMPLE_USER "root@test"
#define EXAMPLE_PASSWORD ""
#define EXAMPLE_DB "test"

using namespace std;

int main(int argc, char *argv[]) {
    try {
        const char   *url = (argc > 1 ? argv[1] : DEFAULT_URI);
        const string user(argc >= 3 ? argv[2] : EXAMPLE_USER);
        const string password(argc >= 4 ? argv[3] : EXAMPLE_PASSWORD);
        const string database(argc >= 5 ? argv[4] : EXAMPLE_DB);
        sql::mysql::MySQL_Driver *driver = sql::mysql::get_mysql_driver_instance();

        cout << "Connected to OceanBase server..." << endl;
        unique_ptr< sql::Connection > con{driver->connect(url, user, password)};
        
        
        con->setSchema(database);
        
        unique_ptr< sql::Statement > stmt{con->createStatement()};
        
        // Drop table if exists
        cout << "Dropping table if exists..." << endl;
        stmt->execute("DROP TABLE IF EXISTS `t_test`");
        
        // Create table
        cout << "Creating table..." << endl;
        stmt->execute("CREATE TABLE `t_test` ("
                     "    `id`   int(10) NOT NULL AUTO_INCREMENT,"
                     "    `name` varchar(20) DEFAULT NULL,"
                     "    PRIMARY KEY (`id`)"
                     ") ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE = utf8_bin");
        
        // Insert data
        cout << "Inserting data..." << endl;
        stmt->execute("INSERT INTO `t_test` VALUES (default, 'Hello OceanBase')");
        
        // Query data
        cout << "Querying data..." << endl;
        std::unique_ptr< sql::ResultSet > res{stmt->executeQuery("SELECT * FROM `t_test`")};
        
        cout << "Query results:" << endl;
        while (res->next()) {
            cout << "id: " << res->getInt("id") 
                 << ", name: " << res->getString("name") << endl;
        }
        
        cout << "Done!" << endl;
    }
    catch (sql::SQLException &e)
  {
    /*
      The JDBC API throws three different exceptions:
 
    - sql::MethodNotImplementedException (derived from sql::SQLException)
    - sql::InvalidArgumentException (derived from sql::SQLException)
    - sql::SQLException (derived from std::runtime_error)
    */
 
    cout << "# ERR: SQLException in " << __FILE__;
    cout << "(" << "EXAMPLE_FUNCTION" << ") on line " << __LINE__ << endl;
 
    /* Use what() (derived from std::runtime_error) to fetch the error message */
 
    cout << "# ERR: " << e.what();
    cout << " (MySQL error code: " << e.getErrorCode();
    cout << ", SQLState: " << e.getSQLState() << " )" << endl;
 
    return EXIT_FAILURE;
  }
 

  return EXIT_SUCCESS;
}