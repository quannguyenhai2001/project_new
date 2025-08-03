import { User } from "@/lib/validations/auth";

declare module "next-auth" {
  interface Session {
    user: User & {
      id: string;
      name?: string | null;
    };
    accessToken: string;
  }

  interface User {
    id: string;
    name?: string;
    email: string;
    accessToken: string;
  }
}

declare module "next-auth/jwt" {
  interface JWT {
    id: string;
    name?: string;
    email: string;
    accessToken: string;
  }
}
