# Mini-CRM System with OceanBase, Sequelize, and Express

This repository contains a mini Customer Relationship Management (CRM) system built with [OceanBase](https://en.oceanbase.com/), Sequelize, and Express. The application allows you to perform Create, Read, Update, and Delete (CRUD) operations on contacts in an OceanBase database.

## Features

The CRM system has the following features:

1. **Create Contacts:** Add new contacts to the database. Each contact has basic information such as name, email, company, and owner.
2. **Read Contacts:** View all the contacts stored in the database. You can also read the record of one specific contact.
3. **Update Contacts:** Modify the details of a specific contact.
4. **Delete Contacts:** Remove a contact and all its related information from the database.

## Prerequisites

Before you begin, ensure you have met the following requirements:

-   You have a running OceanBase cluster. You can install OceanBase in your local environment, [spin up a virtual machine in the cloud](https://medium.com/oceanbase-database/how-to-install-oceanbase-on-an-aws-ec2-instance-step-by-step-guide-aab852c2e0a7) to run it, or use OceanBase Cloud in the AWS marketplace to set up your cluster.
-   You have Node.js and npm installed. If not, you can download and install them from the official [Node.js website](https://nodejs.org/en).

## Running the app

To run the app, follow these steps:

1. Clone this repository:

2. Navigate to the project directory:

```bash
cd sequelize-app
```

3. Install the required packages:

```bash
npm install
```

## Using the Mini-CRM System

To use the Mini-CRM System, follow these steps:

1. Start the application:

```bash
node index.js
```

2. Open your preferred tool for making HTTP requests (like Postman) and start making requests to `http://localhost:3000`.

## Configuring OceanBase Credentials

To connect the application to your own OceanBase instance, you need to replace the database credentials in the `config/config.json` file with your own.

The `config/config.json` file contains the following configuration:

```json
{
    "development": {
        "username": "your_username",
        "password": "your_password",
        "database": "your_database",
        "host": "localhost",
        "dialect": "mysql"
    },
    "test": {
        "username": "your_username",
        "password": "your_password",
        "database": "your_database",
        "host": "localhost",
        "dialect": "mysql"
    },
    "production": {
        "username": "your_username",
        "password": "your_password",
        "database": "your_database",
        "host": "localhost",
        "dialect": "mysql"
    }
}
```

Replace `"your_username"`, `"your_password"`, and `"your_database"` with your actual OceanBase database username, password, and database name. The `"host"` should be the IP address or hostname of your OceanBase instance.
