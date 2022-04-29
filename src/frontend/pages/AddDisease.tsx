import React, { useState } from 'react';
import axios, { AxiosError } from 'axios';

export default function AddDisease() {
  const [inputName, setInputName] = useState('');
  const [inputFile, setInputFile] = useState<File | null>(null);
  const [submitted, setSubmitted] = useState(false);

  const AddDisease = async (e: any) => {
    e.preventDefault();
    try {
      const payload = new FormData();
      payload.append("name", inputName);
      payload.append("file", inputFile!);

      await axios.post(`http://localhost:8080/disease`, payload)
      setSubmitted(true);
      resetData()
    } catch (err) {
      if (err instanceof AxiosError) {
        if (err.response?.status == 400) {
          alert(`Terjadi Error : ${err.response.data.message}`)
        } else {
          alert(`Terjadi kesalahan pada server`)
        }
      }
    };
  }

  function resetData() {
    setInputName("")
  }

  return (
    <div className="bg-white py-4 px-6 rounded">
      {submitted ? (
        <div>
          <h4>Penyakit berhasil ditambahkan</h4>
          <button
            className='w-[100%] p-2 border-2 border-black mt-3'
            onClick={() => {
              resetData()
              setSubmitted(false)
            }}>OK</button>
        </div>
      ) : (
        <div>
          <h1 className="font-bold flex align-middle justify-center">Tambah Penyakit</h1>
          <form onSubmit={AddDisease}>
            <div className="flex flex-col">
              <br></br>
              <p>Nama Penyakit</p>
              <input className='p-2 border-2 border-black mb-2' required placeholder='Nama Penyakit' type="text" name="name" value={inputName} onChange={(e) => setInputName(e.target.value)}></input>
              <p>Sequence DNA</p>
              <input required type="file" name="file" onChange={(e) => {
                if (e.target.files) {
                  setInputFile(e.target.files[0])
                } else {
                  setInputFile(null)
                }
              }
              }></input>
              <br></br>
              <button className="p-3 w-full border-black border-2 my-2 rounded hover:text-white hover:bg-cyan-pallete transition-all ease-in-out">Submit</button>
            </div>
          </form>
        </div>
      )}
    </div>
  )

}
