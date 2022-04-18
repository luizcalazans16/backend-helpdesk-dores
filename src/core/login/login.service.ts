/* eslint-disable prettier/prettier */
import { Injectable, Logger } from '@nestjs/common';
import { LoginRepository } from './repository/login.repository';
import { LoginRequest } from './dto/loginRequest';

@Injectable()
export class LoginService {

  private logger: Logger

  constructor(
    private readonly loginRepository: LoginRepository
  ) {
    this.logger = new Logger(this.constructor.name)
  }

  async attemptLogin(req: LoginRequest): Promise<any> {
    // eslint-disable-next-line @typescript-eslint/no-var-requires
    const firebase =  require('firebase');
    firebase.signInWithEmailAndPassword(req.email, req.password)
    .then((e, res) => {
      firebase.auth().currentUser().getIdToken(true)
      .then((idToken) => {
        res.json({idToken});
      })
    })
    .catch((err) => {
      console.log(err);
    })

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
