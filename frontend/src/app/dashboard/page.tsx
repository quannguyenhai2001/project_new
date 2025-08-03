import { redirect } from "next/navigation";
import { auth } from "@/lib/auth";

export default async function DashboardPage() {
  const session = await auth();

  // Protect this route - redirect to login if not authenticated
  if (!session?.user) {
    redirect("/login");
  }

  return (
    <div className="flex flex-col space-y-6">
      <h1 className="text-3xl font-bold">Dashboard</h1>
      <div className="rounded-lg border p-6">
        <h2 className="text-xl font-semibold mb-4">
          Xin chào, {session.user.name || session.user.email}!
        </h2>
        <p className="mb-2">Bạn đã đăng nhập thành công.</p>
        <p>
          User ID: <code>{session.user.id}</code>
        </p>
        <p>
          Email: <code>{session.user.email}</code>
        </p>
      </div>
    </div>
  );
}
