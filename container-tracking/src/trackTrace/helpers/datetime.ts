const date = require('date-and-time');
const day_of_week = require('date-and-time/plugin/day-of-week');

export interface IDatetime {
    strptime(dateString: string, dateFormat: string): Date;
}

String.prototype.capitalizeFirstLetter = function (): string {
    return this.charAt(0).toUpperCase() + this.slice(1);
}

export class Datetime implements IDatetime {

    public strptime(dateString: string, dateFormat: string): Date {
        if (dateFormat.includes("ddd") || dateFormat.includes("d") || dateFormat.includes("dd")) {
            date.plugin(day_of_week)
        }
        return date.parse(dateString, dateFormat)
    }
}