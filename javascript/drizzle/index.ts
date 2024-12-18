import "dotenv/config";
import { drizzle } from "drizzle-orm/mysql2";
import { usersTable } from "./db/schema";

const db = drizzle(process.env.DATABASE_URL!);

async function main() {
  const user: typeof usersTable.$inferInsert = {
    name: "Alice",
    email: "alice@oceanbase.com",
  };
  await db.insert(usersTable).values(user);

  const allUsers = await db.select().from(usersTable);
  console.log(allUsers);
}

main()
  .then(() => {
    process.exit(0);
  })
  .catch((err) => {
    console.log(err);
  });
