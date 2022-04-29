import React, { useState } from 'react';
import axios, { AxiosError } from 'axios';

interface IHasilTes {
  duration: number,
  id: number,
  disease_name: string,
  disease_id: number,
  is_match: boolean,
  similarity: number
}

export default function disease() {
  const [inputName, setInputName] = useState('');
  const [inputFile, setInputFile] = useState<File | null>(null);
  const [status, setStatus] = useState(null);
  const [inputPenyakit, setInputPenyakit] = useState('');

  const [hasilTes, setHasilTes] = useState<IHasilTes | null>(null);

  const getDiseasePrediction = async (e: any) => {
    e.preventDefault();
    try {
      const formData = new FormData();
      formData.append('file', inputFile!);
      formData.append('name', inputName);
      formData.append('disease_id', inputPenyakit);

      const { data } = await axios.post(`http://localhost:8080/check`, formData)
      setHasilTes(data.data)

    } catch (err) {
      if (err instanceof AxiosError) {

      }
    }
  }

  return (
    <div>
      <div id="tesDNA">
        <div>
          <h3> Tes DNA </h3>
          <div className="column side">
            <p>Username : </p>
            <input className="inputBox" type='text' placeholder='Nama pengguna...' required
              onChange={(e) => {
                setInputName(e.target.value)
              }}></input>
          </div>
          <div className="column middle">
            <p>Sequence DNA :</p>
            <input className="inputBox" type="file" accept=".txt" required
              onChange={(e) => {
                if (e.target.files) {
                  setInputFile(e.target.files[0])
                }
              }}></input>
          </div>
          <div className="column side">
            <p>Disease Prediction :</p>
            <input className="inputBox" type='text' placeholder='ID Penyakit' required
              onChange={(e) => {
                setInputPenyakit(e.target.value)
              }}></input>
          </div>
          <div>
            <button onClick={getDiseasePrediction}>Submit</button>
          </div>
        </div>
        <div id="hasilTes">
          <h3> Hasil Pemeriksaan Terbaru </h3>
          <div>
            {(!status) ?
              <div>
                <p className="Same"> {hasilTes?.id} -</p>
                <p className="Same">&nbsp; {inputName} -</p>
                <p className="Same">&nbsp; {hasilTes?.similarity! * 100}% -</p>
                <p className="Same">&nbsp; {hasilTes?.disease_name} (ID {hasilTes?.disease_id}) -</p>
                {hasilTes?.is_match ? <p className="Same">&nbsp; True</p> : <p className="Same">&nbsp; False</p>}
              </div>
              :
              <p className="Same">DNA Not Found!</p>
            }
          </div>

        </div>
      </div>
    </div>
  )
}
