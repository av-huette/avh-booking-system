export interface Location {
    id?: number
    name: string
    stripeTerminalId?: string
}

export class Location implements Location {

    constructor(location: Location) {
        this.id = location.id;
        this.name = location.name;
        this.stripeTerminalId = location.stripeTerminalId;
    }
}
