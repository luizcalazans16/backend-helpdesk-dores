/* eslint-disable prettier/prettier */
import { Controller, Logger } from "@nestjs/common";
import { TicketsService } from "./tickets.service";

@Controller('/tickets')
export class TicketsController {
  private logger: Logger;

  constructor(private readonly ticketsService: TicketsService) {
    this.logger = new Logger(this.constructor.name);
  }
}
