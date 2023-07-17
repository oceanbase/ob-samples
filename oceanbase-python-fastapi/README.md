# Building REST APIs using FastAPI, SQLAlchemy & Uvicorn with OceanBase Database

This project demonstrates how to build a RESTful API using FastAPI, SQLAlchemy, and Uvicorn with an OceanBase database.

## Introduction

OceanBase is a distributed relational database system developed by Alibaba Group. It is designed to provide high scalability, availability, and performance.

FastAPI is a modern, fast (high-performance), web framework for building APIs with Python 3.7+ based on standard Python type hints. FastAPI is designed to be easy to use and understand, while also being powerful and scalable.

SQLAlchemy is a popular Python library used for working with SQL databases. It provides a set of high-level API for working with relational databases, while also providing a way to work with raw SQL queries.

Uvicorn is an ASGI (Asynchronous Server Gateway Interface) server that is designed to work with async frameworks such as FastAPI.

## Project structure

The project is structured as follows:

```
.
├── README.md
├── main.py
│   ├── __init__.py
│   ├── main.py
│   ├── models.py
│   ├── repositories.py
│   ├── schemas.py
│   └── utils.py
├── database
│   ├── models.py
│   ├── repositories.py
│   └── schemas.py
├── db.py
└── requirements.txt

```

The `main.py` file contains the FastAPI application, while the `database` folder contains the database models and repositories.

## Getting Started

To run this project on your local machine, clone the project and install the necessary packages by running the following command:

```
pip install -r requirements.txt

```

Then, create a `.env` file and substitute the `DB_URL` variable based on your own database setup.

## Running the app

Run the app by typing the following command in the terminal:

```bash
python main.py
```

Now you access a running dev server on `http://127.0.0.1:9000`.
