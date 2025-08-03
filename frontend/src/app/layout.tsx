import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import HydrationFix from "./hydration-fix";

const inter = Inter({
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Next.js App",
  description: "Created with Next.js",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>{children}</body>
    </html>
  );
}

// <html lang="en" suppressHydrationWarning>
//   <body className={inter.className} suppressHydrationWarning>
//     <HydrationFix />
//     {children}
//   </body>
// </html>;
