import Navbar from "@/components/navbar";
import { Suspense } from "react";

export default function RootLayout({
    children,
  }: Readonly<{
    children: React.ReactNode;
  }>) {
    return (
      <div lang="en">
        <Suspense fallback={<div>Loading...</div>}>
          <Navbar />
          <div>{children}</div>
        </Suspense>
      </div>
    );
  }