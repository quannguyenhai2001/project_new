import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { ThemeProvider } from "@/providers/theme-provider";
import { ThemeToggle } from "@/components/theme-toggle";

const inter = Inter({
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Authentication Demo",
  description: "Authentication demo with Next.js and Go",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body className={inter.className}>
        <ThemeProvider
          attribute="class"
          defaultTheme="system"
          enableSystem
          disableTransitionOnChange
        >
          <div className="min-h-screen bg-background">
            <header className="border-b bg-background">
              <div className="container flex h-16 items-center justify-between">
                <div className="font-semibold">Auth Demo</div>
                <ThemeToggle />
              </div>
            </header>
            <main className="container py-6">{children}</main>
          </div>
        </ThemeProvider>
      </body>
    </html>
  );
}
