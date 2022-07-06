import {Column, Entity, PrimaryGeneratedColumn} from "typeorm"

@Entity()
export class Unlocode {
    @PrimaryGeneratedColumn()
    id: number
    @Column()
    unlocode: string
    @Column()
    fullname: string

}
