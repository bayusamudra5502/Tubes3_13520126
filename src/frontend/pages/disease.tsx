import React, { useState } from "react";

export default function App() {
  return (
    <div className="bg-white py-4 px-6 rounded">
      <h1 className="font-bold flex align-middle justify-center">Tambah Penyakit</h1>
      <form>
        <div className="flex flex-col">
          <br></br>
          <p>Nama Penyakit</p>
          <input type="text" id="name" name="name"></input>
          <p>Sequence DNA</p>
          <input type="file" id="sequence" name="file"></input>
          <br></br>
          <input className="p-3 w-full border-black border-2 my-2 rounded hover:text-white hover:bg-cyan-pallete transition-all ease-in-out" type="submit"></input>
        </div>
      </form>
    </div>
  )
}
