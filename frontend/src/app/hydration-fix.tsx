"use client";

import { useEffect } from "react";

export default function HydrationFix() {
  useEffect(() => {
    // This empty useEffect ensures client-side only code runs after hydration
  }, []);

  return null;
}
