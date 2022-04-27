import type { NextPage } from "next";
import Image from "next/image";
import virus from "../public/virus.png"

const Home: NextPage = () => {
  return (
    <div className="bg-white p-4 rounded">
      <div>
        <div className="mb-5 mt-4 flex align-middle justify-center">
          <Image src={virus} width="150px" height="150px"></Image>
        </div>
        <h1 className="text-3xl">Disease Detector</h1>
      </div>
    </div>
  );
};

export default Home;
