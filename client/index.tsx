import ReactDOM from 'react-dom/client'
import './src/assets/styles/index.css'
import App from './src/app'

window.localStorage.setItem('login','false')

ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
).render(<App/>)
