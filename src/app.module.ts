/* eslint-disable prettier/prettier */
import { MiddlewareConsumer, Module, NestModule, RequestMethod } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { PreAuthMiddleware } from './common/auth/preAuthMiddleware';
import { TicketsModule } from './core/tickets/tickets.module';
import { LoginModule } from './core/login/login.module';
import { AuthModule } from './auth/auth.module';

@Module({
  imports: [TicketsModule, LoginModule, AuthModule],
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
