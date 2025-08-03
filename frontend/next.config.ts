import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  experimental: {
    // Hỗ trợ Tailwind CSS v4
    optimizeCss: true,
  },
};

export default nextConfig;
