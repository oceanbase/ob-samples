import { DataSource } from "typeorm";
import { Entity, Column, PrimaryGeneratedColumn } from "typeorm";
import "reflect-metadata";

@Entity({ name: "users" })
export default class User {
  @Column()
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ unique: true })
  email: string;

  @Column()
  name: string;
}

const dataSource = new DataSource({
  type: "mysql",
  url: "mysql://root:@127.0.0.1:2881/test",
  entities: [User],
  synchronize: true,
});

dataSource
  .initialize()
  .then(async () => {
    const userRepository = dataSource.getRepository(User);
    await userRepository.save({
      name: "Alice",
      email: "alice@oceanbase.com",
    });
    const allUsers = await userRepository.find();
    console.log(allUsers);
    dataSource.destroy();
  })
  .catch((err) => {
    console.error(err);
    dataSource.destroy();
  });
