import React, { useState } from 'react';
import Axios from 'axios';

export default function AddDisease() {
  const [inputName, setInputName] = useState('');
  const [inputFile, setInputFile] = useState();
  const [inputFileName, setInputFileName] = useState('');
  const [submitted, setSubmitted] = useState(false);
  const [success, setSuccess] = useState(false);

  const AddDisease = async (e: any) => {
    console.log("Getting data...");
    e.preventDefault();

    const formData = new FormData();
    formData.append('nama', inputName)
    formData.append('file', inputFile!);

    if (inputName != "") {
      const s = inputFileName.split('.')
      if (s[s.length - 1] == "txt") {
        try {
          const {data} = await Axios.post("http://localhost:8080/disease", formData)
        }
        catch (err) {

        };
        setSubmitted(true);
      }
      }
    }
  }

  return (
    <div className="bg-white py-4 px-6 rounded">
    { submitted? (
      <div> 
        <h4>Penyakit berhasil ditambahkan</h4>
        <button onClick={this.reset}>OK</button>
      </div>
    ) : (
    <div>
      <h1 className="font-bold flex align-middle justify-center">Tambah Penyakit</h1>
      <form>
        <div className="flex flex-col">
          <br></br>
          <p>Nama Penyakit</p>
          <input type="text" id="name" name="name" onChange={this.onChangeName}></input>
          <p>Sequence DNA</p>
          <input type="file" id="sequence" name="file" onChange={this.onChangeFile}></input>
          <br></br>
          <input className="p-3 w-full border-black border-2 my-2 rounded hover:text-white hover:bg-cyan-pallete transition-all ease-in-out" onClick={this.onSubmit}></input>
        </div>
      </form>
    </div>
    )}
  </div>
  )

  constructor(props: Props) {
    super(props);
    this.state = {
      name: "",
      file:null,
      submitted: false
    }
    this.onChangeName = this.onChangeName.bind(this);
    this.onChangeFile = this.onChangeFile.bind(this);
    this.onSubmit = this.onSubmit.bind(this);
    this.reset = this.reset.bind(this);
  }

  onChangeName(e) {
    this.setState({name: e.target.value});
  }

  onChangeFile(e) {
    this.setState({file: e.target.files[0]})
  }

  onSubmit() {
    const formData = new FormData();
    formData.append("name", this.state.name);
    formData.append("file", this.state.file);
    console.log(formData);
    axios.post("http://localhost:8080/disease", formData, {
      headers: {
        'Content-Type' : 'multipart/form-data',
        'Access-Control-Allow-Origin' : '*'
      }
    }).then((response: any) => {
      console.log(response);
    }
    );
    this.setState({submitted: true});
  }

  reset() {
    this.setState( {
      name: "",
      file: null,
      submitted: false
    });
  }

  render() {
    const { submitted } = this.state;
    return (
      <div className="bg-white py-4 px-6 rounded">
        { submitted? (
          <div> 
            <h4>Penyakit berhasil ditambahkan</h4>
            <button onClick={this.reset}>OK</button>
          </div>
        ) : (
        <div>
          <h1 className="font-bold flex align-middle justify-center">Tambah Penyakit</h1>
          <form>
            <div className="flex flex-col">
              <br></br>
              <p>Nama Penyakit</p>
              <input type="text" id="name" name="name" onChange={this.onChangeName}></input>
              <p>Sequence DNA</p>
              <input type="file" id="sequence" name="file" onChange={this.onChangeFile}></input>
              <br></br>
              <input className="p-3 w-full border-black border-2 my-2 rounded hover:text-white hover:bg-cyan-pallete transition-all ease-in-out" onClick={this.onSubmit}></input>
            </div>
          </form>
        </div>
        )}
      </div>
    )
  }

  
}