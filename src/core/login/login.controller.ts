/* eslint-disable prettier/prettier */
import { Controller, Get, Post, Body, Patch, Param, Delete, Res, Header, Headers, HttpStatus } from '@nestjs/common';
import { LoginService } from './login.service';
import { LoginRequest } from './dto/loginRequest';
import { UpdateLoginDto } from './dto/update-login.dto';
import { Response } from 'express';

@Controller('login')
export class LoginController {
  constructor(private readonly loginService: LoginService) {}

  @Post()
  async attemptLogin(@Body() request: LoginRequest, @Res() res: Response, @Headers() headers: Record<SVGFESpecularLightingElement, string>) {
    if(!this.validateParams(request))
    res.status(HttpStatus.BAD_REQUEST).send('Invalid params to login');
    return;

    const result = await.this.loginService.attemptLogin(request);

    const status = result.success ? HttpStatus.OK : HttpStatus.UNPROCESSABLE_ENTITY;
    if(result.data.accessToken) {
      res.setHeader('Authorization', `Bearer ${result.data.accessToken}`)
    }
    res.status(status).send();

  }

  private validateParams(request: LoginRequest) {
    if(request) {
      return true;
    }
    return false;
  }

  @Get()
  findAll() {
    return this.loginService.findAll();
  }

  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.loginService.findOne(+id);
  }

  @Patch(':id')
  update(@Param('id') id: string, @Body() updateLoginDto: UpdateLoginDto) {
    return this.loginService.update(+id, updateLoginDto);
  }

  @Delete(':id')
  remove(@Param('id') id: string) {
    return this.loginService.remove(+id);
  }
}
