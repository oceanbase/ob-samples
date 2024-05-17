export function createOceanbaseConnection() {
  return new Promise((resolve, reject) => {
    const mysql = require('mysql2');
    const connection = mysql.createConnection({
      host: '127.0.0.',
      user: 'root@test',
      port: '2881',
      password: '',
      database: 'test',
    });
    connection.connect((err) => {
      if (err) {
        reject(err);
      }
      resolve(connection);
    });
  });
}
