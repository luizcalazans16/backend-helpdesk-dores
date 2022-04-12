/* eslint-disable prettier/prettier */
import { Injectable, Logger } from "@nestjs/common";

@Injectable()
export class TicketsService {
    private logger: Logger;

    constructor() {
        this.logger = new Logger(this.constructor.name);
    }
}