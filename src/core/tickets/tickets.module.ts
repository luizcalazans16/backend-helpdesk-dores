/* eslint-disable prettier/prettier */
import { Module } from '@nestjs/common';
import { TicketsController } from './tickets.controller';
import { TicketsService } from './tickets.service';

@Module({
    providers: [TicketsService],
    controllers: [
        TicketsController
    ]
})
export class TicketsModule { }
