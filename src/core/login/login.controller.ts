/* eslint-disable prettier/prettier */
import { Controller, Post, Body, Res, Headers, HttpStatus } from '@nestjs/common';
import { LoginService } from './login.service';
import { LoginRequest } from './dto/loginRequest';
import { Response } from 'express';

@Controller('login')
export class LoginController {
  constructor(private readonly loginService: LoginService) {}

  @Post()
  async attemptLogin(@Body() request: LoginRequest, @Res() res: Response) {
    if(!this.validateParams(request)) {
      res.status(HttpStatus.BAD_REQUEST).send('Invalid params to login');
      return;
    } else {
    const result = await this.loginService.attemptLogin(request);

    const status = result.success ? HttpStatus.OK : HttpStatus.UNPROCESSABLE_ENTITY;
    if(result.data.accessToken) {
      res.setHeader('Authorization', `Bearer ${result.data.accessToken}`)
    }
    res.status(status).send();
  }
  }

  private validateParams(request: LoginRequest) {
    if(request) {
      return true;
    }
    return false;
  }
}
