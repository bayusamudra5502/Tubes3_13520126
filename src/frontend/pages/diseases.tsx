import React, { useState } from "react";
import IPenyakit from "../model/penyakit";
import { Table } from './Table'

export default function App() {
  const [data, setData] = useState<IPenyakit[]>([]);
  const columns = [
    { accessor: 'id', label: 'id' },
    { accessor: 'Nama_pengguna', label: 'Nama pengguna' },
    { accessor: 'Nama_penyakit', label: 'Nama penyakit' },
    { accessor: 'Tanggal_pemeriksaan', label: 'Tanggal pemeriksaan' },
    { accessor: 'Similarity', label: 'Similarity' },
    { accessor: 'Kebenaran', label: 'Kebenaran', format: (value) => (value ? 'True' : 'False') },
  ]
  const rows = data;
  return (
    <div className="App">
      <h1>Table Urutan Hasil Prediksi</h1>
      <Table rows={rows} columns={columns} />
    </div>
  )
}
