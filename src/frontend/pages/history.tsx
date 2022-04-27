import React, { useState } from "react";
import ICheckUp from "../model/history";

export default function History() {
  const [data, setData] = useState<ICheckUp[]>([]);

  return <div className="bg-white py-4 px-6 rounded">
    <h1>Riwayat Pemeriksaan</h1>
  </div>;
}
