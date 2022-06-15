import "reflect-metadata"
import {DataSource} from "typeorm"
import {Unlocode} from "./entity/Unlocode"
import {POSTGRES_HOST, POSTGRES_PORT, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DATABASE} from "../../config.json"
import {ContainerScac} from "./entity/containerScac";

export const AppDataSource = new DataSource({
    type: "postgres",
    host: POSTGRES_HOST,
    port: POSTGRES_PORT,
    username: POSTGRES_USER,
    password: POSTGRES_PASSWORD,
    database: POSTGRES_DATABASE,
    synchronize: true,
    logging: false,
    migrationsTableName: 'migrations',
    entities: [ContainerScac, Unlocode],
    migrations: [__dirname + "/migration/*{.js,.ts}"],
    subscribers: [],
})



