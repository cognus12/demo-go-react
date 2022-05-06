import { useState, FC } from 'react'
import logo from './logo.svg'
import './App.css'

const HelloFromApi: FC = () => {
  const [hello, setHello] = useState('')
  const [loading, setLoading] = useState(false)

  const fetchHello = async () => {
    setLoading(true)

    const response = await fetch(`${location.origin}/api/hello`)

    const data = await response.text()

    setHello(data)

    setLoading(false)
  }

  return (
    <>
      <p>
        <button type="button" onClick={fetchHello} disabled={loading}>
          Get hello from api
        </button>
      </p>
      <p style={{ height: "39px" }}>
        {hello ? hello : null}
      </p>
    </>
  )
} 

const App: FC = () => {
  const [count, setCount] = useState(0)
  
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>Hello Vite + React!</p>
        <p>
          <button type="button" onClick={() => setCount((count) => count + 1)}>
            count is: {count}
          </button>
        </p>
        <HelloFromApi />
        <p>
          Edit <code>App.tsx</code> and save to test HMR updates.
        </p>
        <p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
          {' | '}
          <a
            className="App-link"
            href="https://vitejs.dev/guide/features.html"
            target="_blank"
            rel="noopener noreferrer"
          >
            Vite Docs
          </a>
        </p>
      </header>
    </div>
  )
}

export default App
