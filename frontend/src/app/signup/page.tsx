"use client";

import * as React from "react";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import Link from "next/link";
import { signIn } from "next-auth/react";
import { SignupSchema, SignupValues } from "@/lib/validations/auth";
import { authService } from "@/lib/api";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

export default function SignupPage() {
  const router = useRouter();
  const [error, setError] = React.useState<string | null>(null);
  const [isPending, setIsPending] = React.useState<boolean>(false);

  const form = useForm<SignupValues>({
    resolver: zodResolver(SignupSchema),
    defaultValues: {
      email: "",
      password: "",
      firstName: "",
      lastName: "",
    },
  });

  async function onSubmit(values: SignupValues) {
    try {
      setError(null);
      setIsPending(true);

      // Register the user with the API
      await authService.signup(values);

      // Then sign in
      const result = await signIn("credentials", {
        redirect: false,
        email: values.email,
        password: values.password,
      });

      if (result?.error) {
        setError("Đăng ký thành công nhưng không thể đăng nhập");
        return;
      }

      router.push("/dashboard");
      router.refresh();
    } catch (error: any) {
      console.error(error);
      if (error.response?.data?.error) {
        setError(error.response.data.error);
      } else {
        setError("Có lỗi xảy ra, vui lòng thử lại");
      }
    } finally {
      setIsPending(false);
    }
  }

  return (
    <div className="mx-auto flex w-full flex-col justify-center space-y-6 sm:w-[350px]">
      <div className="flex flex-col space-y-2 text-center">
        <h1 className="text-2xl font-semibold tracking-tight">
          Đăng ký tài khoản mới
        </h1>
        <p className="text-sm text-muted-foreground">
          Nhập thông tin của bạn để đăng ký tài khoản mới
        </p>
      </div>

      <div className="grid gap-6">
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <div className="grid gap-4">
            <div className="grid grid-cols-2 gap-3">
              <div className="grid gap-2">
                <Label htmlFor="firstName">Tên</Label>
                <Input
                  id="firstName"
                  placeholder="Nguyễn"
                  disabled={isPending}
                  {...form.register("firstName")}
                />
                {form.formState.errors.firstName && (
                  <p className="text-sm text-red-500">
                    {form.formState.errors.firstName.message}
                  </p>
                )}
              </div>
              <div className="grid gap-2">
                <Label htmlFor="lastName">Họ</Label>
                <Input
                  id="lastName"
                  placeholder="Văn A"
                  disabled={isPending}
                  {...form.register("lastName")}
                />
                {form.formState.errors.lastName && (
                  <p className="text-sm text-red-500">
                    {form.formState.errors.lastName.message}
                  </p>
                )}
              </div>
            </div>

            <div className="grid gap-2">
              <Label htmlFor="email">Email</Label>
              <Input
                id="email"
                type="email"
                placeholder="name@example.com"
                autoComplete="email"
                disabled={isPending}
                {...form.register("email")}
              />
              {form.formState.errors.email && (
                <p className="text-sm text-red-500">
                  {form.formState.errors.email.message}
                </p>
              )}
            </div>

            <div className="grid gap-2">
              <Label htmlFor="password">Mật khẩu</Label>
              <Input
                id="password"
                type="password"
                placeholder="••••••••"
                autoComplete="new-password"
                disabled={isPending}
                {...form.register("password")}
              />
              {form.formState.errors.password && (
                <p className="text-sm text-red-500">
                  {form.formState.errors.password.message}
                </p>
              )}
            </div>

            {error && <p className="text-sm text-red-500">{error}</p>}

            <Button type="submit" className="w-full" disabled={isPending}>
              {isPending ? "Đang đăng ký..." : "Đăng ký"}
            </Button>
          </div>
        </form>

        <div className="text-center text-sm">
          Đã có tài khoản?{" "}
          <Link href="/login" className="font-semibold hover:underline">
            Đăng nhập
          </Link>
        </div>
      </div>
    </div>
  );
}
