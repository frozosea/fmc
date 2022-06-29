import {MigrationInterface, QueryRunner} from "typeorm";
import {UnlocodesSeeder} from "../seeders/unlocodesSeeder/UnlocodeSeeder";
import {Unlocode} from "../entity/Unlocode";
import {config} from "dotenv";

export class migration1655113995221 implements MigrationInterface {
    name = 'migration1655113995221'
    public async up(queryRunner: QueryRunner): Promise<void> {
        config()
        console.log(`database : ${process.env.POSTGRES_DATABASE}`)
        await queryRunner.query(`CREATE TABLE IF NOT EXISTS "unlocode"
                                 (
                                     "id"       SERIAL            NOT NULL,
                                     "unlocode" character varying NOT NULL,
                                     "fullname" character varying NOT NULL,
                                     CONSTRAINT "PK_26fa5ecf49e498ebc5c0495d225" PRIMARY KEY ("id")
                                 )`);
        await queryRunner.query(`CREATE TABLE IF NOT EXISTS "container_scac"
                                 (
                                     "id"        SERIAL            NOT NULL,
                                     "container" character varying NOT NULL,
                                     "scac"      character varying NOT NULL,
                                     CONSTRAINT "UQ_b0dee459ca10815467c9ead1a87" UNIQUE ("container"),
                                     CONSTRAINT "PK_31e244382d66d26661accc1ff26" PRIMARY KEY ("id")
                                 )`);
        let unlocodeSeeder = new UnlocodesSeeder()
        let unlocodes = unlocodeSeeder.getUnlocodes()
        for (let array of unlocodes) {
            try {
                let unloObj = new Unlocode()
                unloObj.unlocode = array[0]
                unloObj.fullname = array[1]
                await queryRunner.manager.save(unloObj,{})
            } catch (e) {
            }
        }
    }

    public async down(queryRunner: QueryRunner): Promise<void> {
        await queryRunner.query(`DROP TABLE "container_scac"`);
        await queryRunner.query(`DROP TABLE "unlocode"`);
    }

}
