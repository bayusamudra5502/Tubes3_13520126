import React, { useEffect, useState } from 'react'
import axios, { type AxiosError } from 'axios'
import ICheckUp from '../model/history';

export default function history() {
  const [input, setInput] = useState('');
  const [predictionList, setPredictionList] = useState<ICheckUp[]>([]);

  const getDetailPrediction = async () => {
    const { data } = await axios.get(`http://localhost:8080/check`)

    setPredictionList(data.data.data.map((rec: ICheckUp) => {
      return {
        ...rec,
        created_timestamp: new Date(rec.created_timestamp)
      }
    }))
  }

  useEffect(() => {
    getDetailPrediction()
  }, [])

  return (
    <div id="tesDNA" className='w-[80vw]'>
      <h3> Detail Prediksi DNA </h3>
      <input type='text' placeholder='Input tanggal atau nama penyakit'
        onChange={(e) => {
          setInput(e.target.value)
        }}
      />
      <button onClick={getDetailPrediction}>Process</button>
      <h3>Hasil</h3>
      <div className='w-full'>
        {predictionList.map((val, idx) => {
          return (
            <div className='flex-row ContainerBody w-full'>
              <p>{idx + 1}.</p>
              <p className='margin-left-4'>{
                new Intl.DateTimeFormat("id", {
                  dateStyle: "long"
                }).format(val.created_timestamp)
              }</p>
              <p className='margin-left-4'>{val.name} -</p>
              <p className='margin-left-4'>{val.disease.name} -</p>
              <p className='margin-left-4'>{val.similarity}% -</p>
              {val.is_match ? <p className='margin-left-4'>True</p>
                : <p className='margin-left-4'>False</p>}
            </div>
          )
        })}
      </div>
    </div>
  )
}
