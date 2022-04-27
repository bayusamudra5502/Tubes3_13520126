import type { NextPage } from "next";
import Image from "next/image";
import Link from "next/link";
import virus from "../public/virus.png"

const Home: NextPage = () => {
  console.dir(process.env)
  return (
    <div className="bg-white py-4 px-6 rounded">
      <div>
        <div className="mb-5 mt-4 flex align-middle justify-center">
          <Image src={virus} width="150px" height="150px"></Image>
        </div>
        <h1 className="text-3xl mb-2">Disease Detector</h1>
        <p className="font-bold">Menu Utama</p>
        <div className="flex flex-col">
          <Link href="/diseases">
            <button className="p-3 w-full border-black border-2 my-2 rounded hover:text-white hover:bg-cyan-pallete transition-all ease-in-out">
              Daftar Penyakit
            </button>
          </Link>
          <Link href="/history">
            <button className="p-3 w-full border-black border-2 my-2 rounded hover:text-white hover:bg-cyan-pallete transition-all ease-in-out">
              Riwayat Pemeriksaan
            </button>
          </Link>
        </div>
      </div>
    </div>
  );
};

export default Home;
