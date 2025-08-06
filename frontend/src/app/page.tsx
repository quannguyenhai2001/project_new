"use client";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { authClient } from "@/lib/auth-client";
import { useState } from "react";

export default function Home() {
  const { data: session } = authClient.useSession();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [name, setName] = useState("");
  const onSubmit = () => {
    authClient.signUp.email(
      {
        email,
        password,
        name,
      },
      {
        onSuccess: () => {
          console.log("Sign up successful");
        },
        onError: (error) => {
          console.error("Sign up failed", error);
        },
      }
    );
  };
  const onSignIn = () => {
    authClient.signIn.email(
      {
        email,
        password,
      },
      {
        onSuccess: () => {
          console.log("Sign in successful");
        },
        onError: (error) => {
          console.error("Sign in failed", error);
        },
      }
    );
  };
  if (session) {
    return (
      <div className="flex flex-col items-center justify-center min-h-screen p-8 gap-4">
        <h1 className="text-2xl">Welcome, {session.user.name}!</h1>
        <p className="text-lg">
          You are logged in with email: {session.user.email}
        </p>
        <Button onClick={() => authClient.signOut()}>Sign Out</Button>
      </div>
    );
  }
  return (
    <div className="flex flex-col items-center justify-center min-h-screen p-8 gap-4">
      <div>
        {" "}
        <Input
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <Input
          placeholder="Password"
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <Input
          placeholder="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <Button onClick={onSubmit}>Submit</Button>
      </div>
      <div>
        <Input
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <Input
          placeholder="Password"
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <Button onClick={onSignIn}>Sign in</Button>
      </div>
    </div>
  );
}
