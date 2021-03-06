
import type { NextPage } from "next";
import Image from "next/image";
import Link from "next/link";
import virus from "../public/virus.png";
import diseases from "../public/disease.jpg";

const Home: NextPage = () => {
  return (
    <div className="bg-white py-10 px-12 rounded">
      <div>
        <div className="mb-5 mt-4 flex align-middle justify-center">
          <Image src={virus} width="200px" height="200px"></Image>
        </div>
        <h1>Disease Detector</h1>
        <p className="font-bold flex align-middle justify-center">Main Menu</p>
        <div className="flex flex-col">
          <Link href="/AddDisease">
            <button className="p-3 w-full border-black border-2 my-2 rounded hover:text-white hover:bg-cyan-pallete transition-all ease-in-out">
              Add Disease
            </button>
          </Link>
          <Link href="/diseases">
            <button className="p-3 w-full border-black border-2 my-2 rounded hover:text-white hover:bg-cyan-pallete transition-all ease-in-out">
              Disease Prediction
            </button>
          </Link>
          <Link href="/history">
            <button className="p-3 w-full border-black border-2 my-2 rounded hover:text-white hover:bg-cyan-pallete transition-all ease-in-out">
              Detail Prediction
            </button>
          </Link>
        </div>
      </div>
    </div>
  );
};

export default Home;