# Write data to OceanBase from JSON

This script is designed to import data from a JSON file and insert it into a specified table in a OceanBase Database. It uses Python and the `mysql-connector-python` library to establish a connection to the database, parse the JSON file, and execute the insert statements.

## Requirements

- Python installed on your system
- `mysql-connector-python` library installed. You can install it using the following pip command:

    ```shell
    pip install mysql-connector-python
    ```

- OceanBase server access with permissions to insert data
- A JSON file with the data to be inserted

## Configuration

Before running the script, some initial setup is required:

1. Update the `db_config` dictionary with your OceanBase server and database credentials:

    ```python
    db_config = {
        'user': 'your_username',
        'password': 'your_password',
        'host': 'localhost',
        'port': '3306',   # Change the port if necessary
        'database': 'your_database_name',
    }
    ```

2. Replace `path_to_your_json_file.json` with the actual path to your JSON file:
json_file_path = 'path_to_your_json_file.json'.

3. Modify product_id to match the identifier you want to assign to each record:
product_id = 1  # Change this to the desired ProductID

## Execution

Run the script in your Python environment. During execution, the script will:

- Open and read the JSON file.
- Establish a connection to the OceanBase database.
- Prepare the SQL insert statement.
- Iterate over each item in the JSON data, extracting and preparing the FileName and FileContent from the path and content attributes in the JSON file respectively.
- Execute the insert statement for each item, and increment the inserted_count upon each successful insertion.
- Commit the transaction to the database for data persistence.

After the script finishes running, it will print out the total number of records successfully inserted into the database.

## Troubleshooting

If any error occurs during the data insertion, the error message will be printed to the console, and the script will continue to attempt to insert the remaining records.

Should you have any issues with the database connection, please check your db_config settings and ensure that your OceanBase service is running.

## Notes

- Ensure that the table ProductFiles exists in your OceanBase database and has the appropriate columns ProductID, FileName, and FileContent.
- This script commits the transaction after all insert operations, which means that if an error occurs, any successful inserts prior to the error will still be saved in the database. To implement a rollback mechanism in case of an error, additional error handling can be added.
