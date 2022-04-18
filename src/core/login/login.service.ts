/* eslint-disable prettier/prettier */
import { Injectable, Logger } from '@nestjs/common';
import { LoginRepository } from './repository/login.repository';

@Injectable()
export class LoginService {

  private logger: Logger

  constructor(
    private readonly loginRepository: LoginRepository
  ) {
    this.logger = new Logger(this.constructor.name)
  }

  findAll() {
    return `This action returns all login`;
  }

  findOne(id: number) {
    return `This action returns a #${id} login`;
  }

  remove(id: number) {
    return `This action removes a #${id} login`;
  }
}
