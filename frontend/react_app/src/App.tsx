import { useState,useCallback, useEffect } from 'react'
import './App.css'
import Editor from '@monaco-editor/react'
import { Button } from "@/components/ui/button"
import axios from 'axios'

function App() {
  const [code, setCode] = useState('#type here')
  const [language, setLanguage] = useState('python')
  const [output, setResult] = useState('')

  useEffect(()=>{
    setLanguage('python')
  },[])

  const handleCodeChange = useCallback((value:string|undefined) =>{
    if(value) setCode(value)
  },[])

  const handleRun = useCallback(async ()=>{
      await axios.post("http://localhost:8001/run",{
        code
      }).then((resp) =>{
        setResult(resp.data.output || 'no output')
      }).catch((err) =>{
        if(err) console.error(err)
      })
  },[code,setResult]);

  return (
    <div className="w-full">
      <div className="w-full flex justify-center">
        <Button type="button" className='button' onClick={handleRun}>Run</Button>
      </div>
      <div className="flex justify-between">  
        <div className="w-1/2 h-screen flex flex-col">
          <div className='flex-1'>
            <Editor
              height="100%"
              defaultLanguage="python"
              language={language}
              value={code}
              onChange={handleCodeChange}
              theme="vs-dark"
              />
          </div>
        </div>
        <div className="w-1/2">
          <p>Result: {output}</p>
        </div>
      </div>
      <div className="w-full flex justify-center">
        copyright Ainun Jariya
      </div>
    </div>
  )
}

export default App
