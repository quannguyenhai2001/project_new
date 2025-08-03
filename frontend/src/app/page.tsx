import Link from "next/link";
import { Button } from "@/components/ui/button";

export default function Home() {
  return (
    <div className="flex flex-col items-center justify-center max-w-3xl mx-auto text-center py-12">
      <h1 className="text-4xl font-bold tracking-tight sm:text-5xl">
        Authentication Demo
      </h1>
      <p className="mt-4 text-lg text-muted-foreground max-w-xl">
        A simple authentication demo built with Next.js, shadcn/ui, and Go.
      </p>
      <div className="flex flex-col sm:flex-row gap-4 mt-8">
        <Button asChild size="lg">
          <Link href="/signup">Đăng ký ngay</Link>
        </Button>
        <Button asChild variant="outline" size="lg">
          <Link href="/login">Đăng nhập</Link>
        </Button>
      </div>
    </div>
  );
}
