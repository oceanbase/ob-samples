import { mysqlTable, bigint, varchar } from "drizzle-orm/mysql-core";

export const usersTable = mysqlTable("users", {
  id: bigint({ mode: "bigint" }).autoincrement().primaryKey(),
  email: varchar({ length: 255 }).notNull().unique(),
  name: varchar({ length: 255 }).notNull(),
});
