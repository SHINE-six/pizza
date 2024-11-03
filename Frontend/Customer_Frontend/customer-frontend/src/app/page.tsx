// import Navbar from '@/components/navbar';
import './globals.css';
import Link from "next/link";


export default function Home() {

  return (
    <>
      <main className="flex flex-col items-center justify-center backgroundImage fillPage h-full">
        <Link href={'./menu'}><div className="text-5xl font-extrabold bg-tertiary p-[1rem] uppercase rounded-[2rem] shadow-2xl">Start Order</div></Link>
      </main>
    </>
  );
}
