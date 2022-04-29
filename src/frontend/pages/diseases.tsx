import React, { useState } from 'react';
import Axios from 'axios';

export default function disease() {
  const [inputName, setInputName] = useState('');
  const [inputFile, setInputFile] = useState();
  const [inputFileName, setInputFileName] = useState('');
  const [inputPenyakit, setInputPenyakit] = useState('');

  const [hasilTes, setHasilTes] = useState([]);

  const getDiseasePrediction = (e: any) => {

    console.log("Getting data...");
    e.preventDefault();

    const formData = new FormData();
    formData.append('file', inputFile!);
    formData.append('inputName', inputName);
    formData.append('inputPenyakit', inputPenyakit);
    formData.append('tanggal', new Date().toLocaleString().split(" ")[0]);

    if (inputName !== "") {
      const s = inputFileName.split('.')
      if (s[s.length - 1] == "txt") {
        Axios.post(`${process.env.REACT_APP_DNA_API}/get-diseaseprediction`, formData)
        .then((response) => {
          setHasilTes(response.data)
          console.log(hasilTes)
        })
      } else {
        let arr = [{Status:-5}]
        setHasilTes(arr)
      }
        
    } else {
      let arr = [{Status:-4}]
      setHasilTes(arr)
    }
    console.log("selesai");
  }

  return (
    <div>
      <div id="tesDNA">
        <div>
          <h3> Tes DNA </h3>
          <div className = "column side">
            <p>Username : </p>
            <input className="inputBox" type='text' placeholder='Nama pengguna...' required 
            onChange={(e) => {
              setInputName(e.target.value)
            }}></input>
          </div>
          <div className = "column middle">
            <p>Sequence DNA :</p>
              <input className="inputBox" type="file" accept=".txt" required
              onChange={(e) => {
                setInputFile(e.target.files[0])
                setInputFileName(e.target.files[0].name)
              }}></input>
          </div>
          <div className = "column side">
            <p>Disease Prediction :</p>
            <input className="inputBox" type='text' placeholder='Nama Penyakit...' required
            onChange={(e) => {
              setInputPenyakit(e.target.value)
            }}></input>
          </div>
          <div>
            <button onClick={getDiseasePrediction}>Submit</button>
          </div>
        </div>
        <div id="hasilTes">
          <h3> Hasil </h3>
          {hasilTes.map((val, key) => {
            return (
              <div>
              {(val.Status === 0 || val.Status === 1)&&
                 <div>
                <p className="Same"> {val.TanggalPrediksi} -</p>
                <p className="Same">&nbsp; {val.NamaPasien} -</p>
                <p className="Same">&nbsp; {val.TingkatKemiripan}% -</p>
                <p className="Same">&nbsp; {val.PenyakitPrediksi} -</p>
                {val.Status == 1 && <p className="Same">&nbsp; True</p>}
                {val.Status == 0 && <p className="Same">&nbsp; False</p>}
                </div>
              } 
              {val.Status === -1 &&
                <p className="Same">DNA Not Found!</p>
              }
              {val.Status === -2 &&
                <p className="Same">Disease Not Found!</p>
              }
              {val.Status === -3 &&
                <p className="Same">Disease Not Found!</p>
              }
              {val.Status === -4 &&
                <p className="Same">Name must be filled!</p>
              }
              {val.Status === -5 &&
                <p className="Same">Please upload file format .txt!</p>
              }
              </div>
            )
          })}
        </div>
      </div>
    </div>
  )
}
