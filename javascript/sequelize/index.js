const { Sequelize, DataTypes } = require("sequelize");

const sequelize = new Sequelize("mysql://root:@127.0.0.1:2881/test", {
  dialect: "mysql",
  logging: false,
});

const UserModel = sequelize.define(
  "users",
  {
    id: {
      field: "id",
      type: DataTypes.BIGINT,
      autoIncrement: true,
      primaryKey: true,
      allowNull: false,
    },
    email: {
      field: "email",
      type: DataTypes.STRING,
      unique: true,
      allowNull: false,
    },
    name: {
      field: "name",
      type: DataTypes.STRING,
      allowNull: false,
    },
  },
  {
    timestamps: false,
  }
);

async function main() {
  await UserModel.sync();
  await UserModel.create({
    name: "Alice",
    email: "alice@oceanbase.com",
  });
  const allUsers = await UserModel.findAll();
  console.log(JSON.stringify(allUsers, null, 2));
}

main()
  .then(async () => {
    await sequelize.close();
  })
  .catch(async (e) => {
    console.error(e);
    await sequelize.close();
    process.exit(1);
  });
