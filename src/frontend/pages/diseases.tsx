import React, { useState } from "react";
import IPenyakit from "../model/penyakit";

export default function Diseases() {
  const [data, setData] = useState<IPenyakit[]>([]);

  return <div className="bg-white py-4 px-6 rounded">
    <h1>Daftar Penyakit</h1>
    <div>
      <table>
        <thead>
          <th>
            <td>ID</td>
            <td>Nama Penyakit</td>
            <td>Waktu Penambahan</td>
          </th>
        </thead>
        <tbody>
        </tbody>
      </table>
    </div>
  </div>;
}
