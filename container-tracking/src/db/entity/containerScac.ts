import {Entity, PrimaryGeneratedColumn, Column, Unique} from "typeorm"
import {SCAC_TYPE} from "../../types";

@Entity()
@Unique(["container"])
export class ContainerScac {
    @PrimaryGeneratedColumn()
    id: number
    @Column()
    container: string
    @Column()
    scac: SCAC_TYPE

}
