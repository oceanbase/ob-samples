# Full Stack App with SvelteKit and OceanBase

This project demonstrates how to build a fullstack app using SvelteKit for the frontend and OceanBase as the database. The application is a simple Todo app that allows users to create, read, update, and delete tasks.

## Technologies Used

- SvelteKit: A framework for building web applications of all sizes.
- OceanBase: A distributed relational database management system developed by Ant Group.
- MySQL2: A MySQL driver for Node.js.
- Tailwind CSS: A utility-first CSS framework for rapidly building custom user interfaces.

## Setup

1. Clone the repository to your local machine.

2. Navigate to the project directory.

   ```
   cd applications/sveltekit
   ```

3. Install the necessary dependencies.

   ```
   npm i
   ```

4. Update the [db.js](./src/lib/db.js) with your own OceanBase database credentials.

5. Start the development server.
   ```
   npm run dev
   ```
6. Open your browser and navigate to `http://localhost:3000` to view the app.
