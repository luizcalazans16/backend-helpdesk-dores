import { PartialType } from '@nestjs/mapped-types';
import { LoginRequest } from './loginRequest';

export class UpdateLoginDto extends PartialType(LoginRequest) {}
