// types.ts
import { PrismaClient } from '@prisma/client';
import { User, Session, AuthRequest } from '../models';

export const prisma = new PrismaClient({
  datasources: {
    db: {
      url: 'localhost:5432',
    },
  },
});

export type AuthGatewayError = Error | {
  status: number;
  message: string;
};

export type AuthGatewayResponse = {
  status: number;
  message: string;
  data?: any;
};

export type UserResponse = {
  id: number;
  email: string;
  username: string;
  firstName: string;
  lastName: string;
  role: string;
};

export type SessionResponse = {
  id: number;
  userId: number;
  token: string;
  expiresAt: Date;
};

export type AuthRequestParams = {
  username: string;
  password: string;
  email: string;
  firstName: string;
  lastName: string;
  role: string;
};

export type AuthRequestUpdate = {
  email: string;
  firstName: string;
  lastName: string;
  role: string;
};

export type AuthRequestCreate = {
  username: string;
  password: string;
  email: string;
  firstName: string;
  lastName: string;
  role: string;
};

export type AuthRequestUpdateParams = {
  id: number;
  username?: string;
  password?: string;
  email?: string;
  firstName?: string;
  lastName?: string;
  role?: string;
};

export type AuthRequestCreateParams = {
  username: string;
  password: string;
  email: string;
  firstName: string;
  lastName: string;
  role: string;
};

export type UserUpdateParams = {
  id: number;
  email: string;
  firstName: string;
  lastName: string;
  role: string;
};

export type UserCreateParams = {
  email: string;
  firstName: string;
  lastName: string;
  role: string;
};