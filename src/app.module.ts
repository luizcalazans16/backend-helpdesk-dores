/* eslint-disable prettier/prettier */
import { MiddlewareConsumer, Module, NestModule, RequestMethod } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { PreAuthMiddleware } from './common/auth/preAuthMiddleware';
import { TicketsModule } from './core/tickets/tickets.module';

@Module({
  imports: [TicketsModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule implements NestModule {
  configure(consumer: MiddlewareConsumer) {
      consumer.apply(PreAuthMiddleware).forRoutes({
        path: '*', method: RequestMethod.ALL
      });
  }
}
