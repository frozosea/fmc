import "reflect-metadata"
import {DataSource} from "typeorm"
import {Unlocode} from "./entity/Unlocode"
import {ContainerScac} from "./entity/containerScac";
import {config} from "dotenv";

config()
export const AppDataSource = new DataSource({
    type: "postgres",
    host: process.env.POSTGRES_HOST,
    port: Number(process.env.POSTGRES_PORT),
    username: process.env.POSTGRES_USER,
    password: process.env.POSTGRES_PASSWORD,
    database: process.env.POSTGRES_DATABASE,
    synchronize: true,
    logging: false,
    migrationsTableName: 'migrations',
    entities: [ContainerScac, Unlocode],
    migrations: [__dirname + "/schema/*{.js,.ts}"],
    subscribers: [],
})

AppDataSource.initialize()
    .then(async (_) => {
    })
    .catch((error) => console.log("Error of init db: ", error))



