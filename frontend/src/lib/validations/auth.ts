import { z } from "zod";

// User schema
export const UserSchema = z.object({
  id: z.number(),
  email: z.string().email(),
  firstName: z.string().optional(),
  lastName: z.string().optional(),
  createdAt: z.string().or(z.date()),
});

export type User = z.infer<typeof UserSchema>;

// Signup schema
export const SignupSchema = z.object({
  email: z.string().email({ message: "Email không hợp lệ" }),
  password: z.string().min(6, { message: "Mật khẩu phải có ít nhất 6 ký tự" }),
  firstName: z.string().optional(),
  lastName: z.string().optional(),
});

export type SignupValues = z.infer<typeof SignupSchema>;

// Login schema
export const LoginSchema = z.object({
  email: z.string().email({ message: "Email không hợp lệ" }),
  password: z.string().min(1, { message: "Mật khẩu là bắt buộc" }),
});

export type LoginValues = z.infer<typeof LoginSchema>;

// API response schema for authentication
export const AuthResponseSchema = z.object({
  accessToken: z.string(),
  user: UserSchema,
});

export type AuthResponse = z.infer<typeof AuthResponseSchema>;
